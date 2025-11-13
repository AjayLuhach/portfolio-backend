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

func NewService(log *slog.Logger) *Service {
	return &Service{log: log.With("component", "analytics-app")}
}

func (s *Service) Record(ctx context.Context, metrics domain.Metrics) (domain.Metrics, error) {
	metrics.CapturedAt = time.Now()
	s.log.Info("metrics recorded", "blogId", metrics.BlogID)
	return metrics, nil
}

func (s *Service) Fetch(ctx context.Context, blogID string) (domain.Metrics, error) {
	return domain.Metrics{BlogID: blogID, Views: 128, ReadTime: 4.2, CapturedAt: time.Now()}, nil
}
