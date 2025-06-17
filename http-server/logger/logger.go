package logger

import (
	"log/slog"
	"os"
)

func New(level slog.Leveler) *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	}))
}
