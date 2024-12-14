package helpers

import (
	"log/slog"
	"time"
)

// CartesianProduct computes the Cartesian product of `sets`, where each set is represented as a slice.
func CartesianProduct[T any](sets [][]T) (result [][]T) {
	// Base case: If there are no sets, return an empty slice
	if len(sets) == 0 {
		return [][]T{{}}
	}

	// Take the first set and recursively get the Cartesian product of the rest
	firstSet := sets[0]
	remainingSets := sets[1:]

	// Recursively call CartesianProduct on the remaining sets
	remainingProduct := CartesianProduct[T](remainingSets)

	// Combine elements of the first set with the result of the recursive call
	for _, firstElem := range firstSet {
		for _, remainingElems := range remainingProduct {
			// Create a new combination by appending `firstElem` to each of the remaining products
			result = append(result, append([]T{firstElem}, remainingElems...))
		}
	}

	return result
}

// CartesianProductN computes the Cartesian product of `set` repeated `n` times
func CartesianProductN[T any](set []T, n int) (result [][]T) {
	sets := make([][]T, 0)
	for range n {
		sets = append(sets, set)
	}
	return CartesianProduct(sets)
}

func TrackTime(start time.Time, name string) {
	elapsed := time.Since(start)
	slog.Info("Time:", "name", name, "took", elapsed)
}
