package cache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	lock sync.RWMutex
	data map[string][]byte
}

func New() *Cache {
	return &Cache{
		data: make(map[string][]byte),
	}
}

var _ Cacher = New()

func (c *Cache) Get(key []byte) ([]byte, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	keyStr := string(key)

	val, ok := c.data[keyStr]
	if !ok {
		return nil, fmt.Errorf("key (%s) not found", key)
	}

	return val, nil
}

func (c *Cache) Set(key []byte, value []byte, ttl time.Duration) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.data[string(key)] = value

	if ttl != 0 {
		go c.SetTTL(key, ttl)
	}

	return nil
}

func (c *Cache) Has(key []byte) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()

	_, ok := c.data[string(key)]

	return ok
}

func (c *Cache) Delete(key []byte) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.data, string(key))

	return nil
}

func (c *Cache) SetTTL(key []byte, duration time.Duration) {
	time.AfterFunc(duration, func() {
		c.lock.Lock()
		defer c.lock.Unlock()
		delete(c.data, string(key))
	})
}
