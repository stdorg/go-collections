package set_test

import (
	"sync"
	"testing"
	"time"

	"github.com/stdorg/go-collections/set"
)

// Constants for tuning the concurrency tests.
const (
	routines   = 750
	iterations = 100

	testDuration  = 50 * time.Millisecond // Duration for tests with continuous operations.
	clearInterval = 10 * time.Millisecond // Interval for the clear operation test.
)

// TestSafeSet_AddRemove tests the basic thread-safety of concurrent Add and Remove operations. It spawns goroutines to add elements and, simultaneously, other goroutines to remove them. The test succeeds if the set is empty at the end, proving that every Add was matched by a Remove.
func TestSafeSet_AddRemove(t *testing.T) {
	t.Parallel()

	s := set.NewSafe[int]()
	wg := sync.WaitGroup{}

	// Run Add operations in parallel.
	runInParallel(&wg, routines, func(id int) {
		s.Add(id)
	})

	// Concurrently, run Remove operations in parallel.
	runInParallel(&wg, routines, func(id int) {
		s.Remove(id)
	})

	wg.Wait()

	if s.Len() != 0 {
		t.Errorf("expected set to be empty, but got %d elements", s.Len())
	}
}

// TestSafeSet_MixedOperations is a stress test that simulates a more complex, realistic workload. It involves goroutines that add, remove, and read from the set concurrently.
func TestSafeSet_MixedOperations(t *testing.T) {
	t.Parallel()

	s := set.NewSafe[int]()
	wg := sync.WaitGroup{}

	// Pre-populate the set with initial elements.
	initialCount := routines * 2

	for i := range initialCount {
		s.Add(i)
	}

	// Group 1: Writers that add and then conditionally remove new elements.
	t.Run("AddAndRemoveWriters", func(t *testing.T) {
		runInParallel(&wg, routines, func(id int) {
			element := initialCount + id

			for j := range iterations {
				s.Add(element)

				if (j % 2) == 0 {
					s.Remove(element)
				}
			}
		})
	})

	// Group 2: Removers that target the initial set of elements.
	t.Run("InitialElementsRemovers", func(t *testing.T) {
		runInParallel(&wg, routines, func(id int) {
			element := id % initialCount

			for range iterations {
				s.Remove(element)
			}
		})
	})

	// Group 3: Readers that continuously check for elements and the set's length.
	t.Run("ConcurrentReaders", func(t *testing.T) {
		runInParallel(&wg, routines, func(id int) {
			elementToCheck := id % (initialCount + routines)

			for range iterations {
				s.Contains(elementToCheck)
				s.Len()
			}
		})
	})

	wg.Wait()

	if s.Len() < 0 {
		t.Errorf("set length became negative: %d", s.Len())
	}

	t.Logf("Mixed operations finished. Final set size: %d", s.Len())
}

// TestSafeSet_ReadsDuringWrites checks for race conditions between read operations (Contains, Len) and write operations (Add, Remove).
func TestSafeSet_ReadsDuringWrites(t *testing.T) {
	t.Parallel()

	s := set.NewSafe[int]()
	stop := make(chan struct{})
	wg := sync.WaitGroup{}

	// Start writer goroutines that continuously add and remove elements.
	wg.Add(routines)

	for i := range routines {
		go func(id int) {
			defer wg.Done()

			ticker := time.NewTicker(time.Millisecond)

			defer ticker.Stop()

			for {
				select {
				case <-stop:
					return
				case <-ticker.C:
					if (time.Now().UnixNano() % 2) == 0 {
						s.Add(id)
					} else {
						s.Remove(id)
					}
				}
			}
		}(i)
	}

	// Start reader goroutines that continuously perform read operations.
	wg.Add(routines)

	for i := range routines {
		go func(id int) {
			defer wg.Done()

			elementToCheck := id
			ticker := time.NewTicker(time.Millisecond)

			defer ticker.Stop()

			for {
				select {
				case <-stop:
					return
				case <-ticker.C:
					s.Contains(elementToCheck)
					s.Len()
				}
			}
		}(i)
	}

	// Let the test run for a short duration.
	time.Sleep(testDuration)

	// Signal all goroutines to stop.
	close(stop)

	wg.Wait()

	t.Log("Reads during writes test finished successfully.")
}

// TestSafeSet_ClearConcurrent tests the Clear operation's safety when other modifications are happening concurrently.
func TestSafeSet_ClearConcurrent(t *testing.T) {
	t.Parallel()

	s := set.NewSafe[int]()

	for i := range routines * 2 {
		s.Add(i)
	}

	wg := sync.WaitGroup{}

	// Start goroutines that perform mixed Add/Remove operations.
	runInParallel(&wg, routines, func(id int) {
		for j := range iterations {
			if (j % 2) == 0 {
				s.Add(id)
			} else {
				s.Remove(id)
			}
		}
	})

	// Start a separate goroutine to call Clear periodically.
	wg.Add(1)

	go func() {
		defer wg.Done()

		ticker := time.NewTicker(clearInterval)

		defer ticker.Stop()

		for range 10 {
			<-ticker.C
			s.Clear()
		}
	}()

	wg.Wait()

	if s.Len() < 0 {
		t.Errorf("set length is negative after concurrent clear: %d", s.Len())
	}

	t.Logf("Concurrent clear test finished. Final set size: %d", s.Len())
}

// TestSafeSet_SliceConcurrent tests the safety of creating a slice from the set while it is being actively modified by other goroutines.
func TestSafeSet_SliceConcurrent(t *testing.T) {
	t.Parallel()

	s := set.NewSafe[int]()
	stop := make(chan struct{})
	wg := sync.WaitGroup{}

	// Start writer goroutines.
	wg.Add(routines)

	for i := range routines {
		go func(id int) {
			defer wg.Done()

			ticker := time.NewTicker(time.Millisecond)

			defer ticker.Stop()

			for {
				select {
				case <-stop:
					return
				case <-ticker.C:
					if (time.Now().UnixNano() % 2) == 0 {
						s.Add(id)
					} else {
						s.Remove(id)
					}
				}
			}
		}(i)
	}

	// Start goroutines that create and iterate over slices of the set.
	wg.Add(routines)

	for range routines {
		go func() {
			defer wg.Done()

			ticker := time.NewTicker(time.Millisecond)

			defer ticker.Stop()

			for {
				select {
				case <-stop:
					return
				case <-ticker.C:
					// Get a snapshot and iterate over it. This tests for race conditions during the slice creation.
					slice := s.Slice()

					for range slice {
						// The purpose is just to read the slice, not use the values.
					}
				}
			}
		}()
	}

	time.Sleep(testDuration)

	close(stop)

	wg.Wait()

	t.Log("Concurrent slice test finished successfully.")
}

// runInParallel is a helper function to spawn a specified number of goroutines and execute a given action in each. It uses a WaitGroup to ensure all goroutines complete before returning.
func runInParallel(wg *sync.WaitGroup, total int, action func(id int)) {
	wg.Add(total)

	for i := range total {
		go func(id int) {
			defer wg.Done()

			action(id)
		}(i)
	}
}
