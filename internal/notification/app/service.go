package app

import (
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/ajay/portfolio-backend/internal/notification/domain"
)

// Service manages email/websocket notifications.
type Service struct {
	log *slog.Logger
}

func NewService(log *slog.Logger) *Service {
	return &Service{log: log.With("component", "notification-app")}
}

func (s *Service) Dispatch(ctx context.Context, n domain.Notification) (domain.Notification, error) {
	n.ID = uuid.NewString()
	n.CreatedAt = time.Now()
	s.log.Info("notification dispatched", "id", n.ID, "channel", n.Channel)
	return n, nil
}

func (s *Service) List(ctx context.Context, userID string) ([]domain.Notification, error) {
	return []domain.Notification{{ID: "notif-demo", UserID: userID, Message: "Welcome!", Channel: "email", CreatedAt: time.Now()}}, nil
}
