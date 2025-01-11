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

	maxSteps = 64
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	bytes, _ := os.ReadFile(infile)
	input := strings.ReplaceAll(string(bytes), "\r\n", "\n")

	part1(testInput)
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

	seen := make(map[[2]int]bool)
	reachable := make(map[[2]int]struct{})
	q := [][3]int{{start[0], start[1], 0}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		r := cur[0]
		c := cur[1]
		steps := cur[2]

		if steps > maxSteps {
			continue
		}

		if seen[[2]int{r, c}] {
			continue
		}

		seen[[2]int{r, c}] = true
		if maxSteps%2 == 0 && steps%2 == 0 {
			reachable[[2]int{r, c}] = struct{}{}
		} else if maxSteps%2 == 1 && steps%2 == 1 {
			reachable[[2]int{r, c}] = struct{}{}
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
	slog.Info("Part 1:", "reachable", len(reachable))
}

func part2(input string) {

}
