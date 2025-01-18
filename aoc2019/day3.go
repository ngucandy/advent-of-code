package aoc2019

import (
	"cmp"
	"fmt"
	"github.com/ngucandy/advent-of-code/internal/helpers"
	"slices"
	"strconv"
	"strings"
)

func init() {
	Days["3"] = Day3{
		`R8,U5,L5,D3
U7,R6,D4,L4`,
		`R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`,
		`R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`,
	}
}

type Day3 struct {
	example1, example2, example3 string
}

func (d Day3) Part1(input string) {
	paths := strings.Split(input, "\n")
	var grids []map[[2]int]struct{}
	for _, path := range paths {
		r, c := 0, 0
		grid := make(map[[2]int]struct{})
		for _, move := range strings.Split(path, ",") {
			dir := move[0]
			dist, _ := strconv.Atoi(move[1:])

			switch dir {
			case 'R':
				for nc := c + 1; nc <= c+dist; nc++ {
					grid[[2]int{r, nc}] = struct{}{}
				}
				c += dist
			case 'L':
				for nc := c - 1; nc >= c-dist; nc-- {
					grid[[2]int{r, nc}] = struct{}{}
				}
				c -= dist
			case 'U':
				for nr := r - 1; nr >= r-dist; nr-- {
					grid[[2]int{nr, c}] = struct{}{}
				}
				r -= dist
			case 'D':
				for nr := r + 1; nr <= r+dist; nr++ {
					grid[[2]int{nr, c}] = struct{}{}
				}
				r += dist
			default:
				panic(fmt.Sprintf("unknown direction: %q", dir))
			}
		}
		grids = append(grids, grid)
	}
	var crosses [][3]int
	for loc1 := range grids[0] {
		if _, exists := grids[1][loc1]; exists {
			dist := helpers.AbsInt(loc1[0]) + helpers.AbsInt(loc1[1])
			crosses = append(crosses, [3]int{loc1[0], loc1[1], dist})
		}
	}
	slices.SortFunc(crosses, func(a, b [3]int) int {
		return cmp.Compare(a[2], b[2])
	})
	fmt.Println("part1", crosses[0][2])
}

func (d Day3) Part2(input string) {
	paths := strings.Split(input, "\n")
	var grids []map[[2]int]int
	for _, path := range paths {
		r, c, steps := 0, 0, 0
		grid := make(map[[2]int]int)
		for _, move := range strings.Split(path, ",") {
			dir := move[0]
			dist, _ := strconv.Atoi(move[1:])

			switch dir {
			case 'R':
				for nc := c + 1; nc <= c+dist; nc++ {
					steps++
					if _, exists := grid[[2]int{r, nc}]; !exists {
						grid[[2]int{r, nc}] = steps
					}
				}
				c += dist
			case 'L':
				for nc := c - 1; nc >= c-dist; nc-- {
					steps++
					if _, exists := grid[[2]int{r, nc}]; !exists {
						grid[[2]int{r, nc}] = steps
					}
				}
				c -= dist
			case 'U':
				for nr := r - 1; nr >= r-dist; nr-- {
					steps++
					if _, exists := grid[[2]int{nr, c}]; !exists {
						grid[[2]int{nr, c}] = steps
					}
				}
				r -= dist
			case 'D':
				for nr := r + 1; nr <= r+dist; nr++ {
					steps++
					if _, exists := grid[[2]int{nr, c}]; !exists {
						grid[[2]int{nr, c}] = steps
					}
				}
				r += dist
			default:
				panic(fmt.Sprintf("unknown direction: %q", dir))
			}
		}
		grids = append(grids, grid)
	}
	var crosses [][3]int
	for loc1, steps1 := range grids[0] {
		if steps2, exists := grids[1][loc1]; exists {
			dist := steps1 + steps2
			crosses = append(crosses, [3]int{loc1[0], loc1[1], dist})
		}
	}
	slices.SortFunc(crosses, func(a, b [3]int) int {
		return cmp.Compare(a[2], b[2])
	})
	fmt.Println("part2", crosses[0][2])
}
