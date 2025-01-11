package main

import (
	"log/slog"
	"os"
	"strings"
)

const (
	testInput = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	bytes, _ := os.ReadFile(infile)
	input := strings.ReplaceAll(string(bytes), "\r\n", "\n")

	//part1(testInput)
	part1(input)
	part2(testInput)
	part2(input)
}

func part1(input string) {
	var grid [][]rune
	var start [2]int
	for row, line := range strings.Split(input, "\n") {
		if col := strings.Index(line, "S"); col >= 0 {
			start[0] = row
			start[1] = col
			line = strings.Replace(line, "S", ".", 1)
		}
		grid = append(grid, []rune(line))
	}

	reached := make(map[[2]int]struct{})
	maxSteps := 64
	q := [][3]int{{start[0], start[1], 0}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		r := cur[0]
		c := cur[1]
		steps := cur[2]

		if steps == maxSteps && grid[r][c] == '.' {
			reached[[2]int{r, c}] = struct{}{}
		}

		if steps > maxSteps {
			continue
		}

		for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			nr := r + dir[0]
			nc := c + dir[1]
			if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) || grid[nr][nc] == '#' {
				continue
			}
			q = append(q, [3]int{nr, nc, steps + 1})
		}
	}
	slog.Info("Part 1:", "reached", len(reached))
}

func part2(input string) {

}
