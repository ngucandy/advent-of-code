package aoc2024

import (
	"fmt"
	"slices"
	"strings"
)

func init() {
	Days["18"] = Day18{}
}

type Day18 struct {
	eg1, eg2 string
}

func (d Day18) Part1(input string) {
	size := 71
	var grid [][]rune
	for range size {
		grid = append(grid, slices.Repeat([]rune{'.'}, size))
	}

	n := 1024
	for i, line := range strings.Split(input, "\n") {
		if i == n {
			break
		}
		var r, c int
		_, _ = fmt.Sscanf(line, "%d,%d", &c, &r)
		grid[r][c] = '#'
	}

	sr, sc := 0, 0
	q := [][3]int{{sr, sc, 0}}
	seen := make(map[[2]int]bool)
	for len(q) > 0 {
		r, c, s := q[0][0], q[0][1], q[0][2]
		q = q[1:]

		if r == len(grid)-1 && c == len(grid[0])-1 {
			fmt.Println("part1", s)
			break
		}

		if seen[[2]int{r, c}] {
			continue
		}
		seen[[2]int{r, c}] = true

		for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			nr, nc := r+dir[0], c+dir[1]
			if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) || grid[nr][nc] == '#' {
				continue
			}
			q = append(q, [3]int{nr, nc, s + 1})
		}
	}
}

func (d Day18) Part2(input string) {
	size := 71
	var grid [][]rune
	for range size {
		grid = append(grid, slices.Repeat([]rune{'.'}, size))
	}

	n := 1024
line:
	for i, line := range strings.Split(input, "\n") {
		var y, x int
		_, _ = fmt.Sscanf(line, "%d,%d", &x, &y)
		grid[y][x] = '#'
		if i < n {
			continue
		}

		sr, sc := 0, 0
		q := [][3]int{{sr, sc, 0}}
		seen := make(map[[2]int]bool)
		for len(q) > 0 {
			r, c, s := q[0][0], q[0][1], q[0][2]
			q = q[1:]

			if r == len(grid)-1 && c == len(grid[0])-1 {
				continue line
			}

			if seen[[2]int{r, c}] {
				continue
			}
			seen[[2]int{r, c}] = true

			for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				nr, nc := r+dir[0], c+dir[1]
				if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) || grid[nr][nc] == '#' {
					continue
				}
				q = append(q, [3]int{nr, nc, s + 1})
			}
		}
		fmt.Printf("part2 %d,%d\n", x, y)
		break
	}
}
