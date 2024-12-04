package main

import (
	"bufio"
	"log/slog"
	"os"
	"regexp"
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

	part1(puzzle)
	part2(puzzle)
}

func part2(puzzle [][]string) {
	// M.S
	// .A.
	// M.S
	// M.S.A.M.S

	// S.S
	// .A.
	// M.M
	// S.S.A.M.M

	// S.M
	// .A.
	// S.M
	// S.M.A.S.M

	// M.M
	// .A.
	// S.S
	// M.M.A.S.S

	rexpMSAMS := regexp.MustCompile(`M.S.A.M.S`)
	rexpSSAMM := regexp.MustCompile(`S.S.A.M.M`)
	rexpSMASM := regexp.MustCompile(`S.M.A.S.M`)
	rexpMMASS := regexp.MustCompile(`M.M.A.S.S`)

	count := 0
	for y := range len(puzzle) - 2 {
		for x := range len(puzzle[y]) - 2 {
			block := getBlock(x, y, 3, puzzle)
			if rexpMSAMS.MatchString(block) || rexpSSAMM.MatchString(block) || rexpSMASM.MatchString(block) || rexpMMASS.MatchString(block) {
				count++
			}
		}
	}
	slog.Info("Part 2:", "count", count)
}

func getBlock(x int, y int, n int, puzzle [][]string) string {
	block := ""
	for j := range n {
		for i := range n {
			block += puzzle[y+j][x+i]
		}
	}
	//slog.Info("getBlock:", "block", block)
	return block
}

func part1(puzzle [][]string) {
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
