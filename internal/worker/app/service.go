package app

import (
	"context"
	"log/slog"
	"time"
)

// Service simulates async job orchestration.
type Service struct {
	log *slog.Logger
}

func NewService(log *slog.Logger) *Service {
	return &Service{log: log.With("component", "worker-app")}
}

func (s *Service) Enqueue(ctx context.Context, job string, payload map[string]any) error {
	s.log.Info("job enqueued", "name", job, "payload", payload)
	time.Sleep(100 * time.Millisecond)
	s.log.Info("job completed", "name", job)
	return nil
}
