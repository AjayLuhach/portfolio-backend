package app

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/ajay/portfolio-backend/internal/blog/domain"
)

// Service manages blog posts lifecycle.
type Service struct {
	log *slog.Logger
}

func NewService(log *slog.Logger) *Service {
	return &Service{log: log.With("component", "blog-app")}
}

func (s *Service) CreateDraft(ctx context.Context, draft domain.Post) (domain.Post, error) {
	draft.ID = uuid.NewString()
	draft.Status = "draft"
	draft.UpdatedAt = time.Now()
	s.log.Info("draft created", "id", draft.ID)
	return draft, nil
}

func (s *Service) Publish(ctx context.Context, id string) (domain.Post, error) {
	post := domain.Post{ID: id, Status: "published", PublishedAt: time.Now()}
	s.log.Info("post published", "id", id)
	return post, nil
}

func (s *Service) Trending(ctx context.Context) ([]domain.Post, error) {
	return []domain.Post{
		{ID: "demo-trending", Title: "Building a Go Microservice", Status: "published"},
	}, nil
}

func (s *Service) History(ctx context.Context, id string) ([]domain.Post, error) {
	return []domain.Post{{ID: id, Status: "published"}}, nil
}
