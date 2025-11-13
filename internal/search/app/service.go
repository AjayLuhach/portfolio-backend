package app

import (
	"context"
	"log/slog"

	"github.com/ajay/portfolio-backend/internal/search/domain"
)

// Service handles indexing + query orchestration.
type Service struct {
	log *slog.Logger
}

func NewService(log *slog.Logger) *Service {
	return &Service{log: log.With("component", "search-app")}
}

func (s *Service) Index(ctx context.Context, doc domain.Document) error {
	s.log.Info("document indexed", "id", doc.ID)
	return nil
}

func (s *Service) Query(ctx context.Context, term string) ([]domain.Document, error) {
	return []domain.Document{{ID: "doc-demo", Title: "Distributed Go", Snippet: "Microservices in Go", Score: 0.98}}, nil
}
