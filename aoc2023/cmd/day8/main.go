package main

import (
	"bufio"
	"log/slog"
	"os"
	"regexp"
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
}

func part1(instructions string, elements map[string][2]string) {
	steps := 0
	current := "AAA"
	for current != "ZZZ" {
		for _, instruction := range instructions {
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
	slog.Info("Part 1:", "steps", steps)
}
