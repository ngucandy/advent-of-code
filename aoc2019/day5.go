package aoc2019

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	DayMap["5"] = Day5{
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
	c := NewIntcodeComputer(memory, 1)
	for c.Step() {
	}
	fmt.Println("part1", c.output[len(c.output)-1])
}

func (d Day5) Part2(input string) {
}

type IntcodeComputer struct {
	memory []int
	input  int
	output []int
	ip     int
}

func NewIntcodeComputer(memory []int, input int) *IntcodeComputer {
	c := &IntcodeComputer{
		memory: memory,
		input:  input,
		output: make([]int, 0),
		ip:     0,
	}
	return c
}

func (c *IntcodeComputer) Step() bool {
	var opcode int
	pmodes := make(map[int]int)

	opcode = c.memory[c.ip] % 100
	if opcode == 99 {
		return false
	}
	for i, modes := 0, c.memory[c.ip]/100; modes > 0; i, modes = i+1, modes/10 {
		pmode := modes % 10
		pmodes[i] = pmode
	}

	switch opcode {
	case 1, 2:
		var operands []int
		for i := range 2 {
			param := c.memory[c.ip+i+1]
			if pmode, exists := pmodes[i]; exists && pmode == 1 {
				operands = append(operands, param)
			} else {
				operands = append(operands, c.memory[param])
			}
		}
		param3 := c.memory[c.ip+3]
		if opcode == 1 {
			c.memory[param3] = operands[0] + operands[1]
		} else {
			c.memory[param3] = operands[0] * operands[1]
		}
		c.ip += 4
	case 3:
		param := c.memory[c.ip+1]
		c.memory[param] = c.input
		c.ip += 2
	case 4:
		param := c.memory[c.ip+1]
		c.output = append(c.output, c.memory[param])
		c.ip += 2
	default:
		panic(fmt.Sprintf("unknown opcode at ip %d: %d", c.ip, opcode))
	}
	return true
}
