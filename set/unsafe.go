package set

// concurrentUnsafe is a non-thread-safe set implementation.
type concurrentUnsafe[E comparable] struct {
	table map[E]struct{}
}

// NewUnsafe creates a new non-thread-safe set.
func NewUnsafe[E comparable]() Set[E] {
	return &concurrentUnsafe[E]{
		table: make(map[E]struct{}),
	}
}

// Add adds the specified elements to the set.
func (s *concurrentUnsafe[E]) Add(e ...E) {
	for _, elem := range e {
		s.table[elem] = struct{}{}
	}
}

// Remove removes the specified elements from the set.
func (s *concurrentUnsafe[T]) Remove(e ...T) {
	for _, elem := range e {
		delete(s.table, elem)
	}
}

// Clear removes all elements from the set.
func (s *concurrentUnsafe[E]) Clear() {
	s.table = make(map[E]struct{})
}

// Contains checks if the set contains the specified element.
func (s *concurrentUnsafe[E]) Contains(e E) bool {
	_, ok := s.table[e]
	return ok
}

// Len returns the number of elements in the set.
func (s *concurrentUnsafe[E]) Len() int {
	return len(s.table)
}

// Chan returns a channel to iterate over the elements in the set.
func (s *concurrentUnsafe[E]) Chan() <-chan E {
	ch := make(chan E)

	go func() {
		defer close(ch)

		for key := range s.table {
			ch <- key
		}
	}()

	return ch
}

// Slice returns a slice with the elements in the set.
func (s *concurrentUnsafe[E]) Slice() []E {
	slice := make([]E, 0, len(s.table))

	for key := range s.table {
		slice = append(slice, key)
	}

	return slice
}
