package cache_test

import (
	"testing"
	"time"

	"github.com/kevinstrong/cache"
)

func TestCacheExpiration(t *testing.T) {
	expiration := 1000 * time.Millisecond
	c := cache.NewWithExpiry[string](expiration)

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
