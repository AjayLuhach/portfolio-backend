package app

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/ajay/portfolio-backend/internal/bookmark/domain"
)

// Service stores bookmarks per user.
type Service struct {
	log *slog.Logger
}

func NewService(log *slog.Logger) *Service {
	return &Service{log: log.With("component", "bookmark-app")}
}

func (s *Service) Save(ctx context.Context, bookmark domain.Bookmark) (domain.Bookmark, error) {
	bookmark.ID = uuid.NewString()
	bookmark.CreatedAt = time.Now()
	s.log.Info("bookmark saved", "id", bookmark.ID)
	return bookmark, nil
}

func (s *Service) List(ctx context.Context, userID string) ([]domain.Bookmark, error) {
	return []domain.Bookmark{{ID: "bookmark-demo", UserID: userID, BlogID: "demo"}}, nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	s.log.Info("bookmark deleted", "id", id)
	return nil
}
