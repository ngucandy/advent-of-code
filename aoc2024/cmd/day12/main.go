package main

import (
	"bufio"
	"log/slog"
	"os"
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

var directions = [][2]int{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
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

	for _, direction := range directions {
		aNext, pNext := areaPerimeter(x+direction[0], y+direction[1], plant, grid, visited)
		area += aNext
		perimeter += pNext
	}
	return area, perimeter
}
