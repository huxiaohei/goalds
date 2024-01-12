package deque

import "fmt"

type Segment[T any] struct {
	data  []T
	begin int
	end   int
	nsize int
}

func newSegment[T any](capacity int) *Segment[T] {
	return &Segment[T]{
		data:  make([]T, capacity),
		begin: 0,
		end:   0,
		nsize: 0,
	}
}

func (s *Segment[T]) size() int {
	return s.nsize
}

func (s *Segment[T]) empty() bool {
	return s.nsize == 0
}

func (s *Segment[T]) full() bool {
	return s.nsize == len(s.data)
}

func (s *Segment[T]) at(index int) T {
	if index < 0 || index >= s.nsize {
		panic(fmt.Sprintf("segment: index out of range index: %d size: %d", index, s.nsize))
	}
	return s.data[(s.begin+index)%len(s.data)]
}

func (s *Segment[T]) front() T {
	return s.at(0)
}

func (s *Segment[T]) back() T {
	return s.at(s.nsize - 1)
}

func (s *Segment[T]) set(index int, val T) {
	if index < 0 || index >= s.nsize {
		panic(fmt.Sprintf("segment: index out of range index: %d size: %d", index, s.nsize))
	}
	s.data[(s.begin+index)%len(s.data)] = val
}

func (s *Segment[T]) pushBack(val T) {
	if s.full() {
		panic("segment: pushBack on full segment")
	}
	s.data[s.end] = val
	s.end = s.nextIndex(s.end)
	s.nsize++
}

func (s *Segment[T]) pushFront(val T) {
	if s.full() {
		panic("segment: pushFront on full segment")
	}
	s.begin = s.prevIndex(s.begin)
	s.data[s.begin] = val
	s.nsize++
}

func (s *Segment[T]) insert(index int, val T) {
	if index < s.nsize-index { // move front
		idx := s.prevIndex(s.begin)
		for i := 0; i < index; i++ {
			s.data[idx] = s.data[s.nextIndex(idx)]
			idx = s.nextIndex(idx)
		}
		s.data[idx] = val
		s.begin = s.prevIndex(s.begin)
	} else { // move back
		idx := s.end
		for i := 0; i < s.nsize-index; i++ {
			s.data[idx] = s.data[s.prevIndex(idx)]
			idx = s.prevIndex(idx)
		}
		s.data[idx] = val
		s.end = s.nextIndex(s.end)
	}
	s.nsize += 1
}

func (s *Segment[T]) popBack() T {
	if s.empty() {
		panic("segment: popBack on empty segment")
	}
	s.end = s.prevIndex(s.end)
	s.nsize--
	// todo how to clear the value?
	return s.data[s.end]
}

func (s *Segment[T]) popFront() T {
	if s.empty() {
		panic("segment: popFront on empty segment")
	}
	val := s.data[s.begin]
	s.begin = s.nextIndex(s.begin)
	s.nsize--
	// todo how to clear the value?
	return val
}

func (s *Segment[T]) eraseAt(index int) T {
	e := s.data[(s.begin+index)%len(s.data)]
	if index < s.nsize-index {
		for i := index; i > 0; i-- {
			index = (i + s.begin) % len(s.data)
			s.data[index] = s.data[(i-1+s.begin)%len(s.data)]
		}
		s.begin = s.nextIndex(s.begin)
	} else {
		for i := index; i < s.nsize-1; i++ {
			index = (i + s.begin) % len(s.data)
			s.data[index] = s.data[(i+1+s.begin)%len(s.data)]
		}
		s.end = s.prevIndex(s.end)
	}
	s.nsize--
	return e
}

func (s *Segment[T]) clear() {
	s.begin = 0
	s.end = 0
	s.nsize = 0
	s.data = make([]T, len(s.data))
}

func (s *Segment[T]) nextIndex(index int) int {
	return (index + 1) % len(s.data)
}

func (s *Segment[T]) prevIndex(index int) int {
	return (index - 1 + len(s.data)) % len(s.data)
}
