package app

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"

	"github.com/ajay/portfolio-backend/internal/auth/domain"
)

// Service contains business logic for authentication flows.
type Service struct {
	log *slog.Logger
}

func NewService(log *slog.Logger) *Service {
	return &Service{log: log.With("component", "auth-app")}
}

func (s *Service) Signup(ctx context.Context, creds domain.Credentials) (domain.AuthToken, error) {
	s.log.Info("signup", "email", creds.Email)
	return domain.AuthToken{
		AccessToken:  fmt.Sprintf("access-%s", uuid.NewString()),
		RefreshToken: fmt.Sprintf("refresh-%s", uuid.NewString()),
		ExpiresIn:    time.Hour,
	}, nil
}

func (s *Service) Login(ctx context.Context, creds domain.Credentials) (domain.AuthToken, error) {
	s.log.Info("login", "email", creds.Email)
	return domain.AuthToken{
		AccessToken:  fmt.Sprintf("access-%s", uuid.NewString()),
		RefreshToken: fmt.Sprintf("refresh-%s", uuid.NewString()),
		ExpiresIn:    time.Hour,
	}, nil
}

func (s *Service) Refresh(ctx context.Context, refreshToken string) (domain.AuthToken, error) {
	s.log.Debug("refresh token")
	if refreshToken == "" {
		return domain.AuthToken{}, fmt.Errorf("missing refresh token")
	}
	return domain.AuthToken{
		AccessToken:  fmt.Sprintf("access-%s", uuid.NewString()),
		RefreshToken: refreshToken,
		ExpiresIn:    time.Hour,
	}, nil
}
