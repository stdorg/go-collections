package set_test

import (
	"sync"
	"testing"

	"github.com/stdorg/go-collections/set"
)

const routines = 750

// TestConcurrentSafe_AddRemove ensures concurrent access to `Add` and `Remove` operations.
func TestConcurrentSafe_AddRemove(t *testing.T) {
	set := set.NewConcurrentSafe[int]()

	var wg sync.WaitGroup

	// We'll perform `Add`, and `Remove` operations concurrently.
	wg.Add(routines * 2)

	// Concurrently add elements to the set:
	for i := 0; i < routines; i++ {
		go func(i int) {
			defer wg.Done()

			set.Add(i)
		}(i)
	}

	// Concurrently remove elements from the set:
	for i := 0; i < routines; i++ {
		go func(i int) {
			defer wg.Done()

			set.Remove(i)
		}(i)
	}

	// Wait for all operations to complete.
	wg.Wait()

	// After all operations, the set should be empty.
	if set.Len() != 0 {
		t.Errorf("expected set to be empty, but got %d elements", set.Len())
	}
}

// TestConcurrentSafe_Contains ensures `Contains` works with concurrent operations.
func TestConcurrentSafe_Contains(t *testing.T) {
	set := set.NewConcurrentSafe[int]()

	// Add elements to the set:
	for i := 0; i < 1000; i++ {
		set.Add(i)
	}

	var wg sync.WaitGroup

	wg.Add(routines)

	// Concurrently check for the presence of elements in the set.
	for i := 0; i < routines; i++ {
		go func(i int) {
			defer wg.Done()

			if !set.Contains(i) {
				t.Errorf("expected set to contain %d", i)
			}
		}(i)
	}

	wg.Wait()
}

// TestConcurrentSafe_AddRemoveContains tests a mix of operations concurrently.
func TestConcurrentSafe_AddRemoveContains(t *testing.T) {
	set := set.NewConcurrentSafe[int]()

	var wg sync.WaitGroup

	wg.Add(routines * 3) // We'll perform `Add`, `Remove`, and `Contains` operations concurrently.

	// Concurrently add elements:
	for i := 0; i < routines; i++ {
		go func(i int) {
			defer wg.Done()

			set.Add(i)
		}(i)
	}

	// Concurrently remove elements:
	for i := 0; i < routines; i++ {
		go func(i int) {
			defer wg.Done()

			set.Remove(i)
		}(i)
	}

	// Concurrently check for the presence of elements:
	for i := 0; i < routines; i++ {
		go func(i int) {
			defer wg.Done()

			set.Contains(i)
		}(i)
	}

	wg.Wait()

	// We should have a consistent count, because we are adding and removing the same elements.
	if set.Len() < 0 || set.Len() > routines {
		t.Errorf("set length should be between 0 and %d, but got %d", routines, set.Len())
	}
}
