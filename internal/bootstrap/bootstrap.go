package bootstrap

import (
	"finkit/internal/cache"
	"finkit/internal/cli/currency"
	"finkit/internal/cli/inflation"
	"finkit/internal/cli/interest"
	"finkit/internal/cli/tax"
	"finkit/internal/cli/update"
	"finkit/internal/config"
	"finkit/internal/logger"
	"log/slog"
	"net/http"
	"time"
)

type App struct {
	Currency  *currency.Service
	Tax       *tax.Service
	Compound  *interest.Service
	Update    *update.Service
	Inflation *inflation.Service
	Logger    *slog.Logger
	Config    *config.Config
}

func BuildApp(verbose bool) *App {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	log := logger.InitLogger(verbose)
	fileCache := cache.NewFileCache(log)
	currencyProvider := currency.NewFrankfurterProvider(client, log)
	githubProvider := update.NewGithubProvider(client, log)

	return &App{
		Currency:  currency.NewService(currencyProvider, fileCache, log),
		Tax:       tax.NewService(log),
		Compound:  interest.NewService(log),
		Update:    update.NewService(githubProvider, log, client),
		Inflation: inflation.NewService(log),
		Logger:    log,
		Config:    config.NewConfig(),
	}
}
