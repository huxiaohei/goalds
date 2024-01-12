package container

type Container[T any] interface {
	PushBack(val T)
	PushFront(val T)
	PopBack() T
	PopFront() T
	Front() T
	Back() T
	Empty() bool
	Size() int
	String() string
	Clear()
}
