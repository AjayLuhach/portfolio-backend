package module

import (
	"log/slog"

	"github.com/go-chi/chi/v5"

	"github.com/ajay/portfolio-backend/internal/auth/app"
	authhttp "github.com/ajay/portfolio-backend/internal/auth/transport/http"
	"github.com/ajay/portfolio-backend/internal/common/bootstrap"
	"github.com/ajay/portfolio-backend/internal/common/config"
)

// Registrar exposes the bootstrap registrar for the auth service.
func Registrar() bootstrap.Registrar {
	return func(cfg config.ServiceConfig, log *slog.Logger, r chi.Router) error {
		service := app.NewService(log)
		authhttp.Register(r, service)
		log.Info("auth service wired", "addr", cfg.HTTP.Addr())
		return nil
	}
}
