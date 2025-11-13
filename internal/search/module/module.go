package module

import (
	"log/slog"

	"github.com/go-chi/chi/v5"

	"github.com/ajay/portfolio-backend/internal/common/bootstrap"
	"github.com/ajay/portfolio-backend/internal/common/config"
	"github.com/ajay/portfolio-backend/internal/search/app"
	searchhttp "github.com/ajay/portfolio-backend/internal/search/transport/http"
)

func Registrar() bootstrap.Registrar {
	return func(cfg config.ServiceConfig, log *slog.Logger, r chi.Router) error {
		svc := app.NewService(log)
		searchhttp.Register(r, svc)
		log.Info("search service wired", "addr", cfg.HTTP.Addr())
		return nil
	}
}
