package aoc2019

import "fmt"

func init() {
	Days["17"] = &Day17{}
}

type Day17 struct {
	eg1, eg2 string
}

func (d Day17) Part1(input string) any {
	comp := NewIntcodeComputer(ParseIntcodeProgram(input), []int{})
	for comp.Step() {
	}
	var grid [][]rune
	var line []rune
	for _, ch := range comp.output {
		if ch == 10 {
			if len(line) == 0 {
				continue
			}
			grid = append(grid, line)
			fmt.Println(string(line))
			line = []rune{}
			continue
		}
		line = append(line, rune(ch))
	}

	sum := 0
	for r, row := range grid {
	row:
		for c, ch := range row {
			if ch == '.' {
				continue
			}
			for _, dir := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				nr, nc := r+dir[0], c+dir[1]
				if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) || grid[nr][nc] == '.' {
					continue row
				}
			}
			sum += r * c
		}
	}
	return sum
}

func (d Day17) Part2(input string) any {
	return "no answer yet"
}
