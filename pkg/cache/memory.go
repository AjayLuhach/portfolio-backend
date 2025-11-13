package cache

import "sync"

// MemoryCache is a threadsafe map used for local development.
type MemoryCache struct {
	mu    sync.RWMutex
	store map[string]any
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{store: make(map[string]any)}
}

func (m *MemoryCache) Get(key string) (any, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, ok := m.store[key]
	return val, ok
}

func (m *MemoryCache) Set(key string, value any) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.store[key] = value
}

func (m *MemoryCache) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.store, key)
}
