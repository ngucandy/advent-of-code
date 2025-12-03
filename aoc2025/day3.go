package aoc2025

import (
	"strconv"
	"strings"
)

func init() {
	Days["3"] = Day3{
		example: `987654321111111
811111111111119
234234234234278
818181911112111`,
	}
}

type Day3 struct {
	example string
}

func (d Day3) Part1(input string) any {
	total := 0
	n1, n2 := -1, -1
	for _, line := range strings.Split(input, "\n") {
		for _, n := range "987654321" {
			i := strings.Index(line, string(n))
			if i == -1 || i == len(line)-1 {
				continue
			}
			n1 = int(n - '0')
			line = line[i+1:]
			break
		}

		for _, n := range "9876543210" {
			i := strings.Index(line, string(n))
			if i == -1 {
				continue
			}
			n2 = int(n - '0')
			break
		}

		total += (n1 * 10) + n2
	}
	return total
}

func (d Day3) Part2(input string) any {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		joltage, _ := strconv.Atoi(string(d.largest(line, 12)))
		sum += joltage
	}
	return sum
}

func (d Day3) largest(bank string, remaining int) []rune {
	if remaining == 0 {
		return []rune{}
	}
	pre := bank[:len(bank)-(remaining-1)]
	for _, r := range "987654321" {
		i := strings.Index(pre, string(r))
		if i == -1 {
			continue
		}
		return append([]rune{r}, d.largest(bank[i+1:], remaining-1)...)
	}
	panic("should not reach here")
}
