package module

import (
	"log/slog"

	"github.com/go-chi/chi/v5"

	"github.com/ajay/portfolio-backend/internal/analytics/app"
	analyticshttp "github.com/ajay/portfolio-backend/internal/analytics/transport/http"
	"github.com/ajay/portfolio-backend/internal/common/bootstrap"
	"github.com/ajay/portfolio-backend/internal/common/config"
)

// Registrar returns the bootstrap hook for the analytics service.
// It's called by cmd/analytics/main.go so we can construct the app layer, hook
// up HTTP handlers, and perform any other wiring (DB connections, caches, etc.).
func Registrar() bootstrap.Registrar {
	return func(cfg config.ServiceConfig, log *slog.Logger, r chi.Router) error {
		svc := app.NewService(log)
		analyticshttp.Register(r, svc)
		log.Info("analytics service wired", "addr", cfg.HTTP.Addr())
		return nil
	}
}
