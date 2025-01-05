package main

import (
	"log/slog"
	"os"
	"slices"
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
	grid, trenchSize, tr, tc := buildGrid(input)

TrenchNeighbors:
	for _, dir := range [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}} {
		nr := tr + dir[0]
		nc := tc + dir[1]
		if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) {
			continue
		}
		if grid[nr][nc] == '#' {
			continue
		}
		q := [][2]int{{nr, nc}}
		seen := make(map[[2]int]bool)
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			r := cur[0]
			c := cur[1]
			// if we're outside the grid, then we're outside the trench
			if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
				continue TrenchNeighbors
			}
			if grid[r][c] == '#' {
				continue
			}
			if seen[cur] {
				continue
			}
			seen[cur] = true
			for _, d := range [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
				q = append(q, [2]int{r + d[0], c + d[1]})
			}
		}
		// we've visited every node in the q without hitting the edge of the grid
		slog.Info("Part 1:", "capacity", trenchSize+len(seen))
	}
}

func buildGrid(input string) (grid [][]rune, trenches, r, c int) {
	grid = [][]rune{{'#'}}
	r, c = 0, 0
	trenches = 0
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, " ")
		dir := parts[0]
		n, _ := strconv.Atoi(parts[1])
		trenches += n
		switch dir {
		case "R":
			if c+n >= len(grid[0]) {
				for row := range grid {
					expand := c + n - (len(grid[row]) - 1)
					grid[row] = append(grid[row], slices.Repeat([]rune{'.'}, expand)...)
				}
			}
			for range n {
				grid[r][c+1] = '#'
				c++
			}
		case "L":
			if c-n < 0 {
				expand := 0 - c + n
				for row := range grid {
					grid[row] = append(slices.Repeat([]rune{'.'}, expand), grid[row]...)
				}
				c += expand
			}
			for range n {
				grid[r][c-1] = '#'
				c--
			}
		case "U":
			if r-n < 0 {
				expand := 0 - r + n
				for range expand {
					grid = append([][]rune{slices.Repeat([]rune{'.'}, len(grid[0]))}, grid...)
				}
				r += expand
			}
			for range n {
				grid[r-1][c] = '#'
				r--
			}
		case "D":
			if r+n >= len(grid) {
				expand := r + n - (len(grid) - 1)
				for range expand {
					grid = append(grid, slices.Repeat([]rune{'.'}, len(grid[0])))
				}
			}
			for range n {
				grid[r+1][c] = '#'
				r++
			}
		}
	}
	return grid, trenches, r, c
}

func part2(input string) {

}
