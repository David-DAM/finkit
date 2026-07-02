package inflation

import (
	"log/slog"
	"math"
)

type Service struct {
	logger *slog.Logger
}

func NewService(logger *slog.Logger) *Service {
	return &Service{logger: logger}
}

func (s *Service) Do(amount float64, years int, rate float64) float64 {
	return amount / math.Pow(1+rate/100, float64(years))
}
