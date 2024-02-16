package prioriyqueue

import (
	"container/heap"
	"goalds/utils/comparator"
	"goalds/utils/locker"
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

type ElementHolder[T any] struct {
	elements []any
	cmpFun   comparator.Comparator[T]
}

func (eh *ElementHolder[T]) Push(x any) {
	eh.elements = append(eh.elements, x)
}

func (eh *ElementHolder[T]) Pop() any {
	item := eh.elements[eh.Len()-1]
	eh.elements = eh.elements[:eh.Len()-1]
	return item
}

func (eh *ElementHolder[T]) Len() int {
	return len(eh.elements)
}

func (eh *ElementHolder[T]) Less(i, j int) bool {
	return eh.cmpFun(eh.elements[i].(T), eh.elements[j].(T)) > 0
}

func (eh *ElementHolder[T]) Swap(i, j int) {
	eh.elements[i], eh.elements[j] = eh.elements[j], eh.elements[i]
}

type PriorityQueue[T any] struct {
	holder heap.Interface
	locker locker.Locker
}

func New[T any](cmpFunc comparator.Comparator[T], opts ...Option) *PriorityQueue[T] {
	option := &Options{
		locker: defaultLocker,
	}
	for _, opt := range opts {
		opt(option)
	}
	holder := &ElementHolder[T]{
		elements: make([]any, 0),
		cmpFun:   cmpFunc,
	}
	return &PriorityQueue[T]{
		holder: holder,
		locker: option.locker,
	}
}

func (pq *PriorityQueue[T]) Push(x T) {
	pq.locker.Lock()
	defer pq.locker.Unlock()
	heap.Push(pq.holder, x)
}

func (pq *PriorityQueue[T]) Pop() T {
	pq.locker.Lock()
	defer pq.locker.Unlock()
	return heap.Pop(pq.holder).(T)
}

func (pq *PriorityQueue[T]) Empty() bool {
	pq.locker.RLock()
	defer pq.locker.RUnlock()
	return pq.holder.Len() == 0
}

func (pq *PriorityQueue[T]) Len() int {
	pq.locker.RLock()
	defer pq.locker.RUnlock()
	return pq.holder.Len()
}

func (pq *PriorityQueue[T]) Top() T {
	pq.locker.RLock()
	defer pq.locker.RUnlock()
	return pq.holder.(*ElementHolder[T]).elements[0].(T)
}
