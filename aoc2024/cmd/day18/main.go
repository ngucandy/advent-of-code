package main

import (
	"bufio"
	"container/heap"
	"fmt"
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

	size := 71
	grid := [][]rune{}
	for i := 0; i < size; i++ {
		row := []rune{}
		for j := 0; j < size; j++ {
			row = append(row, '.')
		}
		grid = append(grid, row)
	}

	bytes := 1024
	scanner := bufio.NewScanner(file)
	for i := 0; i < bytes; i++ {
		scanner.Scan()
		line := scanner.Text()
		var x, y int
		fmt.Sscanf(line, "%d,%d", &x, &y)
		grid[y][x] = '#'
	}
	printGrid(grid)
	part1(grid)
}

type state [5]int

type pqueue []state

func (pq pqueue) Len() int {
	return len(pq)
}

func (pq pqueue) Less(i, j int) bool {
	return pq[i][0] < pq[j][0]
}

func (pq pqueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *pqueue) Push(x interface{}) {
	*pq = append(*pq, x.(state))
}

func (pq *pqueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

var directions = [][2]int{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func part1(grid [][]rune) {
	pq := pqueue{state{0, 0, 0}}
	seen := make(map[[2]int]bool)
	heap.Init(&pq)

	for len(pq) > 0 {
		st := heap.Pop(&pq).(state)
		cost := st[0]
		cx := st[1]
		cy := st[2]

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
			heap.Push(&pq, state{cost + 1, nx, ny})
		}
	}
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}
