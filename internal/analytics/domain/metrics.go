package domain

import "time"

// Metrics models aggregated engagement data.
// JSON tags line up with the HTTP handlers so the same struct can be marshaled
// in and out of requests without extra DTOs.
type Metrics struct {
	BlogID     string    `json:"blogId"`
	Views      int       `json:"views"`
	ReadTime   float64   `json:"readTime"`
	BounceRate float64   `json:"bounceRate"`
	CapturedAt time.Time `json:"capturedAt"`
}
