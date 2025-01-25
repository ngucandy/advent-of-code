package aoc2019

import (
	"strconv"
	"strings"
)

func init() {
	Days["1"] = Day1{}
}

type Day1 struct {
}

func (d Day1) Part1(input string) any {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		n, _ := strconv.Atoi(line)
		sum += n/3 - 2
	}
	return sum
}

func (d Day1) Part2(input string) any {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		mass, _ := strconv.Atoi(line)
		sum += d.fuel(mass)
	}
	return sum
}

func (d Day1) fuel(mass int) int {
	f := mass/3 - 2
	if f < 0 {
		return 0
	}
	return f + d.fuel(f)
}
