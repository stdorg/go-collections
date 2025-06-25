package set

import "sync"

// concurrentSafe is a thread-safe set implementation.
type concurrentSafe[T comparable] struct {
	unsafe concurrentUnsafe[T]
	mutex  sync.RWMutex
}

// NewSafe creates a new thread-safe set.
func NewSafe[T comparable]() Set[T] {
	return &concurrentSafe[T]{
		unsafe: concurrentUnsafe[T]{
			table: make(map[T]struct{}),
		},
	}
}

func (s *concurrentSafe[T]) Add(e T) bool {
	s.mutex.Lock()

	defer s.mutex.Unlock()

	return s.unsafe.Add(e)
}

func (s *concurrentSafe[T]) Remove(e T) bool {
	s.mutex.Lock()

	defer s.mutex.Unlock()

	return s.unsafe.Remove(e)
}

func (s *concurrentSafe[T]) Contains(e T) bool {
	s.mutex.RLock()

	defer s.mutex.RUnlock()

	return s.unsafe.Contains(e)
}

func (s *concurrentSafe[T]) Clear() {
	s.mutex.Lock()

	defer s.mutex.Unlock()

	s.unsafe.Clear()
}

func (s *concurrentSafe[T]) Len() int {
	s.mutex.RLock()

	defer s.mutex.RUnlock()

	return s.unsafe.Len()
}

func (s *concurrentSafe[T]) Slice() []T {
	s.mutex.Lock()

	defer s.mutex.Unlock()

	return s.unsafe.Slice()
}
