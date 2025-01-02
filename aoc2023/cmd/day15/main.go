package main

import (
	"log/slog"
	"os"
	"strings"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	bytes, _ := os.ReadFile(infile)
	input := strings.ReplaceAll(string(bytes), "\r\n", "\n")

	part1(input)
	part2(input)
}

func part1(input string) {
	steps := strings.Split(input, ",")
	sum := 0
	for _, step := range steps {
		sum += hash(step)
	}
	slog.Info("Part 1:", "sum", sum)
}

func hash(s string) int {
	current := 0
	for _, ch := range s {
		current += int(ch)
		current *= 17
		current %= 256
	}
	return current
}

func part2(input string) {

}
