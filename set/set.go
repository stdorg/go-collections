package set

type Set[E comparable] interface {
	// Add adds the specified elements to this set if it is not already present.
	Add(e ...E)

	// Add removes the specified element from this set if it is present.
	Remove(e ...E)

	// Clear removes all of the elements from this set.
	Clear()

	// Contains checks whether this set contains the specified element.
	Contains(e E) bool

	// Len returns the number of elements in this set.
	Len() int

	// Chan returns a channel to iterate over the elements in this set.
	Chan() <-chan E

	// Slice returns a slice with the elements in this set.
	Slice() []E
}
