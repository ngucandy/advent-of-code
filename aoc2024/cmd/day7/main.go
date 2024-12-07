package main

import (
	"bufio"
	"log/slog"
	"os"
	"regexp"
	"strconv"

	"github.com/mowshon/iterium"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, _ := os.Open(infile)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var equations [][]int
	rexpNums := regexp.MustCompile(`\d+`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := rexpNums.FindAllString(line, -1)
		var equation []int
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			equation = append(equation, num)
		}
		equations = append(equations, equation)
	}

	part1(equations)
	part2(equations)
}

func part1(equations [][]int) {
	validOperators := []rune{'+', '*'}
	total := 0
	for _, equation := range equations {
		nNeededOperators := len(equation[1:]) - 1
		permutations, _ := iterium.Product(validOperators, nNeededOperators).Slice()
		for _, permutation := range permutations {
			if equation[0] == compute(equation[1:], permutation) {
				total += equation[0]
				break
			}
		}
	}
	slog.Info("Part 1:", "total", total)
}

func part2(equations [][]int) {
	validOperators := []rune{'+', '*', '|'}
	total := 0
	for _, equation := range equations {
		nNeededOperators := len(equation[1:]) - 1
		permutations, _ := iterium.Product(validOperators, nNeededOperators).Slice()
		for _, permutation := range permutations {
			if equation[0] == compute(equation[1:], permutation) {
				total += equation[0]
				break
			}
		}
	}
	slog.Info("Part 2:", "total", total)
}

func compute(operands []int, operators []rune) int {
	if len(operators) != len(operands)-1 {
		slog.Error("Invalid number of operators:", "operators", operators, "operands", operands)
		panic("Invalid number of operators")
	}

	x := operands[0]
	for i, operand := range operators {
		switch operand {
		case '+':
			x += operands[i+1]
		case '*':
			x *= operands[i+1]
		case '|':
			xx, err := strconv.Atoi(strconv.Itoa(x) + strconv.Itoa(operands[i+1]))
			if err != nil {
				panic(err)
			}
			x = xx
		default:
			slog.Error("Invalid operand:", "operand", operand)
			panic("Invalid operand")
		}
	}
	return x
}
