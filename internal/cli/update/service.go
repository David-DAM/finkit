package update

import (
	"context"
	"finkit/internal/version"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type Service struct {
	logger   *slog.Logger
	client   *http.Client
	Provider Provider
}

func NewService(provider Provider, logger *slog.Logger, client *http.Client) *Service {
	return &Service{Provider: provider, logger: logger, client: client}
}

func (s *Service) Do(ctx context.Context) (string, error) {
	release, err := s.Provider.GetLatestRelease(ctx)
	if err != nil {
		s.logger.Error("error getting latest release", "err", err)
		return "", err
	}

	if version.Version == release.TagName {
		return "You are already up to date!", nil
	}

	asset, err := s.findAsset(release)
	if err != nil {
		s.logger.Error("error finding asset", "err", err)
		return "", err
	}

	tmpFile := filepath.Join(os.TempDir(), "finkit_update")

	if err := s.DownloadToFile(ctx, asset.URL, tmpFile); err != nil {
		s.logger.Error("error downloading file", "err", err)
		return "", err
	}

	if err := s.ReplaceSelf(tmpFile); err != nil {
		s.logger.Error("error replacing self", "err", err)
		return "", err
	}

	return fmt.Sprintf("Updated to %s", release.TagName), nil
}

func (s *Service) findAsset(release *Release) (*Asset, error) {

	target := fmt.Sprintf("%s_%s", runtime.GOOS, runtime.GOARCH)

	for _, a := range release.Assets {
		if strings.Contains(a.Name, target) {
			return &a, nil
		}
	}

	return nil, fmt.Errorf("no asset found for %s", target)
}

func (s *Service) DownloadToFile(ctx context.Context, url, path string) error {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		s.logger.Error("error creating request", "err", err)
		return err
	}

	req.Header.Set("User-Agent", "finkit")

	resp, err := s.client.Do(req)
	if err != nil {
		s.logger.Error("error downloading file", "err", err)
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			derr := fmt.Errorf("failed to close response body: %w", err)
			s.logger.Error("failed to close response body", "error", derr)
		}
	}(resp.Body)

	out, err := os.Create(path)
	if err != nil {
		s.logger.Error("error creating output file", "err", err)
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			derr := fmt.Errorf("failed to close output file: %w", err)
			s.logger.Error("failed to close output file", "error", derr)
		}
	}(out)

	var (
		buf      = make([]byte, 32*1024)
		written  int64
		total    = resp.ContentLength
		lastDraw time.Time
	)

	for {
		n, err := resp.Body.Read(buf)
		if n > 0 {
			_, werr := out.Write(buf[:n])
			if werr != nil {
				s.logger.Error("failed to write to output file", "error", werr)
				return werr
			}

			written += int64(n)

			if time.Since(lastDraw) > 50*time.Millisecond && total > 0 {
				s.printProgress(written, total)
				lastDraw = time.Now()
			}
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			s.logger.Error("error reading response body", "err", err)
			return err
		}
	}

	s.printProgress(written, total)

	return nil
}

func (s *Service) printProgress(current, total int64) {

	if total <= 0 {
		s.logger.Info("download complete")
		return
	}

	percent := float64(current) / float64(total)

	width := 30
	filled := int(percent * float64(width))

	bar := strings.Repeat("█", filled) +
		strings.Repeat("░", width-filled)

	fmt.Printf("\r[%s] %3.0f%%", bar, percent*100)
}

func (s *Service) ReplaceSelf(tmpFile string) error {

	exe, err := os.Executable()
	if err != nil {
		s.logger.Error("error getting executable path", "err", err)
		return err
	}

	dir := filepath.Dir(exe)

	backup := filepath.Join(dir, "finkit.old")
	newPath := exe

	// remove old backup
	_ = os.Remove(backup)

	// backup current binary
	if err := os.Rename(exe, backup); err != nil {
		s.logger.Error("error backing up current binary", "err", err)
		return err
	}

	// move new binary into place
	if err := os.Rename(tmpFile, newPath); err != nil {
		// rollback
		_ = os.Rename(backup, exe)
		s.logger.Error("error replacing self", "err", err)
		return err
	}

	return nil
}
