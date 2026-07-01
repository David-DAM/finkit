package interest

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

func (s *Service) Do(initialInvestment float64, monthlyContribution float64, years int, rate float64) float64 {

	months := years * 12
	monthlyRate := rate / 100 / 12

	// Calculate future value of initial investment
	futureValueInitial := initialInvestment * math.Pow(1+monthlyRate, float64(months))

	// Calculate future value of monthly contributions (annuity)
	var futureValueContributions float64
	if monthlyRate > 0 {
		futureValueContributions = monthlyContribution * ((math.Pow(1+monthlyRate, float64(months)) - 1) / monthlyRate)
	} else {
		futureValueContributions = monthlyContribution * float64(months)
	}

	return futureValueInitial + futureValueContributions
}
