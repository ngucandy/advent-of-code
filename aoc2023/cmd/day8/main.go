package main

import (
	"bufio"
	"github.com/ngucandy/advent-of-code/internal/helpers"
	"log/slog"
	"os"
	"regexp"
	"strings"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, _ := os.Open(infile)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	rexpElements := regexp.MustCompile(`(...) = \((...), (...)\)`)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instructions := scanner.Text()
	elements := make(map[string][2]string)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		matches := rexpElements.FindStringSubmatch(line)
		elements[matches[1]] = [2]string{matches[2], matches[3]}
	}

	part1(instructions, elements)
	part2(instructions, elements)
}

func part2(instructions string, elements map[string][2]string) {
	steps := make([]int, 0)
	for k := range elements {
		if strings.HasSuffix(k, "A") {
			steps = append(steps, countSteps(instructions, elements, k))
		}
	}
	lcm := helpers.LCM(steps[0], steps[1], steps[2:]...)

	slog.Info("Part 2:", "steps", steps, "lcm", lcm)
}

func part1(instructions string, elements map[string][2]string) {
	steps := countSteps(instructions, elements, "AAA")
	slog.Info("Part 1:", "steps", steps)
}

func countSteps(instructions string, elements map[string][2]string, start string) int {
	steps := 0
	current := start
	for !strings.HasSuffix(current, "Z") {
		for _, instruction := range instructions {
			if strings.HasSuffix(current, "Z") {
				break
			}
			switch instruction {
			case 'L':
				current = elements[current][0]
			case 'R':
				current = elements[current][1]
			default:
				panic("Invalid instruction")
			}
			steps++
		}
	}
	return steps
}
