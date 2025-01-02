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

const (
	right = iota
	down
	left
	up
)

func part1(input string) {
	var grid [][]rune
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	seen := make(map[[3]int]struct{})
	queue := [][3]int{{0, 0, right}}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		row := current[0]
		col := current[1]
		dir := current[2]

		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
			continue
		}

		if _, exists := seen[current]; exists {
			continue
		}
		seen[current] = struct{}{}

		switch grid[row][col] {
		case '.':
			queue = append(queue, [3]int{row + directions[dir][0], col + directions[dir][1], dir})
		case '\\':
			var newdir int
			switch dir {
			case right:
				newdir = down
			case down:
				newdir = right
			case left:
				newdir = up
			case up:
				newdir = left
			}
			queue = append(queue, [3]int{row + directions[newdir][0], col + directions[newdir][1], newdir})
		case '/':
			var newdir int
			switch dir {
			case right:
				newdir = up
			case down:
				newdir = left
			case left:
				newdir = down
			case up:
				newdir = right
			}
			queue = append(queue, [3]int{row + directions[newdir][0], col + directions[newdir][1], newdir})
		case '|':
			switch dir {
			case left, right:
				queue = append(queue, [3]int{row + directions[up][0], col + directions[up][1], up})
				queue = append(queue, [3]int{row + directions[down][0], col + directions[down][1], down})
			case up, down:
				queue = append(queue, [3]int{row + directions[dir][0], col + directions[dir][1], dir})
			}
		case '-':
			switch dir {
			case left, right:
				queue = append(queue, [3]int{row + directions[dir][0], col + directions[dir][1], dir})
			case up, down:
				queue = append(queue, [3]int{row + directions[left][0], col + directions[left][1], left})
				queue = append(queue, [3]int{row + directions[right][0], col + directions[right][1], right})
			}
		}
	}
	energized := make(map[[2]int]struct{})
	for tile := range seen {
		energized[[2]int{tile[0], tile[1]}] = struct{}{}
	}
	slog.Info("Part 1:", "energized", len(energized))
}

func part2(input string) {

}
