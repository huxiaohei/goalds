package list

import (
	"fmt"
	"goalds/utils/visitor"
)

type Node[T any] struct {
	Val  T
	next *Node[T]
	prev *Node[T]
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func (n *Node[T]) Prev() *Node[T] {
	return n.prev
}

type List[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

func New[T any]() *List[T] {
	return &List[T]{
		size: 0,
		head: nil,
		tail: nil,
	}
}

func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) Empty() bool {
	return l.size == 0
}

func (l *List[T]) Front() *Node[T] {
	return l.head
}

func (l *List[T]) Back() *Node[T] {
	return l.tail
}

func (l *List[T]) PushFront(val T) {
	node := &Node[T]{
		Val:  val,
		next: l.head,
		prev: nil,
	}
	if l.head != nil {
		l.head.prev = node
	}
	l.head = node
	if l.tail == nil {
		l.tail = l.head
	}
	l.size++
}

func (l *List[T]) PopFront() *Node[T] {
	if l.size == 0 {
		panic("list: empty list")
	}
	node := l.head
	l.head = l.head.next
	node.next = nil
	if l.head != nil {
		l.head.prev = nil
	} else {
		l.tail = nil
	}
	l.size--
	return node
}

func (l *List[T]) PushBack(val T) {
	node := &Node[T]{
		Val:  val,
		next: nil,
		prev: l.tail,
	}
	if l.tail != nil {
		l.tail.next = node
	}
	l.tail = node
	if l.head == nil {
		l.head = l.tail
	}
	l.size++
}

func (l *List[T]) PopBack() *Node[T] {
	if l.size == 0 {
		panic("list: empty list")
	}
	node := l.tail
	l.tail = l.tail.prev
	node.prev = nil
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.head = nil
	}
	l.size--
	return node
}

func (l *List[T]) At(index int) *Node[T] {
	if index >= l.size {
		panic(fmt.Sprintf("list: index out of range index: %d size: %d", index, l.size))
	}
	if index > l.size/2 {
		index = l.size - index - 1
		node := l.tail
		for i := 0; i < index; i++ {
			node = node.prev
		}
		return node
	} else {
		node := l.head
		for i := 0; i < index; i++ {
			node = node.next
		}
		return node
	}
}

func (l *List[T]) InsertAt(index int, val T) {
	if index <= 0 {
		l.PushFront(val)
		return
	}
	if index >= l.size {
		l.PushBack(val)
		return
	}
	node := l.At(index)
	newNode := &Node[T]{
		Val:  val,
		next: node,
		prev: node.prev,
	}
	node.prev.next = newNode
	node.prev = newNode
	l.size++
}

func (l *List[T]) RemoveAt(index int) *Node[T] {
	if index < 0 || index >= l.size {
		panic(fmt.Sprintf("list: index out of range index: %d size: %d", index, l.size))
	}
	if index == 0 {
		return l.PopFront()
	}
	if index == l.size-1 {
		return l.PopBack()
	}
	node := l.At(index)
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil
	node.prev = nil
	l.size--
	return node
}

func (l *List[T]) RemoveRange(start, end int) *Node[T] {
	if start < 0 || start >= end || end > l.size {
		panic(fmt.Sprintf("list: index out of range start: %d end: %d size: %d", start, end, l.size))
	}
	if start == 0 && end == l.size {
		node := l.head
		l.size = 0
		l.head = nil
		l.tail = nil
		return node
	}
	if start == 0 {
		node := l.At(end)
		node.prev.next = nil
		node.prev = nil
		head := l.head
		l.head = node
		l.size -= end
		return head
	}
	if end == l.size {
		node := l.At(start)
		l.tail = node.prev
		node.prev.next = nil
		node.prev = nil
		l.size = start
		return node
	}
	startNode := l.At(start)
	endNode := l.At(end)
	startNode.prev.next = endNode
	endNode.prev.next = nil
	endNode.prev = startNode.prev
	startNode.prev = nil
	l.size -= end - start
	return startNode
}

func (l *List[T]) Clear() {
	l.size = 0
	l.head = nil
	l.tail = nil
}

func (l *List[T]) String() string {
	str := "["
	node := l.head
	if node != nil {
		str += fmt.Sprintf("%v", node.Val)
	}
	for node.next != nil {
		node = node.next
		str += fmt.Sprintf(" %v", node.Val)
	}
	str += "]"
	return str
}

func (l *List[T]) Traversal(visitor visitor.VVisitor[T]) {
	node := l.head
	for node != nil {
		if !visitor(node.Val) {
			break
		}
		node = node.next
	}
}
