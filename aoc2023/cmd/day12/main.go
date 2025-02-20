package main

import (
	"github.com/ngucandy/advent-of-code/internal/helpers"
	"log/slog"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	bytes, _ := os.ReadFile(infile)
	input := strings.ReplaceAll(string(bytes), "\r\n", "\n")

	part1(input)
	part2(input)
}

func part1(input string) {
	total := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		springs := parts[0]
		parts = strings.Split(parts[1], ",")
		var damaged []int
		for _, part := range parts {
			n, _ := strconv.Atoi(part)
			damaged = append(damaged, n)
		}
		total += arrangements(springs, damaged)
	}
	slog.Info("Part 1:", "total", total)
}

func arrangements(springs string, damaged []int) int {
	unknownCount := strings.Count(springs, "?")
	combos := helpers.CartesianProductN([]rune("#."), unknownCount)
	matched := 0
	for _, combo := range combos {
		var arrangement string
		for _, spring := range springs {
			if spring == '?' {
				arrangement += string(combo[0])
				combo = combo[1:]
				continue
			}
			arrangement += string(spring)
		}
		if isMatch(arrangement, damaged) {
			matched++
		}
	}
	return matched
}

var reDamaged = regexp.MustCompile(`#+`)

func isMatch(springs string, damaged []int) bool {
	matches := reDamaged.FindAllString(springs, -1)
	damagedMatch := make([]int, len(matches))
	for i, match := range matches {
		damagedMatch[i] = len(match)
	}
	return slices.Equal(damagedMatch, damaged)
}

func part2(input string) {

}
