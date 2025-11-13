package module

import (
	"log/slog"

	"github.com/go-chi/chi/v5"

	"github.com/ajay/portfolio-backend/internal/blog/app"
	bloghttp "github.com/ajay/portfolio-backend/internal/blog/transport/http"
	"github.com/ajay/portfolio-backend/internal/common/bootstrap"
	"github.com/ajay/portfolio-backend/internal/common/config"
)

func Registrar() bootstrap.Registrar {
	return func(cfg config.ServiceConfig, log *slog.Logger, r chi.Router) error {
		svc := app.NewService(log)
		bloghttp.Register(r, svc)
		log.Info("blog service wired", "addr", cfg.HTTP.Addr())
		return nil
	}
}
