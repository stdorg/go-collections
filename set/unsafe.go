package set

// concurrentUnsafe is a non-thread-safe set implementation.
type concurrentUnsafe[T comparable] struct {
	table map[T]struct{}
}

func NewUnsafe[T comparable]() Set[T] {
	return &concurrentUnsafe[T]{
		table: make(map[T]struct{}),
	}
}

func (s *concurrentUnsafe[T]) Add(e T) bool {
	initial := len(s.table)

	s.table[e] = struct{}{}

	return len(s.table) > initial
}

func (s *concurrentUnsafe[T]) Remove(e T) bool {
	_, ok := s.table[e]

	if ok {
		delete(s.table, e)
	}

	return ok
}

func (s *concurrentUnsafe[T]) Contains(e T) bool {
	_, ok := s.table[e]

	return ok
}

func (s *concurrentUnsafe[T]) Clear() {
	s.table = make(map[T]struct{})
}

func (s *concurrentUnsafe[T]) Len() int {
	return len(s.table)
}

func (s *concurrentUnsafe[T]) Slice() []T {
	slice := make([]T, 0, len(s.table))

	for key := range s.table {
		slice = append(slice, key)
	}

	return slice
}
