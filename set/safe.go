package set

import "sync"

// concurrentSafe is a thread-safe set implementation.
type concurrentSafe[E comparable] struct {
	unsafe concurrentUnsafe[E]
	mutex  sync.RWMutex
}

// NewSafe creates a new thread-safe set.
func NewSafe[E comparable]() Set[E] {
	return &concurrentSafe[E]{
		unsafe: concurrentUnsafe[E]{
			table: make(map[E]struct{}),
		},
	}
}

// Add adds the specified elements to the set.
func (s *concurrentSafe[E]) Add(e ...E) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.unsafe.Add(e...)
}

// Remove removes the specified elements from the set.
func (s *concurrentSafe[E]) Remove(e ...E) {
	s.mutex.Lock()

	defer s.mutex.Unlock()

	s.unsafe.Remove(e...)
}

// Clear removes all elements from the set.
func (s *concurrentSafe[E]) Clear() {
	s.mutex.Lock()

	defer s.mutex.Unlock()

	s.unsafe.Clear()
}

// Contains checks if the set contains the specified element.
func (s *concurrentSafe[E]) Contains(e E) bool {
	s.mutex.RLock()

	defer s.mutex.RUnlock()

	return s.unsafe.Contains(e)
}

// Len returns the number of elements in the set.
func (s *concurrentSafe[E]) Len() int {
	s.mutex.RLock()

	defer s.mutex.RUnlock()

	return s.unsafe.Len()
}

// Chan returns a channel to iterate over the elements in the set.
func (s *concurrentSafe[T]) Chan() <-chan T {
	s.mutex.RLock()

	defer s.mutex.RUnlock()

	return s.unsafe.Chan()
}

// Slice returns a slice with the elements in the set.
func (s *concurrentSafe[E]) Slice() []E {
	s.mutex.RLock()

	defer s.mutex.RUnlock()

	return s.unsafe.Slice()
}
