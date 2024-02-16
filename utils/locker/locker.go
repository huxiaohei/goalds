package locker

import "sync"

type Locker interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

var _ Locker = &sync.RWMutex{}

type FakeLocker struct{}

func (l FakeLocker) Lock() {}

func (l FakeLocker) Unlock() {}

func (l FakeLocker) RLock() {}

func (l FakeLocker) RUnlock() {}
