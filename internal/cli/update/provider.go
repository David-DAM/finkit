package update

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type Provider interface {
	GetLatestRelease(ctx context.Context) (*Release, error)
}

type GithubProvider struct {
	client *http.Client
	logger *slog.Logger
}

func NewGithubProvider(client *http.Client, logger *slog.Logger) *GithubProvider {
	return &GithubProvider{client: client, logger: logger}
}

const (
	OWNER = "DAVID-DAM"
	REPO  = "finkit"
)

func (p *GithubProvider) GetLatestRelease(ctx context.Context) (*Release, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", OWNER, REPO)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)

	if err != nil {
		p.logger.Error("failed to create request", "error", err)
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("User-Agent", "finkit")
	resp, err := p.client.Do(req)
	if err != nil {
		p.logger.Error("failed to fetch latest release", "error", err)
		return nil, fmt.Errorf("failed to fetch latest release: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			derr := fmt.Errorf("failed to close response body: %w", err)
			p.logger.Error("failed to close response body", "error", derr)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		p.logger.Error("unexpected status code", "status_code", resp.StatusCode)
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var release Release

	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		p.logger.Error("failed to decode response", "error", err)
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	release.URL = url

	return &release, nil
}
