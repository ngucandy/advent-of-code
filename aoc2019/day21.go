package aoc2019

import "fmt"

func init() {
	Days["21"] = &Day21{}
}

type Day21 struct {
	eg1, eg2 string
}

func (d Day21) Part1(input string) any {
	program := `NOT A J
NOT C T
OR T J
AND D J
WALK
`
	var in []int
	for _, ch := range program {
		in = append(in, int(ch))
	}
	c := NewIntcodeComputer(ParseIntcodeProgram(input), in)
	for c.Step() {
	}
	for _, n := range c.output {
		if n > 127 {
			return n
		}
		fmt.Printf("%c", n)
	}
	return "no answer yet"
}

func (d Day21) Part2(input string) any {
	program := `NOT A J
AND D J
NOT C T
AND D T
AND H T
OR T J
NOT B T
AND D T
OR T J
RUN
`
	var in []int
	for _, ch := range program {
		in = append(in, int(ch))
	}
	c := NewIntcodeComputer(ParseIntcodeProgram(input), in)
	for c.Step() {
	}
	for _, n := range c.output {
		if n > 127 {
			return n
		}
		fmt.Printf("%c", n)
	}
	return "no answer yet"
}
