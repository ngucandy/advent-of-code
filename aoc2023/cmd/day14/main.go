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
	var grid [][]rune
	var rocks [][2]int
	lines := strings.Split(input, "\n")
	for row, line := range lines {
		grid = append(grid, []rune(line))
		for col, ch := range line {
			if ch == 'O' {
				rocks = append(rocks, [2]int{row, col})
			}
		}
	}
	for _, rock := range rocks {
		moveRock(rock, grid)
	}

	total := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 'O' {
				total += len(grid) - row
			}
		}
	}
	slog.Info("Part 1:", "total", total)
}

func moveRock(rock [2]int, grid [][]rune) {
	above := [2]int{rock[0] - 1, rock[1]}
	if above[0] < 0 || grid[above[0]][above[1]] != '.' {
		return
	}
	grid[above[0]][above[1]] = 'O'
	grid[rock[0]][rock[1]] = '.'
	moveRock(above, grid)
}

func part2(input string) {

}
