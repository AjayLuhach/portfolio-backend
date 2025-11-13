package domain

import "time"

// Notification models user facing alerts.
type Notification struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	Message   string    `json:"message"`
	Channel   string    `json:"channel"`
	Read      bool      `json:"read"`
	CreatedAt time.Time `json:"createdAt"`
}
