package set

import "sync"

type Set[E comparable] struct {
	data map[E]struct{}
	lock *sync.Mutex
}

func New[E comparable]() *Set[E] {
	return &Set[E]{
		data: map[E]struct{}{},
		lock: &sync.Mutex{},
	}
}

func (set *Set[E]) Add(elements ...E) {
	set.lock.Lock()
	defer set.lock.Unlock()

	for _, element := range elements {
		set.data[element] = struct{}{}
	}
}

func (set Set[E]) Union(anotherSet Set[E]) *Set[E] {
	unionSet := New[E]()
	unionSet.Add(set.Members()...)
	unionSet.Add(anotherSet.Members()...)
	return unionSet
}

func (set Set[E]) Members() []E {
	set.lock.Lock()
	defer set.lock.Unlock()

	members := make([]E, 0, len(set.data))
	for key := range set.data {
		members = append(members, key)
	}
	return members
}
