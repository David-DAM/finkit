package bootstrap

import (
	"finkit/internal/cache"
	"finkit/internal/cli/currency/convert"
	"finkit/internal/logger"
	"log/slog"
)

type App struct {
	Convert *convert.Service
	Logger  *slog.Logger
}

func BuildApp(verbose bool) *App {

	log := logger.InitLogger(verbose)
	provider := convert.NewFrankfurterProvider()
	fileCache := cache.NewFileCache()

	return &App{
		Convert: convert.NewService(provider, fileCache, log),
		Logger:  log,
	}
}
