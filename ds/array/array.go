package array

import (
	"fmt"
	"goalds/utils/visitor"
)

type Array[T any] struct {
	data []T
}

func New[T any](size int) *Array[T] {
	return &Array[T]{
		data: make([]T, size),
	}
}

func NewFrom[T any](v ...T) *Array[T] {
	return &Array[T]{
		data: v,
	}
}

func Clone[T any](o *Array[T]) *Array[T] {
	s := &Array[T]{data: make([]T, len(o.data))}
	copy(s.data, o.data)
	return s
}

func (a *Array[T]) Size() int {
	return len(a.data)
}

func (a *Array[T]) Empty() bool {
	return len(a.data) == 0
}

func (a *Array[T]) At(idx int) T {
	if idx < 0 || idx >= len(a.data) {
		panic(fmt.Sprintf("array: out of range index: %d size: %d", idx, len(a.data)))
	}
	return a.data[idx]
}

func (a *Array[T]) Front() T {
	return a.At(0)
}

func (a *Array[T]) Back() T {
	return a.At(len(a.data) - 1)
}

func (a *Array[T]) Set(idx int, v T) {
	if idx < 0 || idx >= len(a.data) {
		panic(fmt.Sprintf("array: out of range index: %d size: %d", idx, len(a.data)))
	}
	a.data[idx] = v
}

func (a *Array[T]) Swap(o *Array[T]) {
	if a.Size() != o.Size() {
		panic(fmt.Sprintf("array: two arrays have different lengths len: %d len: %d", a.Size(), o.Size()))
	}
	a.data, o.data = o.data, a.data
}

func (a *Array[T]) String() string {
	return fmt.Sprintf("%v", a.data)
}

func (a *Array[T]) Traversal(visitor visitor.KVVisitor[int, T]) {
	for i, v := range a.data {
		if !visitor(i, v) {
			break
		}
	}
}
