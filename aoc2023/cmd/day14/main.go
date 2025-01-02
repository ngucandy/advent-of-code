package main

import (
	"log/slog"
	"os"
	"slices"
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
		moveRock(rock, [2]int{-1, 0}, grid)
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

func moveRock(rock, direction [2]int, grid [][]rune) bool {
	neighbor := [2]int{rock[0] + direction[0], rock[1] + direction[1]}
	if neighbor[0] < 0 || neighbor[1] < 0 || neighbor[0] >= len(grid) || neighbor[1] >= len(grid[0]) {
		return false
	}
	if grid[neighbor[0]][neighbor[1]] == '#' {
		return false
	}
	if grid[neighbor[0]][neighbor[1]] == '.' {
		grid[neighbor[0]][neighbor[1]] = 'O'
		grid[rock[0]][rock[1]] = '.'
		moveRock(neighbor, direction, grid)
		return true
	}
	if moveRock(neighbor, direction, grid) {
		grid[neighbor[0]][neighbor[1]] = 'O'
		grid[rock[0]][rock[1]] = '.'
		return true
	}
	// neighbor can't be moved
	return false
}

func part2(input string) {
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

	var finalGrid string
	seenMap := make(map[string]struct{})
	var seenList []string
	seenMap[ToString(grid)] = struct{}{}
	seenList = append(seenList, ToString(grid))
	cycles := 1000000000
	directions := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	for cycle := 1; cycle <= cycles; cycle++ {
		for _, direction := range directions {
			for row := range grid {
				for col := range grid[row] {
					if grid[row][col] == 'O' {
						moveRock([2]int{row, col}, direction, grid)
					}
				}
			}
		}

		sgrid := ToString(grid)
		if _, exists := seenMap[sgrid]; exists {
			first := slices.Index(seenList, sgrid)
			i := (cycles-first)%(cycle-first) + first
			finalGrid = seenList[i]
			break
		}
		seenMap[sgrid] = struct{}{}
		seenList = append(seenList, sgrid)
	}

	total := 0
	for row, line := range strings.Split(finalGrid, "\n") {
		for _, char := range line {
			if char == 'O' {
				total += len(grid) - row
			}
		}
	}
	slog.Info("Part 2:", "total", total)
}

func ToString(grid [][]rune) string {
	rows := make([]string, 0)
	for _, row := range grid {
		rows = append(rows, string(row))
	}
	return strings.Join(rows, "\n")
}
