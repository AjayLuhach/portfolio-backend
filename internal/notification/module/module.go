package module

import (
	"log/slog"

	"github.com/go-chi/chi/v5"

	"github.com/ajay/portfolio-backend/internal/common/bootstrap"
	"github.com/ajay/portfolio-backend/internal/common/config"
	"github.com/ajay/portfolio-backend/internal/notification/app"
	notificationhttp "github.com/ajay/portfolio-backend/internal/notification/transport/http"
)

func Registrar() bootstrap.Registrar {
	return func(cfg config.ServiceConfig, log *slog.Logger, r chi.Router) error {
		svc := app.NewService(log)
		notificationhttp.Register(r, svc)
		log.Info("notification service wired", "addr", cfg.HTTP.Addr())
		return nil
	}
}
