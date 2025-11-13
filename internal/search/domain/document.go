package domain

// Document represents an indexed blog entry.
type Document struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Snippet string   `json:"snippet"`
	Tags    []string `json:"tags"`
	Score   float64  `json:"score"`
}
