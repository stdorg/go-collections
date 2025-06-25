package set

// Set is a generic set interface for comparable elements.
type Set[T comparable] interface {
	// Add adds the specified element to this set if it is not already present.
	// Returns true if the element was added (i.e., it was not already in the set).
	Add(element T) bool

	// Remove removes the specified element from this set if it is present.
	// Returns true if the element was removed (i.e., it was in the set).
	Remove(element T) bool

	// Contains checks whether this set contains the specified element.
	Contains(element T) bool

	// Len returns the number of elements in this set.
	Len() int

	// Clear removes all elements from this set.
	Clear()

	// Slice returns a new slice containing all elements of the set.
	Slice() []T
}
