package main

import (
	"bufio"
	"log/slog"
	"os"
	"regexp"
	"strconv"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, _ := os.Open(infile)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	rexpDigits := regexp.MustCompile(`\d+`)
	machines := [][][2]float64{}
	machine := [][2]float64{}
	scanner := bufio.NewScanner(file)
	for {
		for range 3 {
			scanner.Scan()
			line := scanner.Text()
			nums := rexpDigits.FindAllString(line, -1)
			x, _ := strconv.ParseFloat(nums[0], 64)
			y, _ := strconv.ParseFloat(nums[1], 64)
			machine = append(machine, [2]float64{x, y})
		}
		machines = append(machines, machine)
		if !scanner.Scan() {
			break
		}
		machine = [][2]float64{}
	}

	part1(machines)
	part2(machines)
}

func part1(machines [][][2]float64) {
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
			//slog.Error("Invalid number of B presses:", "presses", pressesB)
			continue
		}
		pressesA := ((machine[p][x]) - (pressesB * machine[b][x])) / (machine[a][x])
		if !isWholeNumber(pressesA) {
			// reject fractional button presses
			//slog.Error("Invalid number of A presses:", "presses", pressesB)
			continue
		}
		tokens += (pressesA * 3.0) + pressesB
	}
	slog.Info("Part 1:", "tokens", int64(tokens))
}

func part2(machines [][][2]float64) {
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
			//slog.Error("Invalid number of B presses:", "presses", pressesB)
			continue
		}
		pressesA := ((machine[p][x]) - (pressesB * machine[b][x])) / (machine[a][x])
		if !isWholeNumber(pressesA) {
			// reject fractional button presses
			//slog.Error("Invalid number of A presses:", "presses", pressesB)
			continue
		}
		tokens += (pressesA * 3.0) + pressesB
	}
	slog.Info("Part 2:", "tokens", int64(tokens))
}

func isWholeNumber(number float64) bool {
	return number == float64(int64(number))
}
