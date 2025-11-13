package app

import "log/slog"

// Service proxies requests to downstream APIs (simplified for scaffolding).
type Service struct {
	log       *slog.Logger
	upstreams map[string]string
}

func NewService(log *slog.Logger, upstreams map[string]string) *Service {
	return &Service{log: log.With("component", "gateway-app"), upstreams: upstreams}
}

// Routes returns discovered upstream services for UI wiring/testing.
func (s *Service) Routes() map[string]string {
	return s.upstreams
}
