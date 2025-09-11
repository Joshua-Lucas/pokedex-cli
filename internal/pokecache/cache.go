package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu sync.Mutex
	m  map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {

	cache := &Cache{
		m: make(map[string]cacheEntry),
	}

	go cache.reapLoop(interval)

	return cache
}

// Generates and adds new cache entry
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()

	defer c.mu.Unlock()

	c.m[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Reads cache entry from cache and returns it if there is a value
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()

	defer c.mu.Unlock()

	v, ok := c.m[key]

	if !ok {
		return nil, false
	}

	return v.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:

			c.mu.Lock()

			for k, v := range c.m {
				timeSinceCached := time.Since(v.createdAt)

				if timeSinceCached > interval {
					delete(c.m, k)
				}
			}

			c.mu.Unlock()
		}
	}
}
