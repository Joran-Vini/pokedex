package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
type Cache struct {
	entries map[string]cacheEntry
	mu      sync.RWMutex
}

func NewCache(duration time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]cacheEntry),
		mu:      sync.RWMutex{},
	}
	go cache.cacheReapLoop(duration)
	return cache
}
