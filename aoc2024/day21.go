package aoc2024

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/ngucandy/advent-of-code/internal/helpers"
)

func init() {
	Days["21"] = Day21{
		numpad: [][]rune{
			[]rune("789"),
			[]rune("456"),
			[]rune("123"),
			[]rune("#0A"),
		},

		npButtons: map[rune][2]int{
			'7': {0, 0},
			'8': {0, 1},
			'9': {0, 2},
			'4': {1, 0},
			'5': {1, 1},
			'6': {1, 2},
			'1': {2, 0},
			'2': {2, 1},
			'3': {2, 2},
			'#': {3, 0},
			'0': {3, 1},
			'A': {3, 2},
		},

		dirpad: [][]rune{
			[]rune("#^A"),
			[]rune("<v>"),
		},

		dpButtons: map[rune][2]int{
			'#': {0, 0},
			'^': {0, 1},
			'A': {0, 2},
			'<': {1, 0},
			'v': {1, 1},
			'>': {1, 2},
		},

		directions: map[rune][2]int{
			'^': {-1, 0},
			'v': {1, 0},
			'<': {0, -1},
			'>': {0, 1},
		},

		cache: make(map[[2]rune][]string),
	}
}

type Day21 struct {
	eg1, eg2       string
	numpad, dirpad [][]rune
	npButtons      map[rune][2]int
	dpButtons      map[rune][2]int
	directions     map[rune][2]int
	cache          map[[2]rune][]string
}

func (d Day21) paths(s, e rune, grid [][]rune, buttons map[rune][2]int) []string {
	if paths, exists := d.cache[[2]rune{s, e}]; exists {
		return paths
	}

	var paths []string
	seen := make(map[rune]int)
	q := [][]rune{{s}, {}}
	for len(q) > 0 {
		b := q[0][0]
		path := q[1]
		q = q[2:]

		if seenPath, exists := seen[b]; exists && seenPath < len(path) {
			continue
		}
		seen[b] = len(path)

		if b == e {
			paths = append(paths, string(path)+"A")
		}

		r, c := buttons[b][0], buttons[b][1]
		for arrow, dir := range d.directions {
			nr, nc := r+dir[0], c+dir[1]
			if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) || grid[nr][nc] == '#' {
				continue
			}
			nb := grid[nr][nc]
			q = append(q, []rune{nb}, append(slices.Clone(path), arrow))
		}
	}
	slices.Sort(paths)
	d.cache[[2]rune{s, e}] = paths
	return paths
}

func (d Day21) Part1(input string) {
	defer helpers.TrackTime(time.Now())
	total := 0

	for _, sequence := range strings.Split(input, "\n") {
		current := 'A'
		output := make([][]string, 0)
		for _, next := range sequence {
			output = append(output, d.paths(current, next, d.numpad, d.npButtons))
			current = next
		}
		nextSequences := make([]string, 0)
		combos := helpers.CartesianProduct(output)
		for _, combo := range combos {
			nextSequences = append(nextSequences, strings.Join(combo, ""))
		}

		for range 2 {
			dirSequences := nextSequences
			nextSequences = make([]string, 0)
			for _, dirSequence := range dirSequences {
				current = 'A'
				output = make([][]string, 0)
				for _, next := range dirSequence {
					output = append(output, d.paths(current, next, d.dirpad, d.dpButtons))
					current = next
				}
				combos = helpers.CartesianProduct(output)
				for _, combo := range combos {
					nextSequences = append(nextSequences, strings.Join(combo, ""))
				}
			}
		}

		minLength := math.MaxInt
		for _, seq := range nextSequences {
			minLength = min(minLength, len(seq))
		}
		n, _ := strconv.Atoi(sequence[:len(sequence)-1])
		complexity := n * minLength
		total += complexity
	}
	fmt.Println("part1", total)

}

func (d Day21) Part2(input string) {
	defer helpers.TrackTime(time.Now())
	total := 0

	cache := make(map[[3]int]int)
	for _, numpadSequence := range strings.Split(input, "\n") {
		current := 'A'
		output := make([][]string, 0)
		for _, next := range numpadSequence {
			output = append(output, d.paths(current, next, d.numpad, d.npButtons))
			current = next
		}
		dirpadSequences := make([]string, 0)
		combos := helpers.CartesianProduct(output)
		for _, combo := range combos {
			dirpadSequences = append(dirpadSequences, strings.Join(combo, ""))
		}

		var lengths []int
		for _, seq := range dirpadSequences {
			length := 0
			current = 'A'
			for _, next := range seq {
				length += d.shortestLength(current, next, 25, cache)
				current = next
			}
			lengths = append(lengths, length)
		}
		slices.Sort(lengths)
		n, _ := strconv.Atoi(numpadSequence[:len(numpadSequence)-1])
		complexity := n * lengths[0]
		total += complexity

	}
	fmt.Println("part2", total)
}

func (d Day21) shortestLength(start, end rune, depth int, cache map[[3]int]int) int {
	if depth == 1 {
		return len(d.paths(start, end, d.dirpad, d.dpButtons)[0])
	}
	k := [3]int{int(start), int(end), depth}
	if l, ok := cache[k]; ok {
		return l
	}

	var lengths []int
	for _, path := range d.paths(start, end, d.dirpad, d.dpButtons) {
		length := 0
		current := 'A'
		for _, next := range path {
			length += d.shortestLength(current, next, depth-1, cache)
			current = next
		}
		lengths = append(lengths, length)
	}
	slices.Sort(lengths)
	cache[k] = lengths[0]
	return lengths[0]
}
