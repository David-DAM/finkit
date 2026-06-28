package convert

import (
	"finkit/internal/cache"
	"fmt"
	"log/slog"
	"time"
)

type Service struct {
	provider Provider
	cache    cache.Cache
	logger   *slog.Logger
}

func NewService(provider Provider, cache cache.Cache, logger *slog.Logger) *Service {
	return &Service{provider: provider, cache: cache, logger: logger}
}

const (
	ttlCurrencies = 6 * time.Hour
	ttlRates      = 15 * time.Minute
)

func (s *Service) Convert(amount float64, from string, to string) (float64, error) {
	var rate float64

	err := s.cache.Get(fmt.Sprintf("rate-%s-%s", from, to), &rate)
	if err == nil {
		s.logger.Info("rate found in cache", "from", from, "to", to, "rate", rate)
		return amount * rate, nil
	}

	s.logger.Error("rate not found in cache", "from", from, "to", to, "err", err)

	rate, err = s.provider.GetRate(from, to)
	if err != nil {
		s.logger.Error("error getting rate", "from", from, "to", to, "err", err)
		return 0, err
	}

	err = s.cache.Set(fmt.Sprintf("rate-%s-%s", from, to), rate, ttlRates)
	if err != nil {
		s.logger.Error("error setting rate in cache", "from", from, "to", to, "err", err)
	}

	return amount * rate, nil
}

func (s *Service) SupportedCurrencies() ([]string, error) {
	var currencies []string

	err := s.cache.Get("currencies", &currencies)
	if err != nil {
		s.logger.Error("error getting currencies from cache", "err", err)
	}

	currencies, err = s.provider.SupportedCurrencies()

	err = s.cache.Set("currencies", currencies, ttlCurrencies)
	if err != nil {
		s.logger.Error("error setting currencies in cache", "err", err)
	}

	return currencies, nil
}
