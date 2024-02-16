package queue

import (
	"goalds/ds/deque"
	"goalds/utils/container"
	"goalds/utils/locker"
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

type Queue[T any] struct {
	locker    locker.Locker
	container container.Container[T]
}

func New[T any](opts ...Option[T]) *Queue[T] {
	option := &Options[T]{
		locker:    locker.FakeLocker{},
		container: deque.New[T](),
	}
	for _, opt := range opts {
		opt(option)
	}

	return &Queue[T]{
		locker:    option.locker,
		container: option.container,
	}
}

func (q *Queue[T]) Push(val T) {
	q.locker.Lock()
	defer q.locker.Unlock()

	q.container.PushBack(val)
}

func (q *Queue[T]) Pop() T {
	q.locker.Lock()
	defer q.locker.Unlock()

	return q.container.PopFront()
}

func (q *Queue[T]) Front() T {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Front()
}

func (q *Queue[T]) Empty() bool {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Empty()
}

func (q *Queue[T]) Size() int {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.Size()
}

func (q *Queue[T]) String() string {
	q.locker.RLock()
	defer q.locker.RUnlock()

	return q.container.String()
}

func (q *Queue[T]) Clear() {
	q.locker.Lock()
	defer q.locker.Unlock()

	q.container.Clear()
}
