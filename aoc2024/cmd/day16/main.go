package main

import (
	"bufio"
	"container/heap"
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

type state [4]int

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

func part1(maze [][]rune, s [2]int, e [2]int) {
	pq := pqueue{state{0, s[0], s[1], 0}}
	seen := make(map[[3]int]bool)
	heap.Init(&pq)

	for len(pq) > 0 {
		st := heap.Pop(&pq).(state)
		cost := st[0]
		cx := st[1]
		cy := st[2]
		cdir := st[3]

		if [2]int{cx, cy} == e {
			slog.Info("Part 1:", "cost", cost)
			break
		}
		if seen[[3]int{cx, cy, cdir}] {
			continue
		}

		seen[[3]int{cx, cy, cdir}] = true

		ldir := (len(directions) + cdir - 1) % len(directions)
		rdir := (cdir + 1) % len(directions)

		forward := state{cost + 1, cx + directions[cdir][0], cy + directions[cdir][1], cdir}
		left := state{cost + 1000, cx, cy, ldir}
		right := state{cost + 1000, cx, cy, rdir}
		if maze[forward[2]][forward[1]] != '#' {
			heap.Push(&pq, forward)
		}
		heap.Push(&pq, left)
		heap.Push(&pq, right)
	}
}
