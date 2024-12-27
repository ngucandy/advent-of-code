package main

import (
	"log/slog"
	"os"
	"strconv"
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
	total := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		var seq []int
		for _, num := range parts {
			n, _ := strconv.Atoi(num)
			seq = append(seq, n)
		}
		total += nextVal(seq)
	}
	slog.Info("Part 1:", "total", total)
}

func nextVal(seq []int) int {
	if allZero(seq) {
		return 0
	}

	nextSeq := make([]int, len(seq)-1)
	for i := 0; i < len(seq)-1; i++ {
		nextSeq[i] = seq[i+1] - seq[i]
	}
	next := nextVal(nextSeq)
	return seq[len(seq)-1] + next
}

func allZero(s []int) bool {
	for _, num := range s {
		if num != 0 {
			return false
		}
	}
	return true
}

func part2(input string) {

}
