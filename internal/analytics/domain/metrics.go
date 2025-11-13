package domain

import "time"

// Metrics models aggregated engagement data.
type Metrics struct {
	BlogID     string    `json:"blogId"`
	Views      int       `json:"views"`
	ReadTime   float64   `json:"readTime"`
	BounceRate float64   `json:"bounceRate"`
	CapturedAt time.Time `json:"capturedAt"`
}
