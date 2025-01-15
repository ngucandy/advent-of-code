package aoc2024

import (
	"fmt"
	"strings"
)

func init() {
	DayMap["4"] = Day4{}
}

type Day4 struct {
}

func (d Day4) Part1(input string) {
	var grid [][]rune
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}
	directions := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	count := 0
	word := []rune("XMAS")
	for r := range grid {
		for c := range grid[r] {
			for _, dir := range directions {
				if d.match(word, r, c, 0, dir, grid) {
					count++
				}
			}
		}
	}
	fmt.Println("part1", count)
}

func (d Day4) Part2(input string) {
	var grid [][]rune
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}
	// only diagonal directions
	directions := [][2]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}
	// map of found diagonal `MAS` words keyed by position of `A`
	found := make(map[[2]int]int)
	count := 0
	word := []rune("MAS")
	for r := range grid {
		for c := range grid[r] {
			for _, dir := range directions {
				if d.match(word, r, c, 0, dir, grid) {
					found[[2]int{r + dir[0], c + dir[1]}]++
					// 2 `MAS` diagonals that share an `A` form an X
					if found[[2]int{r + dir[0], c + dir[1]}] == 2 {
						count++
					}
				}
			}
		}
	}
	fmt.Println("part2", count)
}

// match recursively searches for `word` in the grid. The level of recursion
// determines which letter in `word` we're looking for.  Matching follows the
// direction given by `dir`.
func (d Day4) match(word []rune, r, c, level int, dir [2]int, grid [][]rune) bool {
	if level >= len(word) {
		return true
	}
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) || grid[r][c] != word[level] {
		return false
	}
	return d.match(word, r+dir[0], c+dir[1], level+1, dir, grid)
}
