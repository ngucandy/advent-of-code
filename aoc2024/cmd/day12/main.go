package main

import (
	"bufio"
	"cmp"
	"log/slog"
	"os"
	"slices"
)

var (
	DIRECTION_UP    = [2]int{0, -1}
	DIRECTION_DOWN  = [2]int{0, 1}
	DIRECTION_LEFT  = [2]int{-1, 0}
	DIRECTION_RIGHT = [2]int{1, 0}

	DIRECTIONS = [][2]int{DIRECTION_RIGHT, DIRECTION_DOWN, DIRECTION_LEFT, DIRECTION_UP}
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, _ := os.Open(infile)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	grid := [][]rune{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		grid = append(grid, []rune(line))
	}

	part1(grid)
	part2(grid)
}

func part1(grid [][]rune) {
	total := 0
	visited := make(map[[2]int]bool)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if visited[[2]int{x, y}] {
				continue
			}
			a, p := areaPerimeter(x, y, grid[y][x], grid, visited)
			total += a * p
		}
	}
	slog.Info("Part 1:", "total", total)
}

func areaPerimeter(x int, y int, plant rune, grid [][]rune, visited map[[2]int]bool) (area int, perimeter int) {
	if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) || plant != grid[y][x] {
		return 0, 1
	}
	if _, ok := visited[[2]int{x, y}]; ok {
		return 0, 0
	}
	visited[[2]int{x, y}] = true
	area = 1
	perimeter = 0

	for _, direction := range DIRECTIONS {
		aNext, pNext := areaPerimeter(x+direction[0], y+direction[1], plant, grid, visited)
		area += aNext
		perimeter += pNext
	}
	return area, perimeter
}

func part2(grid [][]rune) {
	total := 0
	regions := [][][2]int{}
	visited := make(map[[2]int]bool)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if visited[[2]int{x, y}] {
				continue
			}
			regions = append(regions, fill(x, y, grid[y][x], grid, visited))
		}
	}

	for _, region := range regions {
		m := make(map[[2]int]bool, len(region))
		for _, plot := range region {
			m[plot] = true
		}
		sides := countSides(m)
		area := len(region)
		total += area * sides
	}
	slog.Info("Part 2:", "total", total)
}

func fill(x int, y int, plant rune, grid [][]rune, visited map[[2]int]bool) [][2]int {
	if x < 0 || x >= len(grid[0]) || y < 0 || y >= len(grid) || plant != grid[y][x] {
		return [][2]int{}
	}
	if _, ok := visited[[2]int{x, y}]; ok {
		return [][2]int{}
	}
	visited[[2]int{x, y}] = true
	region := [][2]int{{x, y}}
	for _, direction := range DIRECTIONS {
		region = append(region, fill(x+direction[0], y+direction[1], plant, grid, visited)...)
	}
	return region
}

func countSides(region map[[2]int]bool) int {
	sides := make(map[[2]int][][2]int)
	sides[DIRECTION_UP] = [][2]int{}
	sides[DIRECTION_DOWN] = [][2]int{}
	sides[DIRECTION_RIGHT] = [][2]int{}
	sides[DIRECTION_LEFT] = [][2]int{}

	for plot := range region {
		for _, direction := range DIRECTIONS {
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
	cmpFuncs[DIRECTION_UP] = cmpHoriz
	cmpFuncs[DIRECTION_DOWN] = cmpHoriz
	cmpFuncs[DIRECTION_LEFT] = cmpVert
	cmpFuncs[DIRECTION_RIGHT] = cmpVert

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
	countFuncs[DIRECTION_UP] = countHoriz
	countFuncs[DIRECTION_DOWN] = countHoriz
	countFuncs[DIRECTION_LEFT] = countVert
	countFuncs[DIRECTION_RIGHT] = countVert

	total := 0
	for _, direction := range DIRECTIONS {
		slices.SortFunc(sides[direction], cmpFuncs[direction])
		count := countFuncs[direction](sides[direction])
		total += count
	}

	return total
}
