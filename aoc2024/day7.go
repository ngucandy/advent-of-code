package aoc2024

import (
	"fmt"
	"github.com/ngucandy/advent-of-code/internal/helpers"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	Days["7"] = Day7{
		`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`,
	}
}

type Day7 struct {
	example string
}

func (d Day7) Part1(input string) {
	ops := "+*"
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")
		value, _ := strconv.Atoi(parts[0])
		var operands []int
		for _, operand := range strings.Split(parts[1], " ") {
			n, _ := strconv.Atoi(operand)
			operands = append(operands, n)
		}
		combos := helpers.CartesianProductN([]rune(ops), len(operands)-1)
		for _, operators := range combos {
			if value == d.eval(operands, operators) {
				sum += value
				break
			}
		}
	}
	fmt.Println("part1", sum)
}

func (d Day7) Part2(input string) {
	defer helpers.TrackTime(time.Now())
	ops := "+*|"
	sum := int64(0)
	wg := &sync.WaitGroup{}
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")
		value, _ := strconv.Atoi(parts[0])
		var operands []int
		for _, operand := range strings.Split(parts[1], " ") {
			n, _ := strconv.Atoi(operand)
			operands = append(operands, n)
		}
		combos := helpers.CartesianProductN([]rune(ops), len(operands)-1)
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, operators := range combos {
				if value == d.eval(operands, operators) {
					atomic.AddInt64(&sum, int64(value))
					return
				}
			}
		}()
	}
	wg.Wait()
	fmt.Println("part2", sum)
}

func (d Day7) eval(operands []int, operators []rune) int {
	n1, n2 := operands[0], operands[1]
	op := operators[0]

	var val int
	switch op {
	case '+':
		val = n1 + n2
	case '*':
		val = n1 * n2
	case '|':
		val, _ = strconv.Atoi(strings.Join([]string{strconv.Itoa(n1), strconv.Itoa(n2)}, ""))
	default:
		panic(fmt.Sprintf("unknown operators %c", op))
	}

	if len(operands) == 2 {
		return val
	}
	return d.eval(append([]int{val}, operands[2:]...), operators[1:])
}
