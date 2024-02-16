package gomap

import (
	"goalds/utils/locker"
	"goalds/utils/visitor"
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

type Map[K comparable, V any] struct {
	locker locker.Locker
	data   map[K]V
}

func New[K comparable, V any](options ...Option) *Map[K, V] {
	opt := &Options{locker: defaultLocker}
	for _, option := range options {
		option(opt)
	}
	return &Map[K, V]{locker: opt.locker, data: make(map[K]V)}
}

func (m *Map[K, V]) Set(key K, value V) {
	m.locker.Lock()
	defer m.locker.Unlock()

	m.data[key] = value
}

func (m *Map[K, V]) Get(key K) (V, bool) {
	m.locker.RLock()
	defer m.locker.RUnlock()

	value, ok := m.data[key]
	return value, ok
}

func (m *Map[K, V]) Has(key K) bool {
	m.locker.RLock()
	defer m.locker.RUnlock()

	_, ok := m.data[key]
	return ok
}

func (m *Map[K, V]) Erase(key K) {
	m.locker.Lock()
	defer m.locker.Unlock()

	delete(m.data, key)
}

func (m *Map[K, V]) Size() int {
	m.locker.RLock()
	defer m.locker.RUnlock()

	return len(m.data)
}

func (m *Map[K, V]) Clear() {
	m.locker.Lock()
	defer m.locker.Unlock()

	m.data = make(map[K]V)
}

func (m *Map[K, V]) Empty() bool {
	m.locker.RLock()
	defer m.locker.RUnlock()

	return len(m.data) == 0
}

func (m *Map[K, V]) Traversal(visitor visitor.KVVisitor[K, V]) {
	m.locker.RLock()
	defer m.locker.RUnlock()

	for k, v := range m.data {
		if !visitor(k, v) {
			break
		}
	}
}
