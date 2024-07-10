package skiplist

import (
	"errors"
	"goalds/utils/comparator"
	"goalds/utils/locker"
	"goalds/utils/visitor"
	"math/rand"
	"sync"
	"time"
)

var (
	defaultLocker   locker.FakeLocker
	defaultMaxLevel = 16
)

type Options struct {
	maxLevel int
	locker   locker.Locker
}

type Option func(option *Options)

func WithMaxLevel(maxLevel int) Option {
	return func(option *Options) {
		option.maxLevel = maxLevel
	}
}

func WithGoroutineSafe() Option {
	return func(option *Options) {
		option.locker = &sync.RWMutex{}
	}
}

type Element[K, V any] struct {
	Node[K, V]
	key K
	val V
}

type Node[K, V any] struct {
	next []*Element[K, V]
}

type SkipList[K, V any] struct {
	locker         locker.Locker
	maxLevel       int
	head           Node[K, V]
	cmp            comparator.Comparator[K]
	len            int
	prevNodesCache []*Node[K, V]
	rander         *rand.Rand
}

func New[K, V any](cmp comparator.Comparator[K], opts ...Option) *SkipList[K, V] {
	option := Options{
		maxLevel: defaultMaxLevel,
		locker:   defaultLocker,
	}
	for _, opt := range opts {
		opt(&option)
	}

	sl := &SkipList[K, V]{
		locker:         option.locker,
		maxLevel:       option.maxLevel,
		head:           Node[K, V]{next: make([]*Element[K, V], option.maxLevel)},
		cmp:            cmp,
		len:            0,
		prevNodesCache: make([]*Node[K, V], option.maxLevel),
		rander:         rand.New(rand.NewSource(time.Now().Unix())),
	}
	return sl
}

func (sl *SkipList[K, V]) Insert(key K, val V) {
	sl.locker.Lock()
	defer sl.locker.Unlock()

	prevs := sl.findPrevNodes(key)
	for prevs[0].next[0] != nil && sl.cmp(prevs[0].next[0].key, key) == 0 {
		prevs[0].next[0].val = val
		return
	}

	level := sl.randomLevel()
	e := &Element[K, V]{
		key: key,
		val: val,
		Node: Node[K, V]{
			next: make([]*Element[K, V], level),
		},
	}
	for i := range e.Node.next {
		e.Node.next[i] = prevs[i].next[i]
		prevs[i].next[i] = e
	}
	sl.len++
}

func (sl *SkipList[K, V]) Get(key K) (V, error) {
	sl.locker.RLock()
	defer sl.locker.RUnlock()

	pre := &sl.head
	for i := sl.maxLevel - 1; i >= 0; i-- {
		cur := pre.next[i]
		for ; cur != nil; cur = cur.next[i] {
			cmpRet := sl.cmp(cur.key, key)
			if cmpRet == 0 {
				return cur.val, nil
			}
			if cmpRet > 0 {
				break
			}
			pre = &cur.Node
		}
	}
	return *new(V), errors.New("not found")
}

func (sl *SkipList[K, V]) Remove(key K) bool {
	sl.locker.Lock()
	defer sl.locker.Unlock()

	prevs := sl.findPrevNodes(key)
	element := prevs[0].next[0]
	if element == nil {
		return false
	}
	if sl.cmp(element.key, key) != 0 {
		return false
	}
	for i, v := range element.next {
		prevs[i].next[i] = v
	}
	sl.len--
	return true
}

func (sl *SkipList[K, V]) Len() int {
	sl.locker.RLock()
	defer sl.locker.RUnlock()

	return sl.len
}

func (sl *SkipList[K, V]) Traversal(visitor visitor.KVVisitor[K, V]) {
	sl.locker.RLock()
	defer sl.locker.RUnlock()

	for e := sl.head.next[0]; e != nil; e = e.Node.next[0] {
		if !visitor(e.key, e.val) {
			break
		}
	}
}

func (sl *SkipList[K, V]) Keys() []K {
	sl.locker.RLock()
	defer sl.locker.RUnlock()
	keys := make([]K, 0, sl.len)
	for e := sl.head.next[0]; e != nil; e = e.Node.next[0] {
		keys = append(keys, e.key)
	}
	return keys
}

func (sl *SkipList[K, V]) findPrevNodes(key K) []*Node[K, V] {
	prevs := sl.prevNodesCache
	prev := &sl.head
	for i := sl.maxLevel - 1; i >= 0; i-- {
		if sl.head.next[i] != nil {
			for next := prev.next[i]; next != nil; next = next.next[i] {
				if sl.cmp(next.key, key) >= 0 {
					break
				}
				prev = &next.Node
			}
		}
		prevs[i] = prev
	}
	return prevs
}

func (sl *SkipList[K, V]) randomLevel() int {
	total := uint64(1)<<uint64(sl.maxLevel) - 1 // 2^n-1
	k := sl.rander.Uint64() % total
	levelN := uint64(1) << (uint64(sl.maxLevel) - 1)

	level := 1
	for total -= levelN; total > k; level++ {
		levelN >>= 1
		total -= levelN
	}
	return level
}
