package domain

import "time"

// User models an authenticated account.
type User struct {
	ID        string
	Email     string
	Name      string
	CreatedAt time.Time
}

// Credentials is used for signup/login flows.
type Credentials struct {
	Email    string
	Password string
	Name     string
}

// AuthToken encapsulates access + refresh tokens.
type AuthToken struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    time.Duration
}
