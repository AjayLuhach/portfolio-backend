package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/ajay/portfolio-backend/internal/common/config"
)

// HTTPServer wraps the standard library server with graceful shutdown wiring.
type HTTPServer struct {
	server *http.Server
	log    *slog.Logger
}

// NewHTTPServer wraps the chi router with sensible timeouts supplied by config.
// It deliberately keeps the struct tiny so swapping the http.Handler (REST vs gRPC
// proxies, etc.) stays trivial.
func NewHTTPServer(cfg config.ServiceConfig, handler http.Handler, log *slog.Logger) *HTTPServer {
	return &HTTPServer{
		server: &http.Server{
			Addr:         cfg.HTTP.Addr(),
			Handler:      handler,
			ReadTimeout:  cfg.HTTP.ReadTimeout,
			WriteTimeout: cfg.HTTP.WriteTimeout,
		},
		log: log,
	}
}

// Start launches the HTTP server and blocks until it stops or the context is cancelled.
//
// The context usually comes from bootstrap, which wires SIGINT/SIGTERM into the
// cancellation path. When cancellation happens we shut down gracefully with a
// 5-second timeout so in-flight requests get a chance to finish.
func (s *HTTPServer) Start(ctx context.Context) error {
	shutdownCh := make(chan struct{})
	go func() {
		<-ctx.Done()
		ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := s.server.Shutdown(ctxTimeout); err != nil {
			s.log.Error("graceful shutdown failed", "error", err)
		}
		close(shutdownCh)
	}()

	s.log.Info("http server listening", "addr", s.server.Addr)
	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	<-shutdownCh
	return nil
}
