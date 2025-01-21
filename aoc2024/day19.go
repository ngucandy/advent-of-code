package aoc2024

import (
	"strings"
)

func init() {
	Days["19"] = Day19{}
}

type Day19 struct {
	eg1, eg2 string
}

func (d Day19) Part1(input string) any {
	sections := strings.Split(input, "\n\n")

	patterns := make(map[string]bool)
	for _, pattern := range strings.Split(strings.TrimSpace(sections[0]), ", ") {
		patterns[pattern] = true
	}

	total := 0
	cache := make(map[string]bool)
	for _, design := range strings.Split(sections[1], "\n") {
		if d.isPossible(design, patterns, cache) {
			total++
		}
	}
	return total
}

func (d Day19) isPossible(design string, patterns map[string]bool, cache map[string]bool) bool {
	if possible, ok := cache[design]; ok {
		return possible
	}
	if patterns[design] {
		cache[design] = true
		return true
	}
	for i := 1; i < len(design); i++ {
		if d.isPossible(design[:i], patterns, cache) && d.isPossible(design[i:], patterns, cache) {
			cache[design] = true
			return true
		}
	}
	cache[design] = false
	return false
}

func (d Day19) Part2(input string) any {
	sections := strings.Split(input, "\n\n")

	patterns := make(map[string]bool)
	for _, pattern := range strings.Split(strings.TrimSpace(sections[0]), ", ") {
		patterns[pattern] = true
	}

	total := 0
	cache := make(map[string]int)
	for _, design := range strings.Split(sections[1], "\n") {
		total += d.countPossible(design, patterns, cache)
	}
	return total

}

func (d Day19) countPossible(design string, patterns map[string]bool, cache map[string]int) int {
	if count, ok := cache[design]; ok {
		return count
	}
	if len(design) == 0 {
		return 1
	}
	count := 0
	for i := 1; i <= len(design); i++ {
		if _, ok := patterns[design[:i]]; ok {
			count += d.countPossible(design[i:], patterns, cache)
		}
	}
	cache[design] = count
	return count
}
