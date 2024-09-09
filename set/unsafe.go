package set

type concurrentUnsafe[E comparable] struct {
	table map[E]struct{}
}

func NewConcurrentUnsafe[E comparable]() Set[E] {
	return &concurrentUnsafe[E]{
		table: make(map[E]struct{}),
	}
}

func (set *concurrentUnsafe[E]) Add(e ...E) {
	for _, elem := range e {
		set.table[elem] = struct{}{}
	}
}

func (set *concurrentUnsafe[E]) Remove(e ...E) {
	for _, elem := range e {
		delete(set.table, elem)
	}
}

func (set *concurrentUnsafe[E]) Clear() {
	set.table = make(map[E]struct{})
}

func (set *concurrentUnsafe[E]) Contains(e E) bool {
	_, ok := set.table[e]

	return ok
}

func (set *concurrentUnsafe[E]) Len() int {
	return len(set.table)
}

func (set *concurrentUnsafe[E]) Chan() <-chan E {
	ch := make(chan E)

	go func() {
		defer close(ch)

		for key := range set.table {
			ch <- key
		}
	}()

	return ch
}

func (set *concurrentUnsafe[E]) Slice() []E {
	s := make([]E, 0, len(set.table))

	for key := range set.table {
		s = append(s, key)
	}

	return s
}
