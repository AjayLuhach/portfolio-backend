package domain

import "time"

// Bookmark stores saved articles for later reading.
type Bookmark struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	BlogID    string    `json:"blogId"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"createdAt"`
}
