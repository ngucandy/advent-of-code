package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log/slog"
	"os"

	"github.com/ngucandy/advent-of-code/internal/queue"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, _ := os.Open(infile)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	size := 71
	grid := [][]rune{}
	for i := 0; i < size; i++ {
		row := []rune{}
		for j := 0; j < size; j++ {
			row = append(row, '.')
		}
		grid = append(grid, row)
	}

	bytes := [][2]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var x, y int
		_, _ = fmt.Sscanf(line, "%d,%d", &x, &y)
		bytes = append(bytes, [2]int{x, y})
	}
	part1(grid, bytes)
	part2(grid, bytes)
}

var directions = [][2]int{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func part1(grid [][]rune, bytes [][2]int) {
	fall(grid, bytes, 1024)
	pq := queue.NewPriorityQueue()
	pq.Push(queue.Item{0, 0, 0})
	seen := make(map[[2]int]bool)

	for pq.Len() > 0 {
		item := heap.Pop(pq).(queue.Item)
		cost := item[0]
		cx := item[1]
		cy := item[2]

		if cx == len(grid)-1 && cy == len(grid)-1 {
			slog.Info("Part 1:", "steps", cost)
			break
		}

		if seen[[2]int{cx, cy}] {
			continue
		}

		seen[[2]int{cx, cy}] = true

		for _, direction := range directions {
			nx := cx + direction[0]
			ny := cy + direction[1]
			if nx < 0 || nx >= len(grid) || ny < 0 || ny >= len(grid) || grid[ny][nx] == '#' {
				continue
			}
			heap.Push(pq, queue.Item{cost + 1, nx, ny})
		}
	}
}

func part2(grid [][]rune, bytes [][2]int) {
	fall(grid, bytes, 1024)
	for i := 1024; i < len(bytes); i++ {
		grid[bytes[i][0]][bytes[i][1]] = '#'

		pq := queue.NewPriorityQueue()
		pq.Push(queue.Item{0, 0, 0})
		seen := make(map[[2]int]bool)
		pathFound := false

		for pq.Len() > 0 {
			item := heap.Pop(pq).(queue.Item)
			cost := item[0]
			cx := item[1]
			cy := item[2]

			if cx == len(grid)-1 && cy == len(grid)-1 {
				slog.Info("Found path to end:", "steps", cost)
				pathFound = true
				break
			}

			if seen[[2]int{cx, cy}] {
				continue
			}

			seen[[2]int{cx, cy}] = true

			for _, direction := range directions {
				nx := cx + direction[0]
				ny := cy + direction[1]
				if nx < 0 || nx >= len(grid) || ny < 0 || ny >= len(grid) || grid[ny][nx] == '#' {
					continue
				}
				heap.Push(pq, queue.Item{cost + 1, nx, ny})
			}
		}
		if !pathFound {
			slog.Info("Part 2:", "byte", bytes[i])
			break
		}
	}
}

func fall(grid [][]rune, bytes [][2]int, n int) {
	for i := 0; i < n; i++ {
		grid[bytes[i][0]][bytes[i][1]] = '#'
	}
}
