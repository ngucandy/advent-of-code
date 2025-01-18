package aoc2019

import (
	"fmt"
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

func (d Day5) Part1(input string) {
	var memory []int
	for _, parts := range strings.Split(input, ",") {
		n, _ := strconv.Atoi(parts)
		memory = append(memory, n)
	}
	c := NewIntcodeComputer(memory, []int{1})
	for c.Step() {
	}
	fmt.Println("part1", c.output[len(c.output)-1])
}

func (d Day5) Part2(input string) {
	var memory []int
	for _, parts := range strings.Split(input, ",") {
		n, _ := strconv.Atoi(parts)
		memory = append(memory, n)
	}
	c := NewIntcodeComputer(memory, []int{5})
	for c.Step() {
	}
	fmt.Println("part2", c.output[len(c.output)-1])
}
