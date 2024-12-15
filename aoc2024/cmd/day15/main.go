package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"regexp"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, _ := os.Open(infile)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	rexpMovement := regexp.MustCompile(`[<>v^]`)
	rexpRobot := regexp.MustCompile(`@`)
	scanner := bufio.NewScanner(file)
	movements := ""
	start := [2]int{}
	grid := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if rexpMovement.MatchString(line) {
			movements += line
			continue
		}
		grid = append(grid, []rune(line))
		if rexpRobot.MatchString(line) {
			start[0] = rexpRobot.FindStringIndex(line)[1] - 1
			start[1] = len(grid) - 1
		}
	}

	part1(grid, movements, start)
}

func part1(grid [][]rune, movements string, start [2]int) {
	fmt.Println(start, movements)
	printGrid(grid)
	x := start[0]
	y := start[1]
	for _, dir := range movements {
		if move(x, y, dir, grid) {
			x += directions[dir][0]
			y += directions[dir][1]
		}
	}
	printGrid(grid)

	sum := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 'O' {
				sum += 100*row + col
			}
		}
	}
	slog.Info("Part 1:", "sum", sum)
}

func printGrid(grid [][]rune) {
	for row := range grid {
		fmt.Println(string(grid[row]))
	}
}

var directions = map[rune][2]int{
	'^': {0, -1},
	'v': {0, 1},
	'<': {-1, 0},
	'>': {1, 0},
}

func move(x int, y int, dir rune, grid [][]rune) bool {
	nextX := x + directions[dir][0]
	nextY := y + directions[dir][1]
	if grid[nextY][nextX] == '#' {
		return false
	}
	if grid[nextY][nextX] == '.' {
		grid[nextY][nextX] = grid[y][x]
		grid[y][x] = '.'
		return true
	}
	// try to move obstacle (recursively)
	if move(nextX, nextY, dir, grid) {
		grid[nextY][nextX] = grid[y][x]
		grid[y][x] = '.'
		return true
	}
	// obstacle cannot be moved
	return false
}
