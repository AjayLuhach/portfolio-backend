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
//
// Every service implements a registrar (see internal/<service>/module) that accepts
// the resolved configuration + logger, registers its HTTP routes, and returns an error
// if any dependency (databases, caches, etc.) fails to initialize. Keeping this as a
// function type makes it trivial for the `cmd/<service>` packages to plug in their
// specific wiring while reusing the same bootstrap.
type Registrar func(cfg config.ServiceConfig, log *slog.Logger, r chi.Router) error

// Run initializes shared infrastructure and starts the HTTP server.
//
// High level steps:
//   1. Load the service-specific configuration driven by env vars.
//   2. Build a slog logger that tags every entry with the service name.
//   3. Instantiate the shared chi router + middleware stack.
//   4. Call the registrar so the service can attach handlers/adapters.
//   5. Start the HTTP server with graceful shutdown hooks.
//
// If any step above fails we bubble the error up to main(), which will log.Fatal.
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
