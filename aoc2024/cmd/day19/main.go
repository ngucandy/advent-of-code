package main

import (
	"bufio"
	"log/slog"
	"os"
	"slices"
	"strings"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, _ := os.Open(infile)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	patterns := make(map[string]bool)
	designs := make([]string, 0)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	parts := strings.Split(line, ", ")
	for _, part := range parts {
		patterns[part] = true
	}
	for scanner.Scan() {
		line = scanner.Text()
		if len(line) == 0 {
			continue
		}
		designs = append(designs, line)
	}

	part1(patterns, designs)
	part2(patterns, designs)
}

func part1(patterns map[string]bool, designs []string) {
	total := 0
	cache := make(map[string]bool)
	for _, design := range designs {
		if isPossible(design, patterns, cache) {
			total++
		}
	}
	slog.Info("Part 1:", "total", total)
}

func isPossible(design string, patterns map[string]bool, cache map[string]bool) bool {
	if possible, ok := cache[design]; ok {
		return possible
	}
	if patterns[design] {
		cache[design] = true
		return true
	}
	for i := 1; i < len(design); i++ {
		if isPossible(design[:i], patterns, cache) && isPossible(design[i:], patterns, cache) {
			cache[design] = true
			return true
		}
	}
	cache[design] = false
	return false
}

func part2(patterns map[string]bool, designs []string) {
	total := 0
	cache := make(map[string][][]string)
	for _, design := range designs {
		possible := countPossible(design, patterns, cache)
		total += len(possible)
		slog.Info(design, "possible", possible)
	}
	slog.Info("Part 2:", "total", total)
}

func countPossible(design string, patterns map[string]bool, cache map[string][][]string) [][]string {
	if count, ok := cache[design]; ok {
		return count
	}
	if patterns[design] {
		cache[design] = [][]string{{design}}
		return [][]string{{design}}
	}
	ans := make([][]string, 0)
	for i := 1; i < len(design); i++ {
		left := countPossible(design[:i], patterns, cache)
		right := countPossible(design[i:], patterns, cache)
		if len(left) > 0 && len(right) > 0 {
			for _, l := range left {
				for _, r := range right {
					combined := append(l, r...)
					duplicate := false
					for _, a := range ans {
						if slices.Equal(combined, a) {
							duplicate = true
							break
						}
					}
					if !duplicate {
						ans = append(ans, combined)
					}
				}
			}
		}
	}
	cache[design] = ans
	return ans
}
