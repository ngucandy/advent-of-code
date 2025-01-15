package aoc2024

import (
	"fmt"
	"github.com/ngucandy/advent-of-code/internal/helpers"
	"strings"
	"time"
)

func init() {
	DayMap["6"] = Day6{
		`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`,
	}
}

type Day6 struct {
	example string
}

func (d Day6) Part1(input string) {
	var grid [][]rune
	var sr, sc int
	for r, line := range strings.Split(input, "\n") {
		if c := strings.Index(line, "^"); c != -1 {
			sr, sc = r, c
		}
		grid = append(grid, []rune(line))
	}

	up := [2]int{-1, 0}
	down := [2]int{1, 0}
	left := [2]int{0, -1}
	right := [2]int{0, 1}

	dir := up
	visited := make(map[[2]int]struct{})
	for r, c := sr, sc; d.inbounds(r, c, grid); {
		visited[[2]int{r, c}] = struct{}{}
		nr, nc := r+dir[0], c+dir[1]
		if d.inbounds(nr, nc, grid) && grid[nr][nc] == '#' {
			// blocked so turn right 90 degrees
			switch dir {
			case up:
				dir = right
			case down:
				dir = left
			case right:
				dir = down
			case left:
				dir = up
			default:
				panic(fmt.Sprintf("unknown direction %v", dir))
			}
			continue
		}
		// not blocked
		r, c = nr, nc
	}
	fmt.Println("part1", len(visited))
}

func (d Day6) Part2(input string) {
	defer helpers.TrackTime(time.Now())
	var grid [][]rune
	var sr, sc int
	for r, line := range strings.Split(input, "\n") {
		if c := strings.Index(line, "^"); c != -1 {
			sr, sc = r, c
		}
		grid = append(grid, []rune(line))
	}

	up := [2]int{-1, 0}
	down := [2]int{1, 0}
	left := [2]int{0, -1}
	right := [2]int{0, 1}

	stuck := 0
	for or := range grid {
		for oc := range grid[or] {
			// cannot place new obstacle at starting location or at
			// existing obstacle
			if (or == sr && oc == sc) || grid[or][oc] == '#' {
				continue
			}

			grid[or][oc] = '#'
			dir := up
			visited := make(map[[4]int]bool)
			for r, c := sr, sc; d.inbounds(r, c, grid); {
				if visited[[4]int{r, c, dir[0], dir[1]}] {
					// stuck in a loop
					stuck++
					break
				}
				visited[[4]int{r, c, dir[0], dir[1]}] = true
				nr, nc := r+dir[0], c+dir[1]
				if d.inbounds(nr, nc, grid) && grid[nr][nc] == '#' {
					// blocked so turn right 90 degrees
					switch dir {
					case up:
						dir = right
					case down:
						dir = left
					case right:
						dir = down
					case left:
						dir = up
					default:
						panic(fmt.Sprintf("unknown direction %v", dir))
					}
					continue
				}
				// not blocked
				r, c = nr, nc
			}
			grid[or][oc] = '.'
		}
	}
	fmt.Println("part2", stuck)
}

func (d Day6) inbounds(r, c int, grid [][]rune) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[r])
}
