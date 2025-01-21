package aoc2024

import (
	"fmt"
	"github.com/ngucandy/advent-of-code/internal/helpers"
	"slices"
	"strconv"
	"strings"
)

func init() {
	Days["17"] = Day17{
		eg1: `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
`,
	}
}

type Day17 struct {
	eg1, eg2 string
}

func (d Day17) Part1(input string) any {
	sections := strings.Split(input, "\n\n")

	var a, b, c int
	registers := strings.Split(sections[0], "\n")
	_, _ = fmt.Sscanf(registers[0], "Register A: %d", &a)
	_, _ = fmt.Sscanf(registers[0], "Register B: %d", &b)
	_, _ = fmt.Sscanf(registers[0], "Register C: %d", &c)

	var program []int
	for _, part := range strings.Split(sections[1][9:], ",") {
		n, _ := strconv.Atoi(part)
		program = append(program, n)
	}

	output := d.run(program, a, b, c)
	return helpers.Join(output, ",")
}

func (d Day17) Part2(input string) any {
	sections := strings.Split(input, "\n\n")

	var program []int
	for _, part := range strings.Split(sections[1][9:], ",") {
		n, _ := strconv.Atoi(part)
		program = append(program, n)
	}

	// The input program's logic performs a loop that ultimately decrements the
	// value in `a` until `a` reaches `0` and the program stops. Each time
	// through loop the program outputs a single value.
	//
	// The output value of the program is determined by each 3-bit section of
	// `a`. The input program has 16 values, so `a` will need to be a 48-bit
	// integer.

	// This approach tries to build up `a` to its final value by starting small
	// and incrementing `a` until that produces the last value in the program.
	// It saves this value by shifting `a` left 3 bits and then continues
	// incrementing `a` until that produces the second to last 2 program values.
	// This continues until all program values are produced.
	a := 0
	for i := len(program) - 1; i >= 0; i-- {
		a <<= 3
		for {
			output := d.run(program, a, 0, 0)
			if slices.Equal(program[i:], output) {
				break
			}
			a++
		}
	}
	return fmt.Sprintf("%x", a)
}

func (d Day17) run(program []int, a int, b int, c int) []int {
	var output []int
	for ip := 0; ip < len(program); {
		op := program[ip]
		operand := program[ip+1]
		switch op {
		case 0: // adv
			a = a / (1 << d.combo(operand, a, b, c))
		case 1: // bxl
			b = b ^ operand
		case 2: // bst
			b = d.combo(operand, a, b, c) % 8
		case 3: // jnz
			if a == 0 {
				break
			}
			ip = operand
			continue
		case 4: // bxc
			b = b ^ c
		case 5: // out
			output = append(output, d.combo(operand, a, b, c)%8)
		case 6: // bdv
			b = a / (1 << d.combo(operand, a, b, c))
		case 7: // cdv
			c = a / (1 << d.combo(operand, a, b, c))
		}
		ip += 2
	}
	return output
}

func (d Day17) combo(o, a, b, c int) int {
	switch o {
	case 0, 1, 2, 3:
		return o
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	default:
		panic(fmt.Sprintf("unknown combo operand: %d", o))
	}
}
