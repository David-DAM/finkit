package bootstrap

import (
	"finkit/internal/cache"
	"finkit/internal/cli/currency"
	"finkit/internal/logger"
	"log/slog"
)

type App struct {
	Currency *currency.Service
	Logger   *slog.Logger
}

func BuildApp(verbose bool) *App {

	log := logger.InitLogger(verbose)
	provider := currency.NewFrankfurterProvider(log)
	fileCache := cache.NewFileCache()

	return &App{
		Currency: currency.NewService(provider, fileCache, log),
		Logger:   log,
	}
}
