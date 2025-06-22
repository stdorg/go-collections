package set_test

import (
	"testing"

	"github.com/stdorg/go-collections/set"
)

// TestConcurrentUnsafe_AddRemove ensures `Add` and `Remove` operations work properly in a non current manner.
func TestConcurrentUnsafe_AddRemove(t *testing.T) {
	set := set.NewUnsafe[int]()

	// Add elements to the set:
	for i := 0; i < routines; i++ {
		set.Add(i)
	}

	// Remove elements from the set:
	for i := 0; i < routines; i++ {
		set.Remove(i)
	}

	// After all operations, the set should be empty.
	if set.Len() != 0 {
		t.Errorf("expected set to be empty, but got %d elements", set.Len())
	}
}

// TestConcurrentUnsafe_Contains ensures `Contains` operations work properly in a non current manner.
func TestConcurrentUnsafe_Contains(t *testing.T) {
	set := set.NewUnsafe[int]()

	// Add elements to the set:
	for i := 0; i < 1000; i++ {
		set.Add(i)
	}

	// Check for the presence of elements in the set.
	for i := 0; i < routines; i++ {
		if !set.Contains(i) {
			t.Errorf("expected set to contain %d", i)
		}
	}
}

// TestConcurrentUnsafe_AddRemoveContains tests a mix of operations in a non current manner.
func TestConcurrentUnsafe_AddRemoveContains(t *testing.T) {
	set := set.NewUnsafe[int]()

	// Add elements:
	for i := 0; i < routines; i++ {
		set.Add(i)
	}

	// Check for the presence of elements:
	for i := 0; i < routines; i++ {
		if !set.Contains(i) {
			t.Errorf("expected set to contain %d", i)
		}
	}

	// Remove elements:
	for i := 0; i < routines; i++ {
		set.Remove(i)
	}

	// We should have a empty set, because we are adding and removing elements sequentially.
	if set.Len() != 0 {
		t.Errorf("set length should be equals to 0, but got %d", set.Len())
	}
}
