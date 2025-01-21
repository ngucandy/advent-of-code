package aoc2019

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func init() {
	Days["2"] = Day2{
		`1,9,10,3,2,3,11,0,99,30,40,50`,
	}
}

type Day2 struct {
	example string
}

func (d Day2) Part1(input string) any {
	parts := strings.Split(strings.TrimSpace(input), ",")
	var program []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		program = append(program, n)
	}
	program[1] = 12
	program[2] = 2
	d.runProgram(program)
	return program[0]
}

func (d Day2) runProgram(program []int) {
	for ip := 0; program[ip] != 99; ip += 4 {
		opcode := program[ip]
		operand1 := program[program[ip+1]]
		operand2 := program[program[ip+2]]
		switch opcode {
		case 1:
			program[program[ip+3]] = operand1 + operand2
		case 2:
			program[program[ip+3]] = operand1 * operand2
		default:
			panic(fmt.Sprintf("unknown opcode at ip %d: %d", ip, opcode))
		}
	}
}

func (d Day2) Part2(input string) any {
	parts := strings.Split(strings.TrimSpace(input), ",")
	var program []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		program = append(program, n)
	}
	program[1] = 12
	program[2] = 2

	ch := make(chan int)
	for input1 := 0; input1 <= 99; input1++ {
		for input2 := 0; input2 <= 99; input2++ {
			go func(n1, n2 int) {
				p := slices.Clone(program)
				p[1], p[2] = n1, n2
				d.runProgram(p)
				if p[0] == 19690720 {
					ch <- 100*p[1] + p[2]
				}
			}(input1, input2)
		}
	}
	return <-ch
}
