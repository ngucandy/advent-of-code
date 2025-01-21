package aoc2024

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/ngucandy/advent-of-code/internal/helpers"
)

func init() {
	d := &Day21{
		numpad: [][]rune{
			[]rune("789"),
			[]rune("456"),
			[]rune("123"),
			[]rune("#0A"),
		},

		dirpad: [][]rune{
			[]rune("#^A"),
			[]rune("<v>"),
		},

		directions: map[rune][2]int{
			'^': {-1, 0},
			'v': {1, 0},
			'<': {0, -1},
			'>': {0, 1},
		},

		npButtons: make(map[rune][2]int),
		dpButtons: make(map[rune][2]int),
		cache:     make(map[[2]rune][]string),
	}

	for r, row := range d.numpad {
		for c, ch := range row {
			d.npButtons[ch] = [2]int{r, c}
		}
	}

	for r, row := range d.dirpad {
		for c, ch := range row {
			d.dpButtons[ch] = [2]int{r, c}
		}
	}
	Days["21"] = d
}

type Day21 struct {
	eg1, eg2       string
	numpad, dirpad [][]rune
	npButtons      map[rune][2]int
	dpButtons      map[rune][2]int
	directions     map[rune][2]int
	cache          map[[2]rune][]string
}

func (d Day21) paths(s, e rune, keypad [][]rune, locations map[rune][2]int) []string {
	if paths, exists := d.cache[[2]rune{s, e}]; exists {
		return paths
	}

	// bfs to find the path from `s` to `e` on the given keypad
	var paths []string
	seen := make(map[rune]int)
	// queue holds button name and the path used to get there
	q := [][]rune{{s}, {}}
	for len(q) > 0 {
		b := q[0][0]
		path := q[1]
		q = q[2:]

		// we want all paths with the same length, so only skip if we've seen
		// this button before with a shorter path
		if seenPath, exists := seen[b]; exists && seenPath < len(path) {
			continue
		}
		seen[b] = len(path)

		if b == e {
			// we've reached the desired button so terminate path with an 'A'
			paths = append(paths, string(path)+"A")
		}

		r, c := locations[b][0], locations[b][1]
		for arrow, dir := range d.directions {
			nr, nc := r+dir[0], c+dir[1]
			if nr < 0 || nr >= len(keypad) || nc < 0 || nc >= len(keypad[0]) || keypad[nr][nc] == '#' {
				continue
			}
			nb := keypad[nr][nc]
			q = append(q, []rune{nb}, append(slices.Clone(path), arrow))
		}
	}
	// order of the path doesn't really matter, but we sort so that they're
	// always returned in a consistent order to make testing easier
	slices.Sort(paths)
	d.cache[[2]rune{s, e}] = paths
	return paths
}

func (d Day21) Part1(input string) {
	defer helpers.TrackTime(time.Now())
	total := 0

	for _, line := range strings.Split(input, "\n") {
		seq := []rune("A" + line)

		// left to right, for each pair of neighboring buttons in `seq`,
		// `possibilities` contains a slice of possible moves to go from the
		// left button to its right neighbor
		var possibilities [][]string
		for i := range seq[:len(seq)-1] {
			possibilities = append(possibilities, d.paths(seq[i], seq[i+1], d.numpad, d.npButtons))
		}

		// the cartesian product of `possibilities` will produce a list of
		// candidate dirpad movements
		var candidates []string
		for _, possibility := range helpers.CartesianProduct(possibilities) {
			candidates = append(candidates, strings.Join(possibility, ""))
		}

		// repeat this process using the dirpad
		next := candidates
		for range 2 { // 2 dirpads are controlled by robots
			candidates = make([]string, 0)
			for _, candidate := range next {
				seq = []rune("A" + candidate)
				possibilities = make([][]string, 0)
				for i := range seq[:len(seq)-1] {
					possibilities = append(possibilities, d.paths(seq[i], seq[i+1], d.dirpad, d.dpButtons))
				}
				for _, possibility := range helpers.CartesianProduct(possibilities) {
					candidates = append(candidates, strings.Join(possibility, ""))
				}
				// sort the candidates by their length
				slices.SortFunc(candidates, func(a, b string) int {
					return cmp.Compare(len(a), len(b))
				})
				// only keep candidates with the smallest length
				minl := len(candidates[0])
				if i := slices.IndexFunc(candidates, func(s string) bool {
					return len(s) > minl
				}); i != -1 {
					candidates = candidates[:i]
				}
			}
			next = candidates
		}
		// compute `complexity` as length of shortest sequence * numeric part of keypad sequence
		n, _ := strconv.Atoi(line[:len(line)-1])
		complexity := n * len(candidates[0])
		total += complexity
	}
	fmt.Println("part1", total)

}

func (d Day21) Part2(input string) {
	defer helpers.TrackTime(time.Now())
	total := 0

	cache := make(map[[3]int]int)
	for _, line := range strings.Split(input, "\n") {
		seq := []rune("A" + line)

		// left to right, for each pair of neighboring buttons in `seq`,
		// `possibilities` contains a slice of possible moves to go from the
		// left button to its right neighbor
		var possibilities [][]string
		for i := range seq[:len(seq)-1] {
			possibilities = append(possibilities, d.paths(seq[i], seq[i+1], d.numpad, d.npButtons))
		}

		// the cartesian product of `possibilities` will produce a list of
		// candidate dirpad movements
		var candidates []string
		for _, possibility := range helpers.CartesianProduct(possibilities) {
			candidates = append(candidates, strings.Join(possibility, ""))
		}

		var lengths []int
		for _, candidate := range candidates {
			length := 0
			seq = []rune("A" + candidate)
			for i := range seq[:len(seq)-1] {
				length += d.shortestLength(seq[i], seq[i+1], 25, cache)
			}
			lengths = append(lengths, length)
		}

		slices.Sort(lengths)
		n, _ := strconv.Atoi(line[:len(line)-1])
		complexity := n * lengths[0]
		total += complexity
	}
	fmt.Println("part2", total)
}

func (d Day21) shortestLength(s, e rune, depth int, cache map[[3]int]int) int {
	if depth == 1 {
		return len(d.paths(s, e, d.dirpad, d.dpButtons)[0])
	}
	key := [3]int{int(s), int(e), depth}
	if l, ok := cache[key]; ok {
		return l
	}

	var lengths []int
	for _, path := range d.paths(s, e, d.dirpad, d.dpButtons) {
		length := 0
		seq := []rune("A" + path)
		for i := range seq[:len(seq)-1] {
			length += d.shortestLength(seq[i], seq[i+1], depth-1, cache)
		}
		lengths = append(lengths, length)
	}
	slices.Sort(lengths)
	cache[key] = lengths[0]
	return lengths[0]
}
