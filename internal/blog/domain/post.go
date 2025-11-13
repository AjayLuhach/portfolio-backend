package domain

import "time"

// Post captures authoring data for a blog entry.
type Post struct {
	ID          string    `json:"id"`
	AuthorID    string    `json:"authorId"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	Tags        []string  `json:"tags"`
	Status      string    `json:"status"`
	ScheduledAt time.Time `json:"scheduledAt"`
	PublishedAt time.Time `json:"publishedAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
