package aoc2019

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	DayMap["1"] = Day1{}
}

type Day1 struct {
}

func (d Day1) Part1(input string) {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		n, _ := strconv.Atoi(line)
		sum += n/3 - 2
	}
	fmt.Println("part1", sum)
}

func (d Day1) Part2(input string) {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		mass, _ := strconv.Atoi(line)
		sum += fuel(mass)
	}
	fmt.Println("part2", sum)
}

func fuel(mass int) int {
	f := mass/3 - 2
	if f < 0 {
		return 0
	}
	return f + fuel(f)
}
