package aoc2024

import (
	"fmt"
	"strings"
)

func init() {
	Days["13"] = Day13{}
}

type Day13 struct {
	example string
}

func (d Day13) Part1(input string) {
	var machines [][][2]float64
	for _, section := range strings.Split(input, "\n\n") {
		// Button A: X+94, Y+34
		// Button B: X+22, Y+67
		// Prize: X=8400, Y=5400
		lines := strings.Split(section, "\n")
		var ax, ay, bx, by, px, py int
		_, _ = fmt.Sscanf(lines[0], "Button A: X%d, Y%d", &ax, &ay)
		_, _ = fmt.Sscanf(lines[1], "Button B: X%d, Y%d", &bx, &by)
		_, _ = fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &px, &py)
		machines = append(machines, [][2]float64{
			{float64(ax), float64(ay)},
			{float64(bx), float64(by)},
			{float64(px), float64(py)},
		})
	}

	tokens := 0.0

	a := 0
	b := 1
	p := 2
	x := 0
	y := 1

	// let:
	//   a = number of button A presses
	//   b = number of button B presses
	//   Px = X coordinate of prize
	//   Py = Y coordinate of prize
	//   Ax = movement along X axis for pressing button A
	//   Ay = movement along Y axis for pressing button A
	//   Bx = movement along X axis for pressing button B
	//   By = movement along Y axis for pressing button B
	//
	// system of equations to solve
	//   a * Ax + b * Bx = Px
	//   a * Ay + b * By = Py
	//
	// solving for a in first equation:
	//   a = (Px - (b * Bx)) / Ax
	//
	// solving for b in second equation substituting for a:
	//   b = ((Ax * Py) - (Px * Ay)) / ((Ax * By) - (Bx * Ay))
	for _, machine := range machines {
		pressesB := ((machine[a][x] * machine[p][y]) - (machine[p][x] * machine[a][y])) / ((machine[a][x] * machine[b][y]) - (machine[b][x] * machine[a][y]))
		if !isWholeNumber(pressesB) {
			// reject fractional button presses
			continue
		}
		pressesA := ((machine[p][x]) - (pressesB * machine[b][x])) / (machine[a][x])
		if !isWholeNumber(pressesA) {
			// reject fractional button presses
			continue
		}
		tokens += (pressesA * 3.0) + pressesB
	}
	fmt.Println("part1", int(tokens))
}

func (d Day13) Part2(input string) {
	var machines [][][2]float64
	for _, section := range strings.Split(input, "\n\n") {
		// Button A: X+94, Y+34
		// Button B: X+22, Y+67
		// Prize: X=8400, Y=5400
		lines := strings.Split(section, "\n")
		var ax, ay, bx, by, px, py int
		_, _ = fmt.Sscanf(lines[0], "Button A: X%d, Y%d", &ax, &ay)
		_, _ = fmt.Sscanf(lines[1], "Button B: X%d, Y%d", &bx, &by)
		_, _ = fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &px, &py)
		machines = append(machines, [][2]float64{
			{float64(ax), float64(ay)},
			{float64(bx), float64(by)},
			{float64(px), float64(py)},
		})
	}

	tokens := 0.0

	a := 0
	b := 1
	p := 2
	x := 0
	y := 1

	for _, machine := range machines {
		machine[p][x] += 10000000000000
		machine[p][y] += 10000000000000
		pressesB := ((machine[a][x] * machine[p][y]) - (machine[p][x] * machine[a][y])) / ((machine[a][x] * machine[b][y]) - (machine[b][x] * machine[a][y]))
		if !isWholeNumber(pressesB) {
			// reject fractional button presses
			continue
		}
		pressesA := ((machine[p][x]) - (pressesB * machine[b][x])) / (machine[a][x])
		if !isWholeNumber(pressesA) {
			// reject fractional button presses
			continue
		}
		tokens += (pressesA * 3.0) + pressesB
	}
	fmt.Println("part2", int(tokens))
}

func isWholeNumber(number float64) bool {
	return number == float64(int64(number))
}
