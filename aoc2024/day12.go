package aoc2024

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func init() {
	d := &Day12{
		up:    [2]int{0, -1},
		down:  [2]int{0, 1},
		left:  [2]int{-1, 0},
		right: [2]int{1, 0},
	}
	d.directions = [][2]int{d.right, d.down, d.left, d.up}
	Days["12"] = d

}

type Day12 struct {
	example               string
	up, down, left, right [2]int
	directions            [][2]int
}

func (d Day12) Part1(input string) {
	var grid [][]rune
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}

	total := 0
	visited := make(map[[2]int]bool)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if visited[[2]int{x, y}] {
				continue
			}
			a, p := d.areaPerimeter(x, y, grid[y][x], grid, visited)
			total += a * p
		}
	}
	fmt.Println("part1", total)
}

func (d Day12) Part2(input string) {
	var grid [][]rune
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}

	total := 0
	regions := [][][2]int{}
	visited := make(map[[2]int]bool)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if visited[[2]int{x, y}] {
				continue
			}
			regions = append(regions, d.fill(x, y, grid[y][x], grid, visited))
		}
	}

	for _, region := range regions {
		m := make(map[[2]int]bool, len(region))
		for _, plot := range region {
			m[plot] = true
		}
		sides := d.countSides(m)
		area := len(region)
		total += area * sides
	}
	fmt.Println("part2", total)
}

func (d Day12) areaPerimeter(x int, y int, plant rune, grid [][]rune, visited map[[2]int]bool) (area int, perimeter int) {
	if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) || plant != grid[y][x] {
		return 0, 1
	}
	if _, ok := visited[[2]int{x, y}]; ok {
		return 0, 0
	}
	visited[[2]int{x, y}] = true
	area = 1
	perimeter = 0

	for _, direction := range d.directions {
		aNext, pNext := d.areaPerimeter(x+direction[0], y+direction[1], plant, grid, visited)
		area += aNext
		perimeter += pNext
	}
	return area, perimeter
}

func (d Day12) fill(x int, y int, plant rune, grid [][]rune, visited map[[2]int]bool) [][2]int {
	if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) || plant != grid[y][x] {
		return [][2]int{}
	}
	if _, ok := visited[[2]int{x, y}]; ok {
		return [][2]int{}
	}
	visited[[2]int{x, y}] = true
	region := [][2]int{{x, y}}
	for _, direction := range d.directions {
		region = append(region, d.fill(x+direction[0], y+direction[1], plant, grid, visited)...)
	}
	return region
}

func (d Day12) countSides(region map[[2]int]bool) int {
	sides := make(map[[2]int][][2]int)
	sides[d.up] = [][2]int{}
	sides[d.down] = [][2]int{}
	sides[d.right] = [][2]int{}
	sides[d.left] = [][2]int{}

	for plot := range region {
		for _, direction := range d.directions {
			next := [2]int{plot[0] + direction[0], plot[1] + direction[1]}
			if _, ok := region[next]; !ok {
				sides[direction] = append(sides[direction], plot)
			}
		}
	}

	cmpHoriz := func(a, b [2]int) int {
		if a[1] == b[1] {
			return cmp.Compare(a[0], b[0])
		}
		return cmp.Compare(a[1], b[1])
	}
	cmpVert := func(a, b [2]int) int {
		if a[0] == b[0] {
			return cmp.Compare(a[1], b[1])
		}
		return cmp.Compare(a[0], b[0])
	}

	cmpFuncs := make(map[[2]int]func([2]int, [2]int) int)
	cmpFuncs[d.up] = cmpHoriz
	cmpFuncs[d.down] = cmpHoriz
	cmpFuncs[d.left] = cmpVert
	cmpFuncs[d.right] = cmpVert

	countHoriz := func(items [][2]int) int {
		n := 1
		for i := 1; i < len(items); i++ {
			if items[i-1][1] != items[i][1] {
				n++
				continue
			}
			if items[i][0]-items[i-1][0] != 1 {
				n++
				continue
			}
		}
		return n
	}
	countVert := func(items [][2]int) int {
		n := 1
		for i := 1; i < len(items); i++ {
			if items[i-1][0] != items[i][0] {
				n++
				continue
			}
			if items[i][1]-items[i-1][1] != 1 {
				n++
				continue
			}
		}
		return n
	}
	countFuncs := make(map[[2]int]func([][2]int) int)
	countFuncs[d.up] = countHoriz
	countFuncs[d.down] = countHoriz
	countFuncs[d.left] = countVert
	countFuncs[d.right] = countVert

	total := 0
	for _, direction := range d.directions {
		slices.SortFunc(sides[direction], cmpFuncs[direction])
		count := countFuncs[direction](sides[direction])
		total += count
	}

	return total
}
