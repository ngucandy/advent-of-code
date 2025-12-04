package aoc2025

import (
	"strings"
)

func init() {
	Days["4"] = Day4{
		example: `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`,
	}
}

type Day4 struct {
	example string
}

func (d Day4) Part1(input string) any {
	rollcount := 0
	grid := d.parseGrid(input)
	for r, row := range grid {
		for c, ch := range row {
			if ch == '.' {
				continue
			}
			// found a roll, count adjacent rolls
			adjacent := 0
			for _, drc := range [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
				ar := r + drc[0]
				ac := c + drc[1]
				if ar < 0 || ar >= len(grid) || ac < 0 || ac >= len(grid[r]) {
					// out of grid bounds
					continue
				}
				if grid[ar][ac] == '@' {
					adjacent++
				}
			}
			if adjacent < 4 {
				rollcount++
			}
		}
	}
	return rollcount
}

func (d Day4) parseGrid(input string) [][]rune {
	var grid [][]rune
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}
	return grid
}

func (d Day4) Part2(input string) any {
	total := 0
	grid := d.parseGrid(input)
	// loop through removing rolls until we're unable to remove more
	for removed := d.removeRolls(grid); removed > 0; removed = d.removeRolls(grid) {
		total += removed
	}
	return total
}

// removeRolls modifies the given grid by removing rolls that are removable
// and returns the number of rolls removed
func (d Day4) removeRolls(grid [][]rune) int {
	var removable [][2]int
	for r, row := range grid {
		for c, ch := range row {
			if ch == '.' {
				continue
			}
			// found a roll, count adjacent rolls
			adjacent := 0
			for _, drc := range [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}} {
				ar := r + drc[0]
				ac := c + drc[1]
				if ar < 0 || ar >= len(grid) || ac < 0 || ac >= len(grid[r]) {
					// out of grid bounds
					continue
				}
				if grid[ar][ac] == '@' {
					adjacent++
				}
			}
			if adjacent < 4 {
				removable = append(removable, [2]int{r, c})
			}
		}
	}
	for _, point := range removable {
		grid[point[0]][point[1]] = '.'
	}
	return len(removable)
}
