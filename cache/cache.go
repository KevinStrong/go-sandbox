package cache

import (
	"fmt"
	"time"

	"github.com/kevinstrong/set"
)

type Cache[E comparable] struct {
	set   *set.Set[E]
	adder Add[E]
}

type Add[E comparable] func(...E)

func WithLifetime[E comparable](expiration time.Duration) Option[E] {
	cancelMap := map[E]chan struct{}{}
	return func(c *Cache[E]) {
		cancelMap = map[E]chan struct{}{}
		c.adder = func(elements ...E) {
			for _, value := range elements {
				value := value

				cancelChan, ok := cancelMap[value]
				if ok {
					close(cancelChan)
				}

				c.set.Add(value)

				timer := time.NewTimer(expiration)
				cancelChan = make(chan struct{})
				cancelMap[value] = cancelChan
				go func() {
					kevin := value
					select {
					case <-timer.C:
						fmt.Printf("----------Hello!\n")
						fmt.Printf("----------Timeout of kevin: %v\n", kevin)
						c.set.Delete(kevin)
					case <-cancelChan:
						fmt.Printf("----------Timeout Cancelled for: %v\n", kevin)
						return
					}
				}()
			}
		}
	}
}

type Option[E comparable] func(*Cache[E])

func New[E comparable](options ...Option[E]) *Cache[E] {
	c := &Cache[E]{set: set.New[E]()}
	c.adder = func(e ...E) {
		c.set.Add(e...)
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
	cache.adder(elements...)
}

func (cache *Cache[E]) Members() []E {
	return nil
}

func (cache *Cache[E]) Remove(element E) {
	cache.set.Delete(element)
}
