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
	_, cost := path(s, e, 0, 0, maze, make(map[[2]int]bool))
	slog.Info("Part 1:", "cost", cost)
}

func path(start [2]int, end [2]int, cost int, dir int, maze [][]rune, visited map[[2]int]bool) (bool, int) {
	//fmt.Println(start, end, cost, directions[dir])
	if start == end {
		return true, cost
	}
	if maze[start[1]][start[0]] == '#' {
		return false, 0
	}
	if visited[start] {
		return false, 0
	}

	visited[start] = true
	cost++

	leftDir := (len(directions) + dir - 1) % len(directions)
	rightDir := (dir + 1) % len(directions)
	forward := [2]int{start[0] + directions[dir][0], start[1] + directions[dir][1]}
	left := [2]int{start[0] + directions[leftDir][0], start[1] + directions[leftDir][1]}
	right := [2]int{start[0] + directions[rightDir][0], start[1] + directions[rightDir][1]}

	forwardSuccess, forwardCost := path(forward, end, cost, dir, maze, cloneMap(visited))
	leftSuccess, leftCost := path(left, end, cost+1000, leftDir, maze, cloneMap(visited))
	rightSuccess, rightCost := path(right, end, cost+1000, rightDir, maze, cloneMap(visited))

	if !forwardSuccess && !leftSuccess && !rightSuccess {
		// dead end
		return false, 0
	}

	if forwardSuccess {
		if leftSuccess && rightSuccess {
			return true, min(forwardCost, leftCost, rightCost)
		}
		if leftSuccess {
			return true, min(forwardCost, leftCost)
		}
		if rightSuccess {
			return true, min(forwardCost, rightCost)
		}
		return true, forwardCost
	}
	if leftSuccess {
		if rightSuccess {
			return true, min(leftCost, rightCost)
		}
		return true, leftCost
	}
	return true, rightCost
}

func cloneMap(m map[[2]int]bool) map[[2]int]bool {
	c := make(map[[2]int]bool, len(m))
	for k, v := range m {
		c[k] = v
	}
	return c
}
