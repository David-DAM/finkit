package update

import (
	"context"
	"finkit/internal/version"
	"fmt"
	"log/slog"
)

type Service struct {
	logger   *slog.Logger
	Provider Provider
}

func NewService(provider Provider, logger *slog.Logger) *Service {
	return &Service{Provider: provider, logger: logger}
}

func (s *Service) Check(ctx context.Context) (*Release, error) {
	release, err := s.Provider.GetLatestRelease(ctx)
	if err != nil {
		s.logger.Error("error getting latest release", "err", err)
		return nil, err
	}

	if version.Version == release.TagName {
		return nil, fmt.Errorf("you are already up to date")
	}

	return release, nil
}
