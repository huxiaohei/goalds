package set

import (
	"goalds/utils/locker"
	"goalds/utils/visitor"
	"sync"
)

var defaultLocker locker.FakeLocker

type Options struct {
	locker locker.Locker
}

type Option func(option *Options)

func WithGoroutineSafe() Option {
	return func(option *Options) {
		option.locker = &sync.RWMutex{}
	}
}

type Set[T comparable] struct {
	locker locker.Locker
	data   map[T]bool
}

func New[T comparable](options ...Option) *Set[T] {
	opt := &Options{locker: defaultLocker}
	for _, option := range options {
		option(opt)
	}
	return &Set[T]{locker: opt.locker, data: make(map[T]bool)}
}

func (s *Set[T]) Insert(value T) {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.data[value] = true
}

func (s *Set[T]) Has(value T) bool {
	s.locker.RLock()
	defer s.locker.RUnlock()

	_, ok := s.data[value]
	return ok
}

func (s *Set[T]) Erase(value T) {
	s.locker.Lock()
	defer s.locker.Unlock()

	delete(s.data, value)
}

func (s *Set[T]) Size() int {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return len(s.data)
}

func (s *Set[T]) Clear() {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.data = make(map[T]bool)
}

func (s *Set[T]) Empty() bool {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return len(s.data) == 0
}

func (s *Set[T]) Traversal(visitor visitor.VVisitor[T]) {
	s.locker.RLock()
	defer s.locker.RUnlock()

	for value := range s.data {
		if !visitor(value) {
			break
		}
	}
}

func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	s.locker.RLock()
	defer s.locker.RUnlock()

	other.locker.RLock()
	defer other.locker.RUnlock()

	result := New[T]()
	for value := range s.data {
		if other.Has(value) {
			result.Insert(value)
		}
	}
	return result
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	s.locker.RLock()
	defer s.locker.RUnlock()

	other.locker.RLock()
	defer other.locker.RUnlock()

	result := New[T]()
	for value := range s.data {
		result.Insert(value)
	}
	for value := range other.data {
		result.Insert(value)
	}
	return result
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	s.locker.RLock()
	defer s.locker.RUnlock()

	other.locker.RLock()
	defer other.locker.RUnlock()

	result := New[T]()
	for value := range s.data {
		if !other.Has(value) {
			result.Insert(value)
		}
	}
	return result
}

func (s *Set[T]) IsSubset(other *Set[T]) bool {
	s.locker.RLock()
	defer s.locker.RUnlock()

	other.locker.RLock()
	defer other.locker.RUnlock()

	for value := range s.data {
		if !other.Has(value) {
			return false
		}
	}
	return true
}

func (s *Set[T]) IsSuperset(other *Set[T]) bool {
	return other.IsSubset(s)
}
