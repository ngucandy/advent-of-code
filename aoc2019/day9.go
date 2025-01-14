package aoc2019

import (
	"fmt"
)

func init() {
	DayMap["9"] = Day9{
		`109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99`,
		`1102,34915192,34915192,7,4,7,99,0`,
		`104,1125899906842624,99`,
	}
}

type Day9 struct {
	example1, example2, example3 string
}

func (d Day9) Part1(input string) {
	c := NewIntcodeComputer(ParseIntcodeProgram(input), []int{1})
	for c.Step() {
	}
	fmt.Println("part1", c.output)
}

func (d Day9) Part2(input string) {
	c := NewIntcodeComputer(ParseIntcodeProgram(input), []int{2})
	for c.Step() {
	}
	fmt.Println("part2", c.output)
}
