package main

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, _ := os.Open(infile)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	input := strings.Split(line, " ")

	part1(input)
}

func part1(input []string) {
	stones := input

	for range 25 {
		next := make([]string, 0, len(stones))
		for _, stone := range stones {
			if stone == "0" {
				next = append(next, "1")
				continue
			}
			if len(stone)%2 == 0 {
				next = append(next, stone[:len(stone)/2])
				n, _ := strconv.Atoi(stone[len(stone)/2:])
				next = append(next, strconv.Itoa(n))
				continue
			}
			n, _ := strconv.Atoi(stone)
			next = append(next, strconv.Itoa(n*2024))
		}
		stones = next
	}

	slog.Info("Part 1:", "stones", len(stones))
}
