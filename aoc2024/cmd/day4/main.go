package main

import (
	"bufio"
	"log/slog"
	"os"
	"strings"
)

var (
	up        = [...]int{0, -1}
	down      = [...]int{0, 1}
	left      = [...]int{-1, 0}
	right     = [...]int{1, 0}
	upleft    = [...]int{-1, -1}
	upright   = [...]int{1, -1}
	downleft  = [...]int{-1, 1}
	downright = [...]int{1, 1}
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			slog.Error("Error closing file:", "error", err)
		}
	}(file)

	var puzzle [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		puzzle = append(puzzle, strings.Split(line, ""))
	}

	word := "XMAS"
	count := 0
	for y := range len(puzzle) {
		for x := range len(puzzle[y]) {
			count += search(word, x, y, right, puzzle)
			count += search(word, x, y, left, puzzle)
			count += search(word, x, y, up, puzzle)
			count += search(word, x, y, down, puzzle)
			count += search(word, x, y, upright, puzzle)
			count += search(word, x, y, downright, puzzle)
			count += search(word, x, y, upleft, puzzle)
			count += search(word, x, y, downleft, puzzle)
		}
	}
	slog.Info("Part 1:", "count", count)
}

func search(word string, x int, y int, direction [2]int, puzzle [][]string) int {
	assembled := ""
	for range len(word) {
		assembled += puzzle[y][x]
		x, y = move(x, y, direction)
		if x < 0 || y < 0 || y >= len(puzzle) || x >= len(puzzle[y]) {
			break
		}
	}
	//slog.Info("Search:", "word", word, "assembled", assembled)
	if word == assembled {
		return 1
	}
	return 0
}

func move(x int, y int, direction [2]int) (int, int) {
	return x + direction[0], y + direction[1]
}
