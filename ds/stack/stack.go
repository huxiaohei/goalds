package stack

import (
	"goalds/ds/container"
	"goalds/ds/deque"
	"goalds/ds/locker"
	"sync"
)

type Options[T any] struct {
	locker    locker.Locker
	container container.Container[T]
}

type Option[T any] func(option *Options[T])

func WithGoroutineSafe[T any]() Option[T] {
	return func(option *Options[T]) {
		option.locker = &sync.RWMutex{}
	}
}

func WithContainer[T any](container container.Container[T]) Option[T] {
	return func(option *Options[T]) {
		option.container = container
	}
}

type Stack[T any] struct {
	locker    locker.Locker
	container container.Container[T]
}

func New[T any](opts ...Option[T]) *Stack[T] {
	option := &Options[T]{
		locker:    locker.FakeLocker{},
		container: deque.New[T](),
	}
	for _, opt := range opts {
		opt(option)
	}

	return &Stack[T]{
		locker:    option.locker,
		container: option.container,
	}
}

func (s *Stack[T]) Push(val T) {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.container.PushFront(val)
}

func (s *Stack[T]) Pop() T {
	s.locker.Lock()
	defer s.locker.Unlock()

	return s.container.PopFront()
}

func (s *Stack[T]) Top() T {
	s.locker.Lock()
	defer s.locker.Unlock()

	return s.container.Front()
}

func (s *Stack[T]) Empty() bool {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.Empty()
}

func (s *Stack[T]) Size() int {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.Size()
}

func (s *Stack[T]) Clear() {
	s.locker.Lock()
	defer s.locker.Unlock()

	s.container.Clear()
}

func (s *Stack[T]) String() string {
	s.locker.RLock()
	defer s.locker.RUnlock()

	return s.container.String()
}
