package cache

import (
	"time"

	"github.com/kevinstrong/set"
)

type Cache[E comparable] struct {
	set           *set.Set[E]
	validDuration time.Duration
}

func NewWithExpiry[E comparable](expiration time.Duration) *Cache[E] {
	return &Cache[E]{
		set:           set.New[E](),
		validDuration: expiration,
	}
}

func New[E comparable]() *Cache[E] {
	return &Cache[E]{
		set: set.New[E](),
	}
}

func (cache *Cache[E]) Contains(element E) bool {
	return cache.set.Contains(element)
}

func (cache *Cache[E]) Add(elements ...E) {
	cache.set.Add(elements...)
}

func (cache *Cache[E]) Members() []E {
	return nil
}

func (cache *Cache[E]) Remove(element E) {
	cache.set.Delete(element)
}
