package currency

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type Provider interface {
	GetRate(
		from string,
		to string,
	) (*Rate, error)
	SupportedCurrencies() ([]Currency, error)
}

type FrankfurterProvider struct {
	baseUrl string
	logger  *slog.Logger
}

func NewFrankfurterProvider(logger *slog.Logger) *FrankfurterProvider {
	return &FrankfurterProvider{
		baseUrl: "https://api.frankfurter.dev/v2",
		logger:  logger,
	}
}

func (p *FrankfurterProvider) GetRate(from string, to string) (*Rate, error) {

	response, err := http.Get(p.baseUrl + "/rate/" + from + "/" + to)
	if err != nil {
		p.logger.Error("error getting rate", "from", from, "to", to, "err", err)
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			p.logger.Error("error closing response body", "err", err)
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		p.logger.Error("error reading response body", "err", err)
		return nil, err
	}

	var rate Rate
	err = json.Unmarshal(body, &rate)
	if err != nil {
		p.logger.Error("error unmarshalling response body", "err", err)
		return nil, err
	}

	return &rate, nil
}

func (p *FrankfurterProvider) SupportedCurrencies() ([]Currency, error) {
	response, err := http.Get(p.baseUrl + "/currencies")
	if err != nil {
		p.logger.Error("error getting supported currencies", "err", err)
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			p.logger.Error("error closing response body", "err", err)
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		p.logger.Error("error reading response body", "err", err)
		return nil, err
	}

	var currencies []Currency
	err = json.Unmarshal(body, &currencies)
	if err != nil {
		p.logger.Error("error unmarshalling response body", "err", err)
		return nil, err
	}

	return currencies, nil
}
