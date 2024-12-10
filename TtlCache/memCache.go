package ttlcache

import (
	"sync"
	"time"
)

type cacheValue struct {
	value      any
	expiration int64
}

type MemoryCache struct {
	storage map[string]cacheValue
	mu      sync.RWMutex
}

func (c *MemoryCache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.storage[key] = cacheValue{value: value, expiration: time.Now().Add(ttl).UnixNano()}
}

func (c *MemoryCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.storage[key]
	if !found {
		return nil, false
	}

	if time.Now().UnixNano() > item.expiration {
		c.mu.RUnlock()
		c.Delete(key)
		c.mu.RLock()

		return nil, false
	}

	return item.value, true
}

func (c *MemoryCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.storage, key)
}

func New() *MemoryCache {
	return &MemoryCache{storage: make(map[string]cacheValue)}
}
