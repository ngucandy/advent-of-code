package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"slices"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, _ := os.Open(infile)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	maze := [][]rune{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, []rune(line))
	}

	s := [2]int{}
	e := [2]int{}
	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[0]); x++ {
			if maze[y][x] == 'E' {
				e[0] = x
				e[1] = y
			} else if maze[y][x] == 'S' {
				s[0] = x
				s[1] = y
			}
		}
	}

	part1(maze, s, e)
}

var directions = [][2]int{
	{1, 0},  // east
	{0, 1},  // south
	{-1, 0}, // west
	{0, -1}, // north
}

func part1(maze [][]rune, s [2]int, e [2]int) {
	paths := steps(s, e, 0, maze, make(map[[2]int]bool))
	slices.Sort(paths)
	fmt.Println(paths)
}

func steps(start [2]int, end [2]int, dir int, maze [][]rune, visited map[[2]int]bool) []int {
	if start == end {
		return []int{0}
	}

	if maze[start[1]][start[0]] == '#' {
		return []int{}
	}

	if visited[start] {
		return []int{}
	}

	visited[start] = true
	leftDir := (len(directions) + dir - 1) % len(directions)
	rightDir := (dir + 1) % len(directions)
	forward := [2]int{start[0] + directions[dir][0], start[1] + directions[dir][1]}
	left := [2]int{start[0] + directions[leftDir][0], start[1] + directions[leftDir][1]}
	right := [2]int{start[0] + directions[rightDir][0], start[1] + directions[rightDir][1]}

	forwardSteps := steps(forward, end, dir, maze, visited)
	leftSteps := steps(left, end, leftDir, maze, visited)
	rightSteps := steps(right, end, rightDir, maze, visited)

	if len(forwardSteps) == 0 && len(leftSteps) == 0 && len(rightSteps) == 0 {
		visited[start] = false
		return []int{}
	}

	for i := range rightSteps {
		rightSteps[i] += 1001
	}
	for i := range leftSteps {
		leftSteps[i] += 1001
	}
	for i := range forwardSteps {
		forwardSteps[i]++
	}
	visited[start] = false
	combined := append(forwardSteps, append(leftSteps, rightSteps...)...)
	return combined
}
