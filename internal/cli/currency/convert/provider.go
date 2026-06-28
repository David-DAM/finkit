package convert

import (
	"encoding/json"
	"finkit/internal/cli/currency"
	"io"
	"net/http"
)

type Provider interface {
	GetRate(
		from string,
		to string,
	) (float64, error)
	SupportedCurrencies() ([]string, error)
}

type FrankfurterProvider struct {
	baseUrl string
}

func NewFrankfurterProvider() *FrankfurterProvider {
	return &FrankfurterProvider{
		baseUrl: "https://api.frankfurter.dev/v2",
	}
}

func (p *FrankfurterProvider) GetRate(from string, to string) (float64, error) {

	response, err := http.Get(p.baseUrl + "/rate/" + from + "/" + to)
	if err != nil {
		return 0, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	var rate currency.Rate
	err = json.Unmarshal(body, &rate)
	if err != nil {
		return 0, err
	}

	return rate.Rate, nil
}

func (p *FrankfurterProvider) SupportedCurrencies() ([]string, error) {
	response, err := http.Get(p.baseUrl + "/currencies")
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var currencies currency.Currencies
	err = json.Unmarshal(body, &currencies)
	if err != nil {
		return nil, err
	}

	var supportedCurrencies []string
	for _, c := range currencies.Currencies {
		supportedCurrencies = append(supportedCurrencies, c.IsoCode)
	}

	return supportedCurrencies, nil
}
