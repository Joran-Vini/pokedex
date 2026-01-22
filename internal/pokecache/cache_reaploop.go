package pokecache

import (
	"time"
)

func (c *Cache) cacheReapLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > duration {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
