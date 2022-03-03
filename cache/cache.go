package cache

import (
	"time"

	"github.com/kevinstrong/set"
)

type Cache[E comparable] struct {
	set       *set.Set[E]
	adder     Add[E]
	cancelMap map[E]chan struct{}
}

type Add[E comparable] func(E)

func WithLifetime[E comparable](expiration time.Duration) Option[E] {
	return func(c *Cache[E]) {
		c.adder = func(value E) {
			cancelChan, ok := c.cancelMap[value]
			if ok {
				close(cancelChan)
			}

			c.set.Add(value)

			timer := time.NewTimer(expiration)
			cancelChan = make(chan struct{})
			c.cancelMap[value] = cancelChan
			go func() {
				select {
				case <-timer.C:
					c.set.Delete(value)
				case <-cancelChan:
					return
				}
			}()
		}
	}
}

type Option[E comparable] func(*Cache[E])

func New[E comparable](options ...Option[E]) *Cache[E] {
	c := &Cache[E]{
		set:       set.New[E](),
		cancelMap: map[E]chan struct{}{},
	}
	c.adder = func(e E) {
		c.set.Add(e)
	}

	for i := range options {
		options[i](c)
	}

	return c
}

func (cache *Cache[E]) Contains(element E) bool {
	return cache.set.Contains(element)
}

// this blows up in a concurrent environment
func (cache *Cache[E]) Add(elements ...E) {
	for _, element := range elements {
		cache.adder(element)
	}
}

func (cache *Cache[E]) Members() []E {
	return cache.set.Members()
}

func (cache *Cache[E]) Remove(element E) {
	cache.set.Delete(element)
}
