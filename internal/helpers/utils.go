package helpers

import (
	"fmt"
	"log/slog"
	"runtime"
	"strconv"
	"strings"
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

func TrackTime(start time.Time) {
	elapsed := time.Since(start)
	pc, _, _, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	slog.Info("Time:", "took", elapsed, "func", f.Name())
}

func PrintGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func CloneGrid[E any](grid [][]E) [][]E {
	var dst [][]E
	for _, row := range grid {
		dst = append(dst, append([]E{}, row...))
	}
	return dst
}

func PrintIntGrid(grid [][]int) {
	for _, row := range grid {
		s := ""
		for _, cell := range row {
			s += strconv.Itoa(cell)
		}
		fmt.Println(s)
	}
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Join[E any](a []E, sep string) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), sep), "[]")
}
