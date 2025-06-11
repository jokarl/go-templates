package logger

import (
	"context"
	"log/slog"
	"os"
)

// Logger defines the interface for logging operations.
type Logger interface {
	Debug(msg string, args ...any)
	DebugContext(ctx context.Context, msg string, args ...any)
	Info(msg string, args ...any)
	InfoContext(ctx context.Context, msg string, args ...any)
	Warn(msg string, args ...any)
	WarnContext(ctx context.Context, msg string, args ...any)
	Error(msg string, args ...any)
	ErrorContext(ctx context.Context, msg string, args ...any)
}

// slogger is an implementation of the Logger interface using slog.
type slogger struct {
	logger *slog.Logger
}

// New creates a new Logger instance with a specified minimum log level.
func New(level slog.Leveler) Logger {
	return &slogger{
		logger: slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
			Level: level,
		})),
	}
}

// Debug logs a message at slog.LevelDebug.
func (l *slogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

// DebugContext logs a message at slog.LevelDebug with a context.
func (l *slogger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.logger.DebugContext(ctx, msg, args...)
}

// Info logs a message at slog.LevelInfo.
func (l *slogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

// InfoContext logs a message at slog.LevelInfo with a context.
func (l *slogger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.logger.InfoContext(ctx, msg, args...)
}

// Warn logs a message at WarnLevel.
func (l *slogger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

// WarnContext logs a message at slog.LevelWarn with a context.
func (l *slogger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.logger.WarnContext(ctx, msg, args...)
}

// Error logs a message at ErrorLevel.
func (l *slogger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

// ErrorContext logs a message at slog.LevelError with a context.
func (l *slogger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.logger.ErrorContext(ctx, msg, args...)
}
