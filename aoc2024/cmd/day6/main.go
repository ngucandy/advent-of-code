package main

import (
	"bufio"
	"log/slog"
	"os"
	"regexp"
)

var (
	directions = [][]int{
		{0, -1}, // up
		{1, 0},  // right
		{0, 1},  // down
		{-1, 0}, // left
	}
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			slog.Error("Error closing file:", "error", err)
		}
	}(file)

	rexpObstruction := regexp.MustCompile(`#`)
	rexpGuard := regexp.MustCompile(`\^`)
	obstructions := [][]int{}
	guardLoc := [2]int{-1, -1}
	direction := 0
	visited := [][]int{}
	y := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		visited = append(visited, make([]int, len(line)))

		// locate obstructions
		obstructions = append(obstructions, make([]int, len(line)))
		obstructionIndices := rexpObstruction.FindAllStringIndex(line, -1)
		for _, i := range obstructionIndices {
			obstructions[y][i[0]] = 1
		}

		// locate guard
		loc := rexpGuard.FindStringIndex(line)
		if loc != nil {
			guardLoc[0] = loc[0]
			guardLoc[1] = y
			visited[y][loc[0]] = 1
		}

		y++
	}
	part1(obstructions, guardLoc, direction, visited)
}

func part1(obstructions [][]int, guardLoc [2]int, d int, visited [][]int) {
	for {
		newLoc := [2]int{guardLoc[0] + directions[d][0], guardLoc[1] + directions[d][1]}

		if newLoc[0] < 0 || newLoc[0] == len(visited[0]) || newLoc[1] < 0 || newLoc[1] == len(visited) {
			break
		}

		if obstructions[newLoc[1]][newLoc[0]] == 1 {
			d = (d + 1) % len(directions)
			slog.Info("Changed direction:", "new direction", d)
			continue
		}

		guardLoc = newLoc
		visited[newLoc[1]][newLoc[0]] = 1
	}

	total := 0
	for _, j := range visited {
		for _, i := range j {
			total += i
		}
	}

	slog.Info("Total visited:", "total", total)
}
