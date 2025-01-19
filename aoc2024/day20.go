package aoc2024

import (
	"fmt"
	"strings"

	"github.com/ngucandy/advent-of-code/internal/helpers"
)

func init() {
	Days["20"] = Day20{
		eg1: `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`,
	}
}

type Day20 struct {
	eg1, eg2 string
}

func (d Day20) Part1(input string) {
	grid, sr, sc := d.makeGrid(input)
	path, steps := d.countSteps(sr, sc, grid)

	radius := 2
	save := 100
	count := d.countCheats(path, radius, grid, steps, save)
	fmt.Println("part1", count)
}

func (d Day20) countCheats(path [][2]int, radius int, grid [][]rune, steps map[[2]int]int, save int) int {
	count := 0
	for _, pos := range path {
		r, c := pos[0], pos[1]
		for nr := max(0, r-radius); nr <= min(len(grid)-1, r+radius); nr++ {
			for nc := max(0, c-radius); nc <= min(len(grid[0])-1, c+radius); nc++ {
				distance := helpers.AbsInt(nr-r) + helpers.AbsInt(nc-c)
				if distance > radius || grid[nr][nc] == '#' || steps[[2]int{nr, nc}]-steps[[2]int{r, c}] < save+distance {
					continue
				}
				count++
			}
		}
	}
	return count
}

func (d Day20) countSteps(sr int, sc int, grid [][]rune) ([][2]int, map[[2]int]int) {
	var path [][2]int
	steps := make(map[[2]int]int)
	q := [][3]int{{sr, sc, 0}}
	for len(q) > 0 {
		r, c, s := q[0][0], q[0][1], q[0][2]
		q = q[1:]
		path = append(path, [2]int{r, c})
		steps[[2]int{r, c}] = s

		for _, dir := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			nr, nc := r+dir[0], c+dir[1]
			if _, exists := steps[[2]int{nr, nc}]; exists || grid[nr][nc] == '#' {
				continue
			}
			q = append(q, [3]int{nr, nc, s + 1})
			break
		}
	}
	return path, steps
}

func (d Day20) makeGrid(input string) ([][]rune, int, int) {
	var grid [][]rune
	var sr, sc int
	for r, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
		if c := strings.Index(line, "S"); c != -1 {
			sr, sc = r, c
		}
	}
	return grid, sr, sc
}

func (d Day20) Part2(input string) {
	grid, sr, sc := d.makeGrid(input)
	path, steps := d.countSteps(sr, sc, grid)

	radius := 20
	save := 100
	count := d.countCheats(path, radius, grid, steps, save)
	fmt.Println("part2", count)
}
