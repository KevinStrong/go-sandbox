package cache

import (
	"time"

	"github.com/kevinstrong/set"
)

type Cache[E comparable] struct {
	set           *set.Set[E]
	validDuration time.Duration
	cancelMap     map[E]chan struct{}
}

func NewWithExpiry[E comparable](expiration time.Duration) *Cache[E] {
	return &Cache[E]{
		set:           set.New[E](),
		validDuration: expiration,
		cancelMap:     map[E]chan struct{}{},
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

// this blows up in a concurrent environment
func (cache *Cache[E]) Add(elements ...E) {
	for _, value := range elements {
		value := value

		cancelChan, ok := cache.cancelMap[value]
		if ok {
			close(cancelChan)
		}

		cache.set.Add(value)

		timer := time.NewTimer(cache.validDuration)
		cancelChan = make(chan struct{})
		cache.cancelMap[value] = cancelChan
		go func() {
			select {
			case <-timer.C:
				cache.set.Delete(value)
			case <-cancelChan:
				return
			}
		}()
	}
}

func (cache *Cache[E]) Members() []E {
	return nil
}

func (cache *Cache[E]) Remove(element E) {
	cache.set.Delete(element)
}
