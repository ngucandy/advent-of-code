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
		total += arragnements(springs, damaged)
	}
	slog.Info("Part 1:", "total", total)
}

func arragnements(springs string, damaged []int) int {
	unknownCount := strings.Count(springs, "?")
	combos := helpers.CartesianProductN([]rune("#."), unknownCount)
	matched := 0
	for _, combo := range combos {
		var arrangment string
		for _, spring := range springs {
			if spring == '?' {
				arrangment += string(combo[0])
				combo = combo[1:]
				continue
			}
			arrangment += string(spring)
		}
		if isMatch(arrangment, damaged) {
			matched++
		}
	}
	return matched
}

func isMatch(springs string, damaged []int) bool {
	re := regexp.MustCompile(`#+`)
	matches := re.FindAllString(springs, -1)
	damagedMatch := make([]int, len(matches))
	for i, match := range matches {
		damagedMatch[i] = len(match)
	}
	return slices.Equal(damagedMatch, damaged)
}

func part2(input string) {

}
