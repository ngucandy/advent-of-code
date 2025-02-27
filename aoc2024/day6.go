package aoc2024

import (
	"fmt"
	"github.com/ngucandy/advent-of-code/internal/helpers"
	"strings"
	"sync"
	"sync/atomic"
)

func init() {
	Days["6"] = Day6{
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

func (d Day6) Part1(input string) any {
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
	return len(visited)
}

func (d Day6) Part2(input string) any {
	var grid0 [][]rune
	var sr, sc int
	for r, line := range strings.Split(input, "\n") {
		if c := strings.Index(line, "^"); c != -1 {
			sr, sc = r, c
		}
		grid0 = append(grid0, []rune(line))
	}

	up := [2]int{-1, 0}
	down := [2]int{1, 0}
	left := [2]int{0, -1}
	right := [2]int{0, 1}

	wg := &sync.WaitGroup{}
	stuck := int32(0)
	// or,oc represent coordinates for new obstacle
	for or := range grid0 {
		for oc := range grid0[or] {
			// cannot place new obstacle at starting location or at
			// existing obstacle
			if (or == sr && oc == sc) || grid0[or][oc] == '#' {
				continue
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
				grid := helpers.CloneGrid(grid0)
				grid[or][oc] = '#'
				dir := up
				visited := make(map[[4]int]bool)
				for r, c := sr, sc; d.inbounds(r, c, grid); {
					if visited[[4]int{r, c, dir[0], dir[1]}] {
						// stuck in a loop
						atomic.AddInt32(&stuck, 1)
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
			}()
		}
	}
	wg.Wait()
	return stuck
}

func (d Day6) inbounds(r, c int, grid [][]rune) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[r])
}
