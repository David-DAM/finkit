package tax

import "log/slog"

type Service struct {
	logger *slog.Logger
}

func NewService(logger *slog.Logger) *Service {
	return &Service{logger: logger}
}

const unlimitedBracketLimit = 0

func (s *Service) CalculateTaxSalary(salary float64, country string) float64 {
	if len(country) == 0 {
		s.logger.Error("country must be provided")
		return 0
	}

	if country != "ES" {
		s.logger.Warn("tax calculation not implemented for country", "country", country)
		return 0
	}

	return s.calculateSpainTax(salary)
}

func (s *Service) calculateSpainTax(salary float64) float64 {
	if salary <= 0 {
		return 0
	}

	var tax float64

	brackets := []Bracket{
		{12450, 0.19},
		{20200, 0.24},
		{35200, 0.30},
		{60000, 0.37},
		{300000, 0.45},
		{unlimitedBracketLimit, 0.47}, // no limit for top bracket
	}

	remaining := salary
	previousLimit := 0.0

	for _, bracket := range brackets {
		if remaining <= 0 {
			break
		}

		taxableAmount := taxableAmountForBracket(salary, remaining, previousLimit, bracket)

		tax += taxableAmount * bracket.Rate
		remaining -= taxableAmount
		previousLimit = bracket.Limit
	}

	return tax
}

func taxableAmountForBracket(salary, remaining, previousLimit float64, bracket Bracket) float64 {
	if bracket.Limit == unlimitedBracketLimit {
		return remaining
	}

	if salary > bracket.Limit {
		return bracket.Limit - previousLimit
	}

	return remaining
}
