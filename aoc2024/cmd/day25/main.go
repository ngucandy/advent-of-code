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
	blocks := strings.Split(input, "\n\n")
	var locks, keys [][5]int
	lockOrKey := [5]int{-1, -1, -1, -1, -1}
	for _, block := range blocks {
		lines := strings.Split(block, "\n")
		for _, line := range lines {
			for i, char := range line {
				if char == '#' {
					lockOrKey[i]++
				}
			}
		}
		if lines[0] == "#####" {
			locks = append(locks, lockOrKey)
		} else {
			keys = append(keys, lockOrKey)
		}
		lockOrKey = [5]int{-1, -1, -1, -1, -1}
	}

	count := 0
	for _, lock := range locks {
	KeyLoop:
		for _, key := range keys {
			for i := range key {
				if lock[i]+key[i] > 5 {
					continue KeyLoop
				}
			}
			count++
		}
	}
	slog.Info("Part 1:", "count", count)
}

func part2(input string) {

}
