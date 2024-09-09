package set

import "sync"

type concurrentSafe[E comparable] struct {
	unsafe concurrentUnsafe[E]
	mutex  sync.RWMutex
}

func NewConcurrentSafe[E comparable]() Set[E] {
	return &concurrentSafe[E]{
		unsafe: concurrentUnsafe[E]{
			table: make(map[E]struct{}),
		},
	}
}

func (set *concurrentSafe[E]) Add(e ...E) {
	set.mutex.Lock()

	defer set.mutex.Unlock()

	set.unsafe.Add(e...)
}

func (set *concurrentSafe[E]) Remove(e ...E) {
	set.mutex.Lock()

	defer set.mutex.Unlock()

	set.unsafe.Remove(e...)
}

func (set *concurrentSafe[E]) Clear() {
	set.mutex.Lock()

	defer set.mutex.Unlock()

	set.unsafe.Clear()
}

func (set *concurrentSafe[E]) Contains(e E) bool {
	set.mutex.RLock()

	defer set.mutex.RUnlock()

	return set.unsafe.Contains(e)
}

func (set *concurrentSafe[E]) Len() int {
	set.mutex.RLock()

	defer set.mutex.RUnlock()

	return set.unsafe.Len()
}

func (set *concurrentSafe[E]) Chan() <-chan E {
	set.mutex.RLock()

	defer set.mutex.RUnlock()

	return set.unsafe.Chan()
}

func (set *concurrentSafe[E]) Slice() []E {
	set.mutex.RLock()

	defer set.mutex.RUnlock()

	return set.unsafe.Slice()
}
