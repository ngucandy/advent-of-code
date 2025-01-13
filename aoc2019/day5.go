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
	var memory []int
	for _, parts := range strings.Split(input, ",") {
		n, _ := strconv.Atoi(parts)
		memory = append(memory, n)
	}
	c := NewIntcodeComputer(memory, 5)
	for c.Step() {
	}
	fmt.Println("part2", c.output[len(c.output)-1])
}

type IntcodeComputer struct {
	memory []int
	input  int
	output []int
	ip     int
	opfns  map[int]func(map[int]int)
}

func NewIntcodeComputer(memory []int, input int) *IntcodeComputer {
	c := &IntcodeComputer{
		memory: memory,
		input:  input,
		output: make([]int, 0),
		ip:     0,
		opfns:  make(map[int]func(map[int]int)),
	}
	c.opfns[1] = c.opcode1
	c.opfns[2] = c.opcode2
	c.opfns[3] = c.opcode3
	c.opfns[4] = c.opcode4
	c.opfns[5] = c.opcode5
	c.opfns[6] = c.opcode6
	c.opfns[7] = c.opcode7
	c.opfns[8] = c.opcode8
	return c
}

func (c *IntcodeComputer) opcode1(pmodes map[int]int) {
	operands := c.readParams(2, pmodes)
	param3 := c.memory[c.ip+3]
	c.memory[param3] = operands[0] + operands[1]
	c.ip += 4
}

func (c *IntcodeComputer) opcode2(pmodes map[int]int) {
	operands := c.readParams(2, pmodes)
	param3 := c.memory[c.ip+3]
	c.memory[param3] = operands[0] * operands[1]
	c.ip += 4
}

// Opcode 3 takes a single integer as input and saves it to the position
// given by its only parameter. For example, the instruction 3,50 would
// take an input value and store it at address 50.
func (c *IntcodeComputer) opcode3(_ map[int]int) {
	param := c.memory[c.ip+1]
	c.memory[param] = c.input
	c.ip += 2
}

// Opcode 4 outputs the value of its only parameter. For example, the
// instruction 4,50 would output the value at address 50.
func (c *IntcodeComputer) opcode4(_ map[int]int) {
	param := c.memory[c.ip+1]
	c.output = append(c.output, c.memory[param])
	c.ip += 2
}

// Opcode 5 is jump-if-true: if the first parameter is non-zero, it
// sets the instruction pointer to the value from the second parameter.
// Otherwise, it does nothing.
func (c *IntcodeComputer) opcode5(pmodes map[int]int) {
	params := c.readParams(2, pmodes)
	if params[0] != 0 {
		c.ip = params[1]
	} else {
		c.ip += 3
	}
}

// Opcode 6 is jump-if-false: if the first parameter is zero, it sets
// the instruction pointer to the value from the second parameter.
// Otherwise, it does nothing.
func (c *IntcodeComputer) opcode6(pmodes map[int]int) {
	params := c.readParams(2, pmodes)
	if params[0] == 0 {
		c.ip = params[1]
	} else {
		c.ip += 3
	}
}

// Opcode 7 is less than: if the first parameter is less than the
// second parameter, it stores 1 in the position given by the third
// parameter. Otherwise, it stores 0.
func (c *IntcodeComputer) opcode7(pmodes map[int]int) {
	operands := c.readParams(2, pmodes)
	param3 := c.memory[c.ip+3]
	if operands[0] < operands[1] {
		c.memory[param3] = 1
	} else {
		c.memory[param3] = 0
	}
	c.ip += 4
}

// Opcode 8 is equals: if the first parameter is equal to the second
// parameter, it stores 1 in the position given by the third parameter.
// Otherwise, it stores 0.
func (c *IntcodeComputer) opcode8(pmodes map[int]int) {
	operands := c.readParams(2, pmodes)
	param3 := c.memory[c.ip+3]
	if operands[0] == operands[1] {
		c.memory[param3] = 1
	} else {
		c.memory[param3] = 0
	}
	c.ip += 4
}

func (c *IntcodeComputer) Step() bool {
	// opcode is one's and ten's digit, e.g., 1002 is opcode 2
	var opcode int
	opcode = c.memory[c.ip] % 100

	// opcode 99 halts program
	if opcode == 99 {
		return false
	}

	// parameter modes are 100's, 1000's, 10,000's,... digits
	// e.g., for 1002, param1 mode is 0, param2 mode is 1
	pmodes := make(map[int]int)
	for i, modes := 0, c.memory[c.ip]/100; modes > 0; i, modes = i+1, modes/10 {
		pmode := modes % 10
		pmodes[i] = pmode
	}

	if fn, exists := c.opfns[opcode]; exists {
		fn(pmodes)
	} else {
		panic(fmt.Sprintf("unknown opcode at ip %d: %d", c.ip, c.memory[c.ip]))
	}

	return true
}

func (c *IntcodeComputer) readParams(n int, pmodes map[int]int) []int {
	var params []int
	for i := range n {
		param := c.memory[c.ip+i+1]
		if pmode, exists := pmodes[i]; exists && pmode == 1 {
			params = append(params, param)
		} else {
			params = append(params, c.memory[param])
		}
	}
	return params
}
