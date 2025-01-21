package aoc2019

import (
	"strconv"
	"strings"
)

func init() {
	Days["5"] = Day5{
		`1002,4,3,4,33`,
	}
}

type Day5 struct {
	example string
}

func (d Day5) Part1(input string) any {
	var memory []int
	for _, parts := range strings.Split(input, ",") {
		n, _ := strconv.Atoi(parts)
		memory = append(memory, n)
	}
	c := NewIntcodeComputer(memory, []int{1})
	for c.Step() {
	}
	return c.output[len(c.output)-1]
}

func (d Day5) Part2(input string) any {
	var memory []int
	for _, parts := range strings.Split(input, ",") {
		n, _ := strconv.Atoi(parts)
		memory = append(memory, n)
	}
	c := NewIntcodeComputer(memory, []int{5})
	for c.Step() {
	}
	return c.output[len(c.output)-1]
}
