package main

import (
	"github.com/ngucandy/advent-of-code/internal/helpers"
	"log/slog"
	"os"
	"strings"
	"time"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	bytes, _ := os.ReadFile(infile)
	input := string(bytes)

	part1(input)
	part2(input)
}

func part1(input string) {
	defer helpers.TrackTime(time.Now(), "part1")
	grid := make([][]rune, 0)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	var start [2]int
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 'S' {
				start = [2]int{row, col}
				break
			}
		}
	}

	distances := make(map[[2]int]int)
	distances[start] = 0
	current := start
	for grid[current[0]][current[1]] != 'E' {
		for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			next := [2]int{current[0] + dir[0], current[1] + dir[1]}
			if next[0] < 0 || next[1] < 0 || next[0] >= len(grid) || next[1] >= len(grid) {
				continue
			}
			if grid[next[0]][next[1]] == '#' {
				continue
			}
			if _, ok := distances[next]; ok {
				continue
			}
			distances[next] = distances[current] + 1
			current = next
		}
	}

	count := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '#' {
				continue
			}
			for _, dir := range [][2]int{{2, 0}, {1, 1}, {0, 2}, {-1, 1}} {
				next := [2]int{row + dir[0], col + dir[1]}
				if next[0] < 0 || next[1] < 0 || next[0] >= len(grid) || next[1] >= len(grid) {
					continue
				}
				if grid[next[0]][next[1]] == '#' {
					continue
				}
				if _, ok := distances[[2]int{row, col}]; !ok {
					continue
				}
				if _, ok := distances[next]; !ok {
					continue
				}
				if helpers.AbsInt(distances[next]-distances[[2]int{row, col}]) >= 102 {
					count++
				}
			}
		}
	}
	slog.Info("Part 1:", "count", count)
}

func part2(input string) {
	defer helpers.TrackTime(time.Now(), "part2")
	grid := make([][]rune, 0)
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	var start [2]int
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 'S' {
				start = [2]int{row, col}
				break
			}
		}
	}

	distances := make(map[[2]int]int)
	distances[start] = 0
	current := start
	for grid[current[0]][current[1]] != 'E' {
		for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			next := [2]int{current[0] + dir[0], current[1] + dir[1]}
			if next[0] < 0 || next[1] < 0 || next[0] >= len(grid) || next[1] >= len(grid) {
				continue
			}
			if grid[next[0]][next[1]] == '#' {
				continue
			}
			if _, ok := distances[next]; ok {
				continue
			}
			distances[next] = distances[current] + 1
			current = next
		}
	}

	count := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '#' {
				continue
			}
			endLocations := make(map[[2]int]int)
			for radius := 2; radius <= 20; radius++ {
				for dr := 0; dr <= radius; dr++ {
					dc := radius - dr
					endLocations[[2]int{row + dr, col + dc}] = radius
					endLocations[[2]int{row + dr, col - dc}] = radius
					endLocations[[2]int{row - dr, col + dc}] = radius
					endLocations[[2]int{row - dr, col - dc}] = radius
				}
			}

			for next, radius := range endLocations {
				if next[0] < 0 || next[1] < 0 || next[0] >= len(grid) || next[1] >= len(grid) {
					continue
				}
				if grid[next[0]][next[1]] == '#' {
					continue
				}
				if _, ok := distances[[2]int{row, col}]; !ok {
					continue
				}
				if _, ok := distances[next]; !ok {
					continue
				}
				if distances[[2]int{row, col}]-distances[next] >= 100+radius {
					count++
				}
			}

		}
	}
	slog.Info("Part 2:", "count", count)

}
