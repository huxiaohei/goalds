package deque

import "fmt"

const segmentCapacity = 128

type Deque[T any] struct {
	pool  *Pool[T]
	segs  []*Segment[T]
	begin int
	end   int
	size  int
}

func New[T any]() *Deque[T] {
	return &Deque[T]{
		pool:  newPoll[T](),
		segs:  make([]*Segment[T], 0),
		begin: 0,
		end:   0,
		size:  0,
	}
}

func (d *Deque[T]) Size() int {
	return d.size
}

func (d *Deque[T]) Empty() bool {
	return d.size == 0
}

func (d *Deque[T]) PushFront(val T) {
	seg := d.firstAvailableSegment()
	seg.pushFront(val)
	d.size++
}

func (d *Deque[T]) PushBack(val T) {
	d.lastAvailableSegment().pushBack(val)
	d.size++
	if d.segmentUsed() >= len(d.segs) {
		d.expand()
	}
}

func (d *Deque[T]) InsertAt(index int, val T) {
	if index <= 0 {
		d.PushFront(val)
		return
	}
	if index >= d.size {
		d.PushBack(val)
		return
	}
	segPos, valPos := d.pos(index)
	if segPos < d.segmentUsed()-segPos {
		d.moveFrontInsert(segPos, valPos, val)
	} else {
		d.moveBackInsert(segPos, valPos, val)
	}
	d.size++
	if d.segmentUsed() >= len(d.segs) {
		d.expand()
	}
}

func (d *Deque[T]) Front() T {
	return d.firstSegment().front()
}

func (d *Deque[T]) Back() T {
	return d.lastSegment().back()
}

func (d *Deque[T]) At(index int) T {
	if index < 0 || index >= d.size {
		panic(fmt.Sprintf("deque: out of range index: %d size: %d", index, d.size))
	}
	segPos, valPos := d.pos(index)
	return d.segmentAt(segPos).at(valPos)
}

func (d *Deque[T]) Set(index int, val T) {
	if index < 0 || index >= d.size {
		panic(fmt.Sprintf("deque: out of range index: %d size: %d", index, d.size))
	}
	segPos, valPos := d.pos(index)
	d.segmentAt(segPos).set(valPos, val)
}

func (d *Deque[T]) PopFront() T {
	if d.size == 0 {
		panic("deque: PopFront on empty deque")
	}
	s := d.segs[d.begin]
	val := s.popFront()
	if s.empty() {
		d.putToPool(s)
		d.segs[d.begin] = nil
		d.begin = d.nextIndex(d.begin)
	}
	d.size--
	d.shrinkIfNeeded()
	return val
}

func (d *Deque[T]) PopBack() T {
	if d.size == 0 {
		panic("deque: PopBack on empty deque")
	}
	s := d.segs[d.prevIndex(d.end)]
	val := s.popBack()
	if s.empty() {
		d.putToPool(s)
		d.segs[d.prevIndex(d.end)] = nil
		d.end = d.prevIndex(d.end)
	}
	d.size--
	d.shrinkIfNeeded()
	return val
}

func (d *Deque[T]) EraseAt(index int) T {
	if index < 0 || index >= d.size {
		panic(fmt.Sprintf("deque: out of range index: %d size: %d", index, d.size))
	}
	segPos, valPos := d.pos(index)
	e := d.segmentAt(segPos).eraseAt(valPos)
	if segPos < d.segmentUsed()-segPos-1 {
		for i := segPos; i > 0; i-- {
			cur := d.segmentAt(i)
			prev := d.segmentAt(i - 1)
			cur.pushFront(prev.popBack())
		}
		if d.firstSegment().empty() {
			d.putToPool(d.firstSegment())
			d.segs[d.begin] = nil
			d.begin = d.nextIndex(d.begin)
			d.shrinkIfNeeded()
		}
	} else {
		for i := segPos; i < d.segmentUsed()-1; i++ {
			cur := d.segmentAt(i)
			next := d.segmentAt(i + 1)
			cur.pushBack(next.popFront())
		}
		if d.lastSegment().empty() {
			d.putToPool(d.lastSegment())
			d.segs[d.prevIndex(d.end)] = nil
			d.end = d.prevIndex(d.end)
			d.shrinkIfNeeded()
		}
	}
	d.size--
	return e
}

func (d *Deque[T]) EraseRange(startIndex, endIndex int) bool {
	if startIndex < 0 || startIndex >= d.size || endIndex < 0 || endIndex > d.size || startIndex >= endIndex {
		return false
	}
	num := endIndex - startIndex
	if d.size-startIndex < endIndex {
		// move back
		for index := startIndex; index+num < d.size; index++ {
			d.Set(index, d.At(index+num))
		}
		for i := 0; i < num; i++ {
			d.PopBack()
		}
	} else {
		// move front
		for index := endIndex - 1; index-num >= 0; index-- {
			d.Set(index, d.At(index-num))
		}
		for i := 0; i < num; i++ {
			d.PopFront()
		}
	}
	return true
}

func (d *Deque[T]) Clear() {
	d.EraseRange(0, d.size)
}

