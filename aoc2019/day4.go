package aoc2019

import (
	"strconv"
	"strings"
)

func init() {
	Days["4"] = Day4{}
}

type Day4 struct {
}

func (d Day4) Part1(input string) any {
	parts := strings.Split(strings.TrimSpace(input), "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	count := 0
outer:
	for i := start; i <= end; i++ {
		counts := make(map[rune]int)
		pwd := strconv.Itoa(i)
		for _, char := range pwd {
			counts[char]++
		}
		if len(counts) == 6 {
			continue
		}
		for j := 1; j < len(pwd); j++ {
			if pwd[j] < pwd[j-1] {
				continue outer
			}
		}
		count++

	}
	return count
}

func (d Day4) Part2(input string) any {
	parts := strings.Split(strings.TrimSpace(input), "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	count := 0
outer:
	for i := start; i <= end; i++ {
		counts := make(map[rune]int)
		pwd := strconv.Itoa(i)
		for _, char := range pwd {
			counts[char]++
		}
		found := false
		for _, c := range counts {
			if c == 2 {
				found = true
				break
			}
		}
		if !found {
			continue
		}
		for j := 1; j < len(pwd); j++ {
			if pwd[j] < pwd[j-1] {
				continue outer
			}
		}
		count++

	}
	return count
}
