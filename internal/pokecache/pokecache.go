package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mu           sync.Mutex
	interval     time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type cache interface {
	Add()
	Get()
	reapLoop()
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cacheEntries: make(map[string]cacheEntry),
		interval:     interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	fmt.Println("add to cache")
	c.mu.Lock()
	cacheEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cacheEntries[key] = cacheEntry
	defer c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	fmt.Println("get from cache")
	c.mu.Lock()
	defer c.mu.Unlock()
	cache, ok := c.cacheEntries[key]
	if !ok {
		return nil, false
	}
	return cache.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for {
		<-ticker.C
		c.mu.Lock()
		for key, entry := range c.cacheEntries {
			if time.Since(entry.createdAt) > c.interval {
				fmt.Printf("Reaping cache entry (key: %s)\n", key)
				delete(c.cacheEntries, key)
			}
		}
		c.mu.Unlock()
	}
}
