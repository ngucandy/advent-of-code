package aoc2019

import (
	"fmt"
	"github.com/ngucandy/advent-of-code/internal/helpers"
)

func init() {
	DayMap["11"] = Day11{}
}

type Day11 struct {
}

func (d Day11) Part1(input string) {
	comp := NewIntcodeComputer(ParseIntcodeProgram(input), []int{0})
	up := [2]int{-1, 0}
	down := [2]int{1, 0}
	left := [2]int{0, -1}
	right := [2]int{0, 1}
	directions := [][2]int{up, right, down, left}
	cdir := 0 // up
	r, c := 0, 0
	grid := make(map[[2]int]int)
	grid[[2]int{0, 0}] = 0
	for comp.Step() {
		if len(comp.output) < 2 {
			continue
		}
		color := comp.output[0]
		turn := comp.output[1]
		comp.output = comp.output[2:]

		grid[[2]int{r, c}] = color
		switch turn {
		case 0: // counterclockwise
			cdir = (cdir - 1 + len(directions)) % len(directions)
		case 1: // clockwise
			cdir = (cdir + 1) % len(directions)
		default:
			panic(fmt.Sprintf("unknown turn value: %d", turn))
		}
		r, c = r+directions[cdir][0], c+directions[cdir][1]
		nextColor := grid[[2]int{r, c}]
		comp.input = append(comp.input, nextColor)
	}
	fmt.Println("part1", len(grid))
}

func (d Day11) Part2(input string) {
	comp := NewIntcodeComputer(ParseIntcodeProgram(input), []int{1})
	up := [2]int{-1, 0}
	down := [2]int{1, 0}
	left := [2]int{0, -1}
	right := [2]int{0, 1}
	directions := [][2]int{up, right, down, left}
	cdir := 0 // up
	r, c := 0, 0
	grid := make(map[[2]int]int)
	grid[[2]int{0, 0}] = 1
	minRow, minCol, maxRow, maxCol := 0, 0, 0, 0
	for comp.Step() {
		if len(comp.output) < 2 {
			continue
		}
		color := comp.output[0]
		turn := comp.output[1]
		comp.output = comp.output[2:]

		if color == 1 {
			grid[[2]int{r, c}] = color
			minRow = min(minRow, r)
			maxRow = max(minRow, r)
			minCol = min(minCol, c)
			maxCol = max(maxCol, c)
		}

		switch turn {
		case 0: // counterclockwise
			cdir = (cdir - 1 + len(directions)) % len(directions)
		case 1: // clockwise
			cdir = (cdir + 1) % len(directions)
		default:
			panic(fmt.Sprintf("unknown turn value: %d", turn))
		}
		r, c = r+directions[cdir][0], c+directions[cdir][1]
		nextColor := grid[[2]int{r, c}]
		comp.input = append(comp.input, nextColor)
	}
	fmt.Println("part2", len(grid), minRow, maxRow, minCol, maxCol)
	var canvas [][]rune
	for range maxRow + 1 {
		var row []rune
		for range maxCol + 1 {
			row = append(row, ' ')
		}
		canvas = append(canvas, row)
	}

	for pos := range grid {
		canvas[pos[0]][pos[1]] = '\u2588'
	}
	helpers.PrintGrid(canvas)
}
