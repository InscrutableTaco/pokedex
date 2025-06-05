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
	mu       sync.RWMutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := new(Cache)
	cache.cacheMap = make(map[string]cacheEntry)
	cache.interval = interval
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cacheMap[key] = newEntry
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	var returnSlice []byte
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.cacheMap[key]
	if !ok {
		return returnSlice, false
	}
	returnSlice = entry.val
	return returnSlice, true
}
