package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Data     map[string]cacheEntry
	Mu       *sync.Mutex
	Interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{Data: make(map[string]cacheEntry), Mu: &sync.Mutex{}, Interval: interval}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	entry := cacheEntry{createdAt: time.Now(), val: val}
	c.Data[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	if entry, ok := c.Data[key]; ok {
		return entry.val, ok
	}
	return nil, false

}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.Interval)
	for range ticker.C {
		c.Mu.Lock()
		for key, v := range c.Data {
			if time.Since(v.createdAt) > c.Interval {
				delete(c.Data, key)
			}
		}
		c.Mu.Unlock()
	}
}
