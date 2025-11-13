package bootstrap

import (
	"context"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"

	"github.com/ajay/portfolio-backend/internal/common/config"
	"github.com/ajay/portfolio-backend/internal/common/server"
	"github.com/ajay/portfolio-backend/pkg/logger"
)

// Registrar represents the function each service provides to wire routes + dependencies.
type Registrar func(cfg config.ServiceConfig, log *slog.Logger, r chi.Router) error

// Run initializes shared infrastructure and starts the HTTP server.
func Run(serviceName string, registrar Registrar) error {
	cfg := config.Load(serviceName)
	log := logger.New(serviceName)
	router := server.NewRouter(log)

	if err := registrar(cfg, log, router); err != nil {
		return err
	}

	srv := server.NewHTTPServer(cfg, router, log)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := srv.Start(ctx); err != nil {
		log.Error("http server stopped", "error", err)
		return err
	}

	log.Info("service shutdown complete")
	return nil
}
