package cache_test

import (
	"testing"
	"time"

	"github.com/KevinStrong/go-sandbox/cache"
)

func TestCacheExpiration(t *testing.T) {
	expiration := 5 * time.Millisecond
	c := cache.New[string](
		cache.WithLifetime[string](expiration),
	)

	testString := "/v1/users/1"
	c.Add(testString)

	if !c.Contains(testString) {
		t.Fatal("cache should contain testString after add")
	}

	time.Sleep(expiration + (10 * time.Millisecond))

	if c.Contains(testString) {
		t.Fatal("cache should remove testString after expiration duration")
	}
}

func TestCacheExpiration_DoubleAdd(t *testing.T) {
	expiration := 1000 * time.Millisecond
	c := cache.New[string](
		cache.WithLifetime[string](expiration),
	)

	testString := "/v1/users/1"
	c.Add(testString)
	time.Sleep(500 * time.Millisecond)

	c.Add(testString)

	time.Sleep(600 * time.Millisecond)

	if !c.Contains(testString) {
		t.Fatal("cache failed to reset expiration timer on a 2nd add")
	}

	time.Sleep(400*time.Millisecond + 10*time.Millisecond)

	if c.Contains(testString) {
		t.Fatal("cache failed to remove updated value from 2nd add")
	}
}

// our code demands an expiry time currently, please fix.
func TestCacheGet_WithExistingValue_ReturnTrue(t *testing.T) {
	c := cache.New[string]()
	testString := "/v1/users/1"

	isInBeforeAdd := c.Contains(testString)
	if isInBeforeAdd {
		t.Fatal("value is present in cache before add is called")
	}

	c.Add(testString)
	isInCache := c.Contains(testString)

	if !isInCache {
		t.Fatal("test string should be in cache")
	}
}

func TestCacheGet_WithNonExistingValue_ReturnFalse(t *testing.T) {
	c := cache.New[string]()
	testString := "/v1/users/1"

	isInBeforeAdd := c.Contains(testString)
	if isInBeforeAdd {
		t.Error("value is present in cache before add is called")
	}

	isInCache := c.Contains(testString)

	if isInCache {
		t.Fatal("test string should not be in cache")
	}
}

func TestCacheRemove(t *testing.T) {
	c := cache.New[int]()
	testValue := 1

	c.Add(testValue)
	if !c.Contains(testValue) {
		t.Fatal("test value should be in cache after add is called")
	}

	c.Remove(testValue)

	if c.Contains(testValue) {
		t.Fatal("test value still present after remove is called")
	}
}
