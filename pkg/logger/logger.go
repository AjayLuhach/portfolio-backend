package logger

import (
	"log/slog"
	"os"
	"strings"
)

// New creates a slog.Logger configured for structured JSON output.
func New(service string) *slog.Logger {
	level := slog.LevelInfo
	if raw := strings.ToLower(os.Getenv("LOG_LEVEL")); raw != "" {
		switch raw {
		case "debug":
			level = slog.LevelDebug
		case "warn":
			level = slog.LevelWarn
		case "error":
			level = slog.LevelError
		}
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	return slog.New(handler).With("service", service)
}
