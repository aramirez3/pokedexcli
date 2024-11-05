package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mu           sync.Mutex
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

func NewCache(interval time.Duration) *cacheEntry {
	var cacheEntry = cacheEntry{
		createdAt: time.Now(),
		val:       []byte{},
	}
	fmt.Println("Hello from fn")
	return &cacheEntry
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	interval := 5 * time.Second
	newCache := NewCache(interval)
	newCache.createdAt = time.Now()
	newCache.val = val
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cache, ok := c.cacheEntries[key]
	if !ok {
		return nil, false
	}
	return cache.val, true
}
