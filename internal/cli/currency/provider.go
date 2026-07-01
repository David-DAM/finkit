package currency

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type Provider interface {
	GetRate(
		ctx context.Context,
		from string,
		to string,
	) (*Rate, error)
	SupportedCurrencies(
		ctx context.Context,
	) ([]Currency, error)
}

type FrankfurterProvider struct {
	baseUrl string
	client  *http.Client
	logger  *slog.Logger
}

const ApiUrl = "https://api.frankfurter.dev/v2"

func NewFrankfurterProvider(client *http.Client, logger *slog.Logger) *FrankfurterProvider {
	return &FrankfurterProvider{
		baseUrl: ApiUrl,
		client:  client,
		logger:  logger,
	}
}

func (p *FrankfurterProvider) GetRate(ctx context.Context, from string, to string) (*Rate, error) {

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		p.baseUrl+"/rate/"+from+"/"+to,
		nil,
	)
	if err != nil {
		p.logger.Error("error creating request", "err", err)
		return nil, err
	}

	response, err := p.client.Do(req)
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

func (p *FrankfurterProvider) SupportedCurrencies(ctx context.Context) ([]Currency, error) {

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		p.baseUrl+"/currencies",
		nil,
	)
	if err != nil {
		p.logger.Error("error creating request", "err", err)
		return nil, err
	}

	response, err := p.client.Do(req)
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
