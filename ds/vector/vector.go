package vector

import (
	"fmt"
	"goalds/utils/locker"
	"goalds/utils/visitor"
	"sync"
)

var defaultLocker locker.FakeLocker

type Options struct {
	capacity int
	locker   locker.Locker
}

type Option func(option *Options)

func WithCapacity(capacity int) Option {
	return func(option *Options) {
		option.capacity = capacity
	}
}

func WithGoroutineSafe() Option {
	return func(option *Options) {
		option.locker = &sync.RWMutex{}
	}
}

type Vector[T any] struct {
	data   []T
	locker locker.Locker
}

func New[T any](options ...Option) *Vector[T] {
	opt := Options{
		capacity: 0,
		locker:   defaultLocker,
	}
	for _, option := range options {
		option(&opt)
	}
	return &Vector[T]{
		data:   make([]T, 0, opt.capacity),
		locker: opt.locker,
	}
}

func NewFromVector[T any](other *Vector[T], options ...Option) *Vector[T] {
	defer other.locker.RUnlock()
	other.locker.RLock()
	opt := Options{
		capacity: other.Capacity(),
		locker:   defaultLocker,
	}
	for _, option := range options {
		option(&opt)
	}
	v := &Vector[T]{
		data:   make([]T, len(other.data), opt.capacity),
		locker: opt.locker,
	}
	copy(v.data, other.data)
	return v
}

func (v *Vector[T]) Size() int {
	defer v.locker.RUnlock()
	v.locker.RLock()
	return len(v.data)
}

func (v *Vector[T]) Capacity() int {
	defer v.locker.RUnlock()
	v.locker.RLock()
	return cap(v.data)
}

func (v *Vector[T]) Empty() bool {
	defer v.locker.RUnlock()
	v.locker.RLock()
	return len(v.data) == 0
}

func (v *Vector[T]) At(idx int) T {
	defer v.locker.RUnlock()
	v.locker.RLock()
	if idx < 0 || idx >= len(v.data) {
		panic(fmt.Sprintf("vector: out of range index: %d size: %d", idx, len(v.data)))
	}
	return v.data[idx]
}

func (v *Vector[T]) Front() T {
	return v.At(0)
}

func (v *Vector[T]) Back() T {
	return v.At(v.Size() - 1)
}

func (v *Vector[T]) PushBack(value T) {
	defer v.locker.Unlock()
	v.locker.Lock()
	v.data = append(v.data, value)
}

func (v *Vector[T]) PopBack() T {
	defer v.locker.Unlock()
	v.locker.Lock()
	if len(v.data) == 0 {
		panic("vector: pop back from empty array")
	}
	val := v.data[len(v.data)-1]
	v.data = v.data[:len(v.data)-1]
	return val
}

func (v *Vector[T]) InsertAt(idx int, value T) {
	defer v.locker.Unlock()
	v.locker.Lock()
	if idx < 0 || idx > len(v.data) {
		panic(fmt.Sprintf("vector: out of range index: %d size: %d", idx, len(v.data)))
	}
	v.data = append(v.data, value)
	copy(v.data[idx+1:], v.data[idx:])
	v.data[idx] = value
}

func (v *Vector[T]) EraseAt(idx int) {
	v.EraseRange(idx, idx+1)
}

func (v *Vector[T]) EraseRange(first, last int) {
	defer v.locker.Unlock()
	v.locker.Lock()
	if first < 0 || last > len(v.data) || first > last {
		panic(fmt.Sprintf("vector: out of range index: %d %d size: %d", first, last, len(v.data)))
	}
	v.data = append(v.data[:first], v.data[last:]...)
}

func (v *Vector[T]) ShrinkToFit() {
	defer v.locker.Unlock()
	if len(v.data) == cap(v.data) {
		return
	}
	len := len(v.data)
	data := make([]T, len)
	copy(data, v.data)
	v.data = data
}

func (v *Vector[T]) Clear() {
	defer v.locker.Unlock()
	v.locker.Lock()
	v.data = v.data[:0]
}

func (v *Vector[T]) String() string {
	defer v.locker.RUnlock()
	v.locker.RLock()
	return fmt.Sprintf("vector: %v", v.data)
}

func (v *Vector[T]) Traversal(visitor visitor.KVVisitor[int, T]) {
	defer v.locker.RUnlock()
	v.locker.RLock()

	for i, v := range v.data {
		if !visitor(i, v) {
			break
		}
	}
}
