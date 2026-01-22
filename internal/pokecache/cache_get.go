package pokecache

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, exists := c.entries[key]
	if !exists {
		return nil, false
	}
	return value.val, true
}
