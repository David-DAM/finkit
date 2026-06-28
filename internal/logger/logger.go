package logger

import (
	"io"
	"log/slog"
	"os"
)

func InitLogger(verbose bool) *slog.Logger {
	var logger *slog.Logger
	if verbose {
		logger = slog.New(slog.NewJSONHandler(os.Stderr, nil))
	} else {
		logger = slog.New(slog.NewJSONHandler(io.Discard, nil))
	}
	return logger
}
