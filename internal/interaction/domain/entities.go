package domain

import "time"

// Comment models threaded comments.
type Comment struct {
	ID        string    `json:"id"`
	BlogID    string    `json:"blogId"`
	ParentID  string    `json:"parentId"`
	AuthorID  string    `json:"authorId"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
}

// Reaction expresses emoji style interactions.
type Reaction struct {
	ID       string `json:"id"`
	BlogID   string `json:"blogId"`
	Type     string `json:"type"`
	AuthorID string `json:"authorId"`
}
