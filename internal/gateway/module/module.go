package module

import (
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/ajay/portfolio-backend/internal/common/bootstrap"
	"github.com/ajay/portfolio-backend/internal/common/config"
	"github.com/ajay/portfolio-backend/internal/gateway/app"
	gatewayhttp "github.com/ajay/portfolio-backend/internal/gateway/transport/http"
)

func Registrar() bootstrap.Registrar {
	return func(cfg config.ServiceConfig, log *slog.Logger, r chi.Router) error {
		upstreams := map[string]string{
			"auth":         getenv("AUTH_SERVICE_URL", "http://localhost:9101"),
			"blog":         getenv("BLOG_SERVICE_URL", "http://localhost:9102"),
			"interaction":  getenv("INTERACTION_SERVICE_URL", "http://localhost:9103"),
			"notification": getenv("NOTIFICATION_SERVICE_URL", "http://localhost:9104"),
			"analytics":    getenv("ANALYTICS_SERVICE_URL", "http://localhost:9105"),
			"bookmark":     getenv("BOOKMARK_SERVICE_URL", "http://localhost:9106"),
			"search":       getenv("SEARCH_SERVICE_URL", "http://localhost:9107"),
		}

		svc := app.NewService(log, upstreams)
		gatewayhttp.Register(r, svc)
		log.Info("gateway wired", "addr", cfg.HTTP.Addr())
		return nil
	}
}

func getenv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}
