package aoc2024

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/ngucandy/advent-of-code/internal/helpers"
)

func init() {
	Days["21"] = Day21{
		numpadGraph: map[rune]string{
			'A': "03",
			'0': "2A",
			'1': "42",
			'2': "1530",
			'3': "26A",
			'4': "751",
			'5': "4862",
			'6': "593",
			'7': "84",
			'8': "795",
			'9': "86",
		},

		numpadDirections: map[string]rune{
			"A0": '<',
			"A3": '^',
			"02": '^',
			"0A": '>',
			"14": '^',
			"12": '>',
			"21": '<',
			"25": '^',
			"23": '>',
			"20": 'v',
			"32": '<',
			"36": '^',
			"3A": 'v',
			"47": '^',
			"45": '>',
			"41": 'v',
			"54": '<',
			"58": '^',
			"56": '>',
			"52": 'v',
			"65": '<',
			"69": '^',
			"63": 'v',
			"78": '>',
			"74": 'v',
			"87": '<',
			"89": '>',
			"85": 'v',
			"98": '<',
			"96": 'v',
		},

		dirpadGraph: map[rune]string{
			'A': "^>",
			'^': "vA",
			'v': "<^>",
			'<': "v",
			'>': "vA",
		},

		dirpadDirections: map[string]rune{
			"A^": '<',
			"A>": 'v',
			"^v": 'v',
			"^A": '>',
			"v<": '<',
			"v^": '^',
			"v>": '>',
			"<v": '>',
			">v": '<',
			">A": '^',
		},
	}
}

type Day21 struct {
	eg1, eg2         string
	numpadGraph      map[rune]string
	numpadDirections map[string]rune
	dirpadGraph      map[rune]string
	dirpadDirections map[string]rune
}

func (d Day21) Part1(input string) {
	defer helpers.TrackTime(time.Now())
	total := 0
	numpadPaths := make(map[string][]string)
	for _, r1 := range "A0123456789" {
		for _, r2 := range "A0123456789" {
			numpadPaths[string(r1)+string(r2)] = d.shortestPaths(r1, r2, d.numpadGraph, d.numpadDirections)
		}
	}

	dirpadPaths := make(map[string][]string)
	for _, r1 := range "A^<v>" {
		for _, r2 := range "A^<v>" {
			dirpadPaths[string(r1)+string(r2)] = d.shortestPaths(r1, r2, d.dirpadGraph, d.dirpadDirections)
		}
	}

	for _, sequence := range strings.Split(input, "\n") {
		current := 'A'
		output := make([][]string, 0)
		for _, next := range sequence {
			output = append(output, numpadPaths[string(current)+string(next)])
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
					output = append(output, dirpadPaths[string(current)+string(next)])
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
	numpadPaths := make(map[string][]string)
	for _, r1 := range "A0123456789" {
		for _, r2 := range "A0123456789" {
			numpadPaths[string(r1)+string(r2)] = d.shortestPaths(r1, r2, d.numpadGraph, d.numpadDirections)
		}
	}

	dirpadPaths := make(map[string][]string)
	for _, r1 := range "A^<v>" {
		for _, r2 := range "A^<v>" {
			dirpadPaths[string(r1)+string(r2)] = d.shortestPaths(r1, r2, d.dirpadGraph, d.dirpadDirections)
		}
	}

	cache := make(map[[3]int]int)
	for _, numpadSequence := range strings.Split(input, "\n") {
		current := 'A'
		output := make([][]string, 0)
		for _, next := range numpadSequence {
			output = append(output, numpadPaths[string(current)+string(next)])
			current = next
		}
		dirpadSequences := make([]string, 0)
		combos := helpers.CartesianProduct(output)
		for _, combo := range combos {
			dirpadSequences = append(dirpadSequences, strings.Join(combo, ""))
		}

		shortest := math.MaxInt
		for _, seq := range dirpadSequences {
			length := 0
			current = 'A'
			for _, next := range seq {
				length += d.shortestLength(current, next, 25, dirpadPaths, cache)
				current = next
			}
			if length > shortest {
				continue
			}
			shortest = length
		}
		n, _ := strconv.Atoi(numpadSequence[:len(numpadSequence)-1])
		complexity := n * shortest
		total += complexity

	}
	fmt.Println("part2", total)
}

func (d Day21) shortestPaths(start, end rune, graph map[rune]string, directions map[string]rune) []string {
	queue := make([]string, 0)
	seen := make(map[rune]int)
	queue = append(queue, string(start))
	paths := make([]string, 0)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		currentChar := rune(current[len(current)-1])
		if seenLength, ok := seen[currentChar]; ok && seenLength < len(current) {
			continue
		}
		seen[currentChar] = len(current)
		if currentChar == end {
			paths = append(paths, current[:len(current)-1]+"A")
			continue
		}
		neighbors := graph[currentChar]
		for _, n := range neighbors {
			queue = append(queue, current[:len(current)-1]+string(directions[string(currentChar)+string(n)])+string(n))
		}
	}
	return paths
}

func (d Day21) shortestLength(start, end rune, depth int, paths map[string][]string, cache map[[3]int]int) int {
	if depth == 1 {
		return len(paths[string(start)+string(end)][0])
	}
	k := [3]int{int(start), int(end), depth}
	if l, ok := cache[k]; ok {
		return l
	}

	shortest := math.MaxInt
	for _, path := range paths[string(start)+string(end)] {
		length := 0
		current := 'A'
		for _, next := range path {
			length += d.shortestLength(current, next, depth-1, paths, cache)
			current = next
		}
		if length > shortest {
			continue
		}
		shortest = length
	}
	cache[k] = shortest
	return shortest
}
