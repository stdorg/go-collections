package set

// Set is a generic set interface for comparable elements.
type Set[T comparable] interface {
	// Add adds the specified elements to this set if they are not already present.
	Add(e ...T)

	// Remove removes the specified elements from this set if they are present.
	Remove(e ...T)

	// Clear removes all elements from this set.
	Clear()

	// Contains checks whether this set contains the specified element.
	Contains(e T) bool

	// Len returns the number of elements in this set.
	Len() int

	// Chan returns a channel to iterate over the elements in this set.
	Chan() <-chan T

	// Slice returns a slice with the elements in this set.
	Slice() []T
}
