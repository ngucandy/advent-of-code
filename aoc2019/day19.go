package aoc2019

import (
	"slices"
)

func init() {
	Days["19"] = &Day19{}
}

type Day19 struct {
	eg1, eg2 string
}

func (d Day19) Part1(input string) any {
	program := ParseIntcodeProgram(input)
	count := 0
	n := 50
	for y := range n {
		for x := range n {
			c := NewIntcodeComputer(slices.Clone(program), []int{x, y})
			for len(c.output) == 0 {
				c.Step()
			}
			if c.output[0] == 1 {
				count++
			}
		}
	}
	return count
}

func (d Day19) Part2(input string) any {
	return "no answer yet"
}
