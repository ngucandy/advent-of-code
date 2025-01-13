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
	c := NewIntcodeComputer51(memory, 1)
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
	c := NewIntcodeComputer52(memory, 5)
	for c.Step() {
	}
	fmt.Println("part2", c.output[len(c.output)-1])
}

type Memory interface {
	Read(int) int
	IP() int
}
type IntcodeComputer struct {
	memory []int
	input  int
	output []int
	ip     int
}
type IntcodeComputer51 IntcodeComputer
type IntcodeComputer52 IntcodeComputer

func NewIntcodeComputer51(memory []int, input int) *IntcodeComputer51 {
	c := &IntcodeComputer51{
		memory: memory,
		input:  input,
		output: make([]int, 0),
		ip:     0,
	}
	return c
}

func (c *IntcodeComputer51) Read(i int) int {
	return c.memory[i]
}

func (c *IntcodeComputer51) IP() int {
	return c.ip
}

func (c *IntcodeComputer51) Step() bool {
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
	case 1:
		operands := readParams(c, 2, pmodes)
		param3 := c.memory[c.ip+3]
		c.memory[param3] = operands[0] + operands[1]
		c.ip += 4
	case 2:
		operands := readParams(c, 2, pmodes)
		param3 := c.memory[c.ip+3]
		c.memory[param3] = operands[0] * operands[1]
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

func NewIntcodeComputer52(memory []int, input int) *IntcodeComputer52 {
	c := &IntcodeComputer52{
		memory: memory,
		input:  input,
		output: make([]int, 0),
		ip:     0,
	}
	return c
}

func (c *IntcodeComputer52) Step() bool {
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
	case 1:
		operands := readParams(c, 2, pmodes)
		param3 := c.memory[c.ip+3]
		c.memory[param3] = operands[0] + operands[1]
		c.ip += 4
	case 2:
		operands := readParams(c, 2, pmodes)
		param3 := c.memory[c.ip+3]
		c.memory[param3] = operands[0] * operands[1]
		c.ip += 4
	case 3:
		param := c.memory[c.ip+1]
		c.memory[param] = c.input
		c.ip += 2
	case 4:
		param := c.memory[c.ip+1]
		c.output = append(c.output, c.memory[param])
		c.ip += 2
	case 5:
		params := readParams(c, 2, pmodes)
		if params[0] != 0 {
			c.ip = params[1]
		} else {
			c.ip += 3
		}
	case 6:
		params := readParams(c, 2, pmodes)
		if params[0] == 0 {
			c.ip = params[1]
		} else {
			c.ip += 3
		}
	case 7:
		operands := readParams(c, 2, pmodes)
		param3 := c.memory[c.ip+3]
		if operands[0] < operands[1] {
			c.memory[param3] = 1
		} else {
			c.memory[param3] = 0
		}
		c.ip += 4
	case 8:
		operands := readParams(c, 2, pmodes)
		param3 := c.memory[c.ip+3]
		if operands[0] == operands[1] {
			c.memory[param3] = 1
		} else {
			c.memory[param3] = 0
		}
		c.ip += 4
	default:
		panic(fmt.Sprintf("unknown opcode at ip %d: %d", c.ip, opcode))
	}
	return true
}

func (c *IntcodeComputer52) Read(i int) int {
	return c.memory[i]
}

func (c *IntcodeComputer52) IP() int {
	return c.ip
}

func readParams(m Memory, n int, pmodes map[int]int) []int {
	var params []int
	for i := range n {
		param := m.Read(m.IP() + i + 1)
		if pmode, exists := pmodes[i]; exists && pmode == 1 {
			params = append(params, param)
		} else {
			params = append(params, m.Read(param))
		}
	}
	return params
}
