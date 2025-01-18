package aoc2024

import (
	"fmt"
	"strings"
	"time"

	"github.com/ngucandy/advent-of-code/internal/helpers"
)

func init() {
	Days["10"] = Day10{
		`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`,
	}
}

type Day10 struct {
	example string
}

func (d Day10) Part1(input string) {
	d.part1Iterative(input)
	d.part1Recursive(input)
}

func (d Day10) part1Iterative(input string) {
	defer helpers.TrackTime(time.Now())
	var grid [][]int
	var heads [][2]int
	for r, line := range strings.Split(input, "\n") {
		grid = append(grid, []int{})
		for c, ch := range line {
			n := int(ch - '0')
			grid[r] = append(grid[r], n)
			if n == 0 {
				heads = append(heads, [2]int{r, c})
			}
		}
	}

	score := 0
	for _, head := range heads {
		ends := make(map[[2]int]struct{})
		q := [][2]int{head}
		for len(q) > 0 {
			r, c := q[0][0], q[0][1]
			q = q[1:]

			if grid[r][c] == 9 {
				ends[[2]int{r, c}] = struct{}{}
				continue
			}

			for _, dir := range [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
				nr, nc := r+dir[0], c+dir[1]
				if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) {
					continue
				}
				if grid[nr][nc] == grid[r][c]+1 {
					q = append(q, [2]int{nr, nc})
				}
			}
		}
		score += len(ends)
	}

	fmt.Println("part1 iterative", score)
}

func (d Day10) part1Recursive(input string) {
	defer helpers.TrackTime(time.Now())
	var grid [][]int
	var heads, ends [][2]int

	for r, line := range strings.Split(input, "\n") {
		grid = append(grid, []int{})
		for c, ch := range line {
			n := int(ch - '0')
			grid[r] = append(grid[r], n)
			switch n {
			case 0:
				heads = append(heads, [2]int{r, c})
			case 9:
				ends = append(ends, [2]int{r, c})
			}
		}
	}

	score := 0
	for _, head := range heads {
		for _, end := range ends {
			if d.reachable(head, end, grid) {
				score++
			}
		}
	}
	fmt.Println("part1 recursive", score)
}

func (d Day10) reachable(start [2]int, end [2]int, grid [][]int) bool {
	if start == end {
		return true
	}
	for _, dir := range [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
		next := [2]int{start[0] + dir[0], start[1] + dir[1]}
		if next[0] < 0 || next[0] >= len(grid) || next[1] < 0 || next[1] >= len(grid[0]) {
			continue
		}
		if grid[next[0]][next[1]]-grid[start[0]][start[1]] != 1 {
			continue
		}
		if d.reachable(next, end, grid) {
			return true
		}
	}
	return false
}

func (d Day10) Part2(input string) {
	d.part2Iterative(input)
	d.part2Recursive(input)
}

func (d Day10) part2Iterative(input string) {
	defer helpers.TrackTime(time.Now())
	var grid [][]int
	var heads [][2]int
	for r, line := range strings.Split(input, "\n") {
		grid = append(grid, []int{})
		for c, ch := range line {
			n := int(ch - '0')
			grid[r] = append(grid[r], n)
			if n == 0 {
				heads = append(heads, [2]int{r, c})
			}
		}
	}

	score := 0
	for _, head := range heads {
		q := [][2]int{head}
		for len(q) > 0 {
			r, c := q[0][0], q[0][1]
			q = q[1:]

			if grid[r][c] == 9 {
				score++
				continue
			}

			for _, dir := range [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
				nr, nc := r+dir[0], c+dir[1]
				if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) {
					continue
				}
				if grid[nr][nc] == grid[r][c]+1 {
					q = append(q, [2]int{nr, nc})
				}
			}
		}
	}

	fmt.Println("part2 iterative", score)
}

func (d Day10) part2Recursive(input string) {
	defer helpers.TrackTime(time.Now())
	var grid [][]int
	var heads, ends [][2]int

	for r, line := range strings.Split(input, "\n") {
		grid = append(grid, []int{})
		for c, ch := range line {
			n := int(ch - '0')
			grid[r] = append(grid[r], n)
			switch n {
			case 0:
				heads = append(heads, [2]int{r, c})
			case 9:
				ends = append(ends, [2]int{r, c})
			}
		}
	}

	score := 0
	for _, head := range heads {
		for _, end := range ends {
			score += d.count(head, end, grid, make(map[[2][2]int]int))
		}
	}
	fmt.Println("part2 recursive", score)
}

func (d Day10) count(start [2]int, end [2]int, grid [][]int, cache map[[2][2]int]int) int {
	if start == end {
		return 1
	}
	if n, ok := cache[[2][2]int{start, end}]; ok {
		return n
	}
	n := 0
	for _, dir := range [][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
		next := [2]int{start[0] + dir[0], start[1] + dir[1]}
		if next[0] < 0 || next[0] >= len(grid) || next[1] < 0 || next[1] >= len(grid[0]) {
			continue
		}
		if grid[next[0]][next[1]]-grid[start[0]][start[1]] != 1 {
			continue
		}
		n += d.count(next, end, grid, cache)
	}
	cache[[2][2]int{start, end}] = n
	return n
}
