package app

import (
	"context"
	"log/slog"
	"time"

	"github.com/ajay/portfolio-backend/internal/analytics/domain"
)

// Service aggregates analytics for blogs/authors.
type Service struct {
	log *slog.Logger
}

// NewService wires the slog logger. In a real system this is also the spot where
// you'd inject repositories, caches, queue publishers, etc.
func NewService(log *slog.Logger) *Service {
	return &Service{log: log.With("component", "analytics-app")}
}

// Record pretends to persist the provided metrics. For now we just stamp the
// timestamp, log the call, and echo the payload back to the caller.
func (s *Service) Record(ctx context.Context, metrics domain.Metrics) (domain.Metrics, error) {
	metrics.CapturedAt = time.Now()
	s.log.Info("metrics recorded", "blogId", metrics.BlogID)
	return metrics, nil
}

// Fetch simulates retrieving stored metrics for a blog. Hard-coded numbers keep
// the sample predictable while still showing how a query path would look.
func (s *Service) Fetch(ctx context.Context, blogID string) (domain.Metrics, error) {
	return domain.Metrics{BlogID: blogID, Views: 128, ReadTime: 4.2, CapturedAt: time.Now()}, nil
}
