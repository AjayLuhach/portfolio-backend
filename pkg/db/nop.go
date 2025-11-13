package db

import "context"

// NopStore is a placeholder database implementation used during scaffolding.
type NopStore struct{}

func NewNopStore() *NopStore { return &NopStore{} }

func (n *NopStore) Ping(_ context.Context) error { return nil }
