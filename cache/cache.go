package cache

import (
	"time"

	"github.com/kevinstrong/set"
)

type Cache[E comparable] struct {
	set *set.Set[E]
}

func New[E comparable](expiration time.Duration) *Cache[E] {
	return &Cache[E]{
		set: set.New[E](),
	}
}

func (cache *Cache[E]) Add(elements ...E) {

}

func (cache *Cache[E]) Members() []E {
	return nil
}
