package pokecache

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	interval := 5 * time.Second
	cache := NewCache(interval)
	t.Log(cache)
}
