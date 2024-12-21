package main

import (
	"bufio"
	"fmt"
	"github.com/ngucandy/advent-of-code/internal/helpers"
	"log/slog"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	bytes, _ := os.ReadFile(infile)
	input := string(bytes)

	part1(input)
	part2(input)

}

var (
	numpadGraph = map[rune]string{
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
	}

	numpadDirections = map[string]rune{
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
	}

	dirpadGraph = map[rune]string{
		'A': "^>",
		'^': "vA",
		'v': "<^>",
		'<': "v",
		'>': "vA",
	}

	dirpadDirections = map[string]rune{
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
	}
)

func shortestPaths(start, end rune, graph map[rune]string, directions map[string]rune) []string {
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

func part1(input string) {
	total := 0
	numpadPaths := make(map[string][]string)
	for _, r1 := range "A0123456789" {
		for _, r2 := range "A0123456789" {
			numpadPaths[string(r1)+string(r2)] = shortestPaths(r1, r2, numpadGraph, numpadDirections)
		}
	}

	dirpadPaths := make(map[string][]string)
	for _, r1 := range "A^<v>" {
		for _, r2 := range "A^<v>" {
			dirpadPaths[string(r1)+string(r2)] = shortestPaths(r1, r2, dirpadGraph, dirpadDirections)
		}
	}

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		sequence := scanner.Text()
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
		fmt.Println(sequence, n, minLength)
	}
	slog.Info("Part 1:", "total", total)
}

func part2(input string) {
	total := 0

	slog.Info("Part 2:", "total", total)
}
