package deque

type Pool[T any] struct {
	segs []*Segment[T]
}

func newPoll[T any]() *Pool[T] {
	return &Pool[T]{
		segs: make([]*Segment[T], 0),
	}
}

func (p *Pool[T]) get() *Segment[T] {
	if len(p.segs) == 0 {
		return newSegment[T](segmentCapacity)
	}
	seg := p.segs[len(p.segs)-1]
	p.segs = p.segs[:len(p.segs)-1]
	return seg
}

func (p *Pool[T]) put(seg *Segment[T]) {
	p.segs = append(p.segs, seg)
}

func (p *Pool[T]) shrinkToSize(size int) {
	if len(p.segs) > size {
		_segs := make([]*Segment[T], size)
		copy(_segs, p.segs)
		p.segs = _segs
	}
}

func (p *Pool[T]) size() int {
	return len(p.segs)
}
