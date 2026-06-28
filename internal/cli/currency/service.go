package currency

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

func (s *Service) Convert(from string, to string) (*Rate, error) {
	var rate *Rate

	err := s.cache.Get(fmt.Sprintf("rate-%s-%s", from, to), &rate)
	if err == nil {
		s.logger.Info("rate found in cache", "from", from, "to", to, "rate", rate)
		return rate, nil
	}

	s.logger.Error("rate not found in cache", "from", from, "to", to, "err", err)

	rate, err = s.provider.GetRate(from, to)
	if err != nil {
		s.logger.Error("error getting rate", "from", from, "to", to, "err", err)
		return nil, err
	}

	err = s.cache.Set(fmt.Sprintf("rate-%s-%s", from, to), rate, ttlRates)
	if err != nil {
		s.logger.Error("error setting rate in cache", "from", from, "to", to, "err", err)
	}

	return rate, nil
}

func (s *Service) Currencies() ([]Currency, error) {
	var currencies []Currency

	err := s.cache.Get("currencies", &currencies)
	if err == nil {
		s.logger.Info("currencies found in cache")
		return currencies, err
	}

	s.logger.Error("error getting currencies from cache", "err", err)

	currencies, err = s.provider.SupportedCurrencies()
	if err != nil {
		s.logger.Error("error getting supported currencies", "err", err)
		return nil, err
	}

	err = s.cache.Set("currencies", currencies, ttlCurrencies)
	if err != nil {
		s.logger.Error("error setting currencies in cache", "err", err)
	}

	return currencies, nil
}
