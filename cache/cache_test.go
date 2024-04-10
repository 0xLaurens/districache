package cache

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCache_Get_EmptyCacheThrowsKeyNotFoundError(t *testing.T) {
	cache := New()
	key := "nomatch"
	_, err := cache.Get([]byte(key))
	assert.Error(t, err, fmt.Sprintf("key (%s) not found", key))
}

func TestCache_Get_KeyFound(t *testing.T) {
	cache := New()
	key := []byte("hello")
	value := []byte("mom")

	_ = cache.Set(key, value, 0)
	resp, err := cache.Get(key)
	assert.NoError(t, err)

	assert.Equal(t, string(value), string(resp))
}

func TestCache_Set(t *testing.T) {
	cache := New()

	key := []byte("hello")
	value := []byte("mom")
	ttl := 100 * time.Millisecond

	err := cache.Set(key, value, ttl)
	assert.NoError(t, err)
	// assert the value was inserted
	assert.Equal(t, true, cache.Has(key))

	time.Sleep(100 * time.Millisecond)
	// check if value is deleted after 100ms
	assert.Equal(t, false, cache.Has(key))
}

func TestCache_Delete(t *testing.T) {
	cache := New()
	key := []byte("hello")
	value := []byte("mom")
	ttl := time.Duration(0)

	err := cache.Delete(key)
	assert.NoError(t, err)

	_ = cache.Set(key, value, ttl)
	res, err := cache.Get(key)
	assert.NoError(t, err)
	assert.NotNil(t, res)

	err = cache.Delete(key)
	assert.NoError(t, err)

	assert.Equal(t, false, cache.Has(key))
}

func BenchmarkSetGetDelete(b *testing.B) {
	cache := New()
	key := []byte("hello")
	value := []byte("mom")
	ttl := time.Duration(0)

	for i := 0; i < b.N; i++ {
		_ = cache.Set(key, value, ttl)
		_, _ = cache.Get(key)
		_ = cache.Delete(key)
	}
}
