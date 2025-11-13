package app

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/ajay/portfolio-backend/internal/interaction/domain"
)

// Service owns reactions/comments orchestration.
type Service struct {
	log *slog.Logger
}

func NewService(log *slog.Logger) *Service {
	return &Service{log: log.With("component", "interaction-app")}
}

func (s *Service) CreateComment(ctx context.Context, comment domain.Comment) (domain.Comment, error) {
	comment.ID = uuid.NewString()
	comment.CreatedAt = time.Now()
	s.log.Info("comment created", "id", comment.ID)
	return comment, nil
}

func (s *Service) CreateReaction(ctx context.Context, reaction domain.Reaction) (domain.Reaction, error) {
	reaction.ID = uuid.NewString()
	s.log.Info("reaction created", "id", reaction.ID)
	return reaction, nil
}

func (s *Service) ListComments(ctx context.Context, blogID string) ([]domain.Comment, error) {
	return []domain.Comment{{ID: "comment-demo", BlogID: blogID, Body: "First!", CreatedAt: time.Now()}}, nil
}