func (d *Deque[T]) String() string {
	str := "["
	if !d.Empty() {
		str += fmt.Sprintf("%v", d.Front())
	}
	for i := 1; i < d.size; i++ {
		str += fmt.Sprintf(" %v", d.At(i))
	}
	str += "]"
	return str
}

func (d *Deque[T]) segmentUsed() int {
	if d.size == 0 {
		return 0
	}
	if d.end > d.begin {
		return d.end - d.begin
	}
	return len(d.segs) - d.begin + d.end
}

func (d *Deque[T]) firstSegment() *Segment[T] {
	if len(d.segs) == 0 {
		return nil
	}
	return d.segs[d.begin]
}

func (d *Deque[T]) lastSegment() *Segment[T] {
	if len(d.segs) == 0 {
		return nil
	}
	return d.segs[d.prevIndex(d.end)]
}

func (d *Deque[T]) firstAvailableSegment() *Segment[T] {
	firstSegment := d.firstSegment()
	if firstSegment != nil && !firstSegment.full() {
		return firstSegment
	}
	if d.segmentUsed() >= len(d.segs) {
		d.expand()
	}
	firstSegment = d.firstSegment()
	if firstSegment == nil || firstSegment.full() {
		d.begin = d.prevIndex(d.begin)
		s := d.pool.get()
		d.segs[d.begin] = s
		return s
	}
	return firstSegment
}

func (d *Deque[T]) lastAvailableSegment() *Segment[T] {
	lastSegment := d.lastSegment()
	if lastSegment != nil && !lastSegment.full() {
		return lastSegment
	}
	if d.segmentUsed() >= len(d.segs) {
		d.expand()
	}
	lastSegment = d.lastSegment()
	if lastSegment == nil || lastSegment.full() {
		s := d.pool.get()
		d.segs[d.end] = s
		d.end = d.nextIndex(d.end)
		return s
	}
	return lastSegment
}

func (d *Deque[T]) expand() {
	capacity := d.segmentUsed() * 2
	if capacity == 0 {
		capacity = 1
	}
	segs := make([]*Segment[T], capacity)
	for i := 0; i < d.segmentUsed(); i++ {
		segs[i] = d.segs[(d.begin+i)%d.segmentUsed()]
	}
	n := d.segmentUsed()
	d.begin = 0
	d.end = n
	d.segs = segs
}

func (d *Deque[T]) shrinkIfNeeded() {
	if int(float64(d.segmentUsed()*2)*1.2) < cap(d.segs) {
		capacity := cap(d.segs) / 2
		segs := make([]*Segment[T], capacity)
		for i := 0; i < d.segmentUsed(); i++ {
			segs[i] = d.segs[(d.begin+i)%len(d.segs)]
		}
		n := d.segmentUsed()
		d.begin = 0
		d.end = n
		d.segs = segs
	}
}

func (d *Deque[T]) nextIndex(index int) int {
	return (index + 1) % len(d.segs)
}

func (d *Deque[T]) prevIndex(index int) int {
	return (index - 1 + len(d.segs)) % len(d.segs)
}

func (d *Deque[T]) pos(index int) (int, int) {
	if index <= d.firstSegment().size()-1 {
		return 0, index
	}
	index -= d.firstSegment().size()
	return index/segmentCapacity + 1, index % segmentCapacity
}

func (d *Deque[T]) segmentAt(seg int) *Segment[T] {
	return d.segs[(d.begin+seg)%cap(d.segs)]
}

func (d *Deque[T]) moveFrontInsert(segPos, insertPos int, val T) {
	if d.firstSegment().full() {
		if d.segmentUsed() >= len(d.segs) {
			d.expand()
		}
		d.begin = d.prevIndex(d.begin)
		d.segs[d.begin] = d.pool.get()
		if insertPos == 0 {
			insertPos = segmentCapacity - 1
		} else {
			segPos++
			insertPos--
		}
	} else {
		if insertPos == 0 {
			segPos--
			insertPos = segmentCapacity - 1
		} else {
			if segPos != 0 {
				insertPos--
			}
		}
	}
	for i := 0; i < segPos; i++ {
		cur := d.segmentAt(i)
		next := d.segmentAt(i + 1)
		cur.pushBack(next.popFront())
	}
	d.segmentAt(segPos).insert(insertPos, val)
}

func (d *Deque[T]) moveBackInsert(segPos, insertPos int, val T) {
	if d.lastSegment().full() {
		if d.segmentUsed() >= len(d.segs) {
			d.expand()
		}
		d.segs[d.end] = d.pool.get()
		d.end = d.nextIndex(d.end)
	}
	for i := d.segmentUsed() - 1; i > segPos; i-- {
		cur := d.segmentAt(i)
		prev := d.segmentAt(i - 1)
		cur.pushFront(prev.popBack())
	}
	d.segmentAt(segPos).insert(insertPos, val)
}

func (d *Deque[T]) putToPool(s *Segment[T]) {
	s.clear()
	d.pool.put(s)
	if d.pool.size()*6/5 > d.segmentUsed() {
		d.pool.shrinkToSize(d.segmentUsed() / 5)
	}
}
