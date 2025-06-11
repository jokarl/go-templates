package logger

import (
	"log/slog"
	"os"
)

// Logger defines the interface for logging operations.
type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

// slogger is an implementation of the Logger interface using slog.
type slogger struct {
	logger *slog.Logger
}

// New creates a new Logger instance with a specified minimum log level.
func New(level slog.Leveler) Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	}))
	return &slogger{
		logger: logger,
	}
}

// Debug logs a message at DebugLevel.
func (l *slogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

// Info logs a message at InfoLevel.
func (l *slogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

// Warn logs a message at WarnLevel.
func (l *slogger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

// Error logs a message at ErrorLevel.
func (l *slogger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}
