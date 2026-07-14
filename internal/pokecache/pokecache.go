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
	cacheMap map[string]cacheEntry
	mu       sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{cacheMap: make(map[string]cacheEntry)}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, data []byte) {
	c.mu.Lock()
	c.cacheMap[key] = cacheEntry{createdAt: time.Now(), val: data}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	data, ok := c.cacheMap[key]
	c.mu.Unlock()
	if !ok {
		return nil, false
	}
	return data.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key, val := range c.cacheMap {
			if time.Since(val.createdAt) > interval {
				delete(c.cacheMap, key)
			}
		}
		c.mu.Unlock()
	}
}
