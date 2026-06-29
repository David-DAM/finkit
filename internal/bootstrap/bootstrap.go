package bootstrap

import (
	"finkit/internal/cache"
	"finkit/internal/cli/currency"
	"finkit/internal/cli/tax"
	"finkit/internal/logger"
	"log/slog"
)

type App struct {
	Currency *currency.Service
	Tax      *tax.Service
	Logger   *slog.Logger
}

func BuildApp(verbose bool) *App {

	log := logger.InitLogger(verbose)
	provider := currency.NewFrankfurterProvider(log)
	fileCache := cache.NewFileCache(log)

	return &App{
		Currency: currency.NewService(provider, fileCache, log),
		Tax:      tax.NewService(log),
		Logger:   log,
	}
}
