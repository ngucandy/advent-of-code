package main

import (
	"bufio"
	"log/slog"
	"math"
	"os"
	"time"

	"github.com/ngucandy/advent-of-code/internal/helpers"
	"github.com/ngucandy/advent-of-code/internal/queue"
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
	part2(maze, s, e)
}

var directions = [][2]int{
	{1, 0},  // east
	{0, 1},  // south
	{-1, 0}, // west
	{0, -1}, // north
}

type state struct {
	cost      int
	x         int
	y         int
	direction int
	nodes     [][2]int
}

func (s state) Cost() int {
	return s.cost
}

func part1(maze [][]rune, s [2]int, e [2]int) {
	defer helpers.TrackTime(time.Now())
	pq := queue.PQ[state]{}
	pq.Push(state{0, s[0], s[1], 0, [][2]int{}})
	seen := make(map[[3]int]bool)

	for pq.Len() > 0 {
		st := pq.Pop()
		cost := st.cost
		cx := st.x
		cy := st.y
		cdir := st.direction

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

		forward := state{cost + 1, cx + directions[cdir][0], cy + directions[cdir][1], cdir, st.nodes}
		left := state{cost + 1000, cx, cy, ldir, st.nodes}
		right := state{cost + 1000, cx, cy, rdir, st.nodes}
		if maze[forward.y][forward.x] != '#' {
			pq.Push(forward)
		}
		pq.Push(left)
		pq.Push(right)
	}
}

func part2(maze [][]rune, s [2]int, e [2]int) {
	defer helpers.TrackTime(time.Now())
	pq := queue.PQ[state]{}
	pq.Push(state{0, s[0], s[1], 0, [][2]int{}})
	seen := make(map[[3]int]int)
	bestCost := math.MaxInt
	bestPaths := [][2]int{}

	for pq.Len() > 0 {
		st := pq.Pop()
		cost := st.cost
		cx := st.x
		cy := st.y
		cdir := st.direction

		if [2]int{cx, cy} == e {
			if cost > bestCost {
				continue
			}
			if cost < bestCost {
				bestCost = cost
				bestPaths = make([][2]int, 0)
			}
			slog.Info("Path found:", "cost", cost, "count", len(st.nodes))
			for _, node := range st.nodes {
				maze[node[1]][node[0]] = 'O'
			}
			for _, node := range st.nodes {
				maze[node[1]][node[0]] = '.'
			}
			bestPaths = append(bestPaths, st.nodes...)
			bestPaths = append(bestPaths, [2]int{cx, cy})
			continue
		}
		if seenCost, ok := seen[[3]int{cx, cy, cdir}]; ok && seenCost != cost {
			continue
		}

		seen[[3]int{cx, cy, cdir}] = cost

		ldir := (len(directions) + cdir - 1) % len(directions)
		rdir := (cdir + 1) % len(directions)

		forward := state{cost + 1, cx + directions[cdir][0], cy + directions[cdir][1], cdir, append(st.nodes, [2]int{cx, cy})}
		left := state{cost + 1000, cx, cy, ldir, append([][2]int{}, st.nodes...)}
		right := state{cost + 1000, cx, cy, rdir, append([][2]int{}, st.nodes...)}
		if maze[forward.y][forward.x] != '#' {
			pq.Push(forward)
		}
		pq.Push(left)
		pq.Push(right)
	}

	paths := make(map[[2]int]bool)
	for _, path := range bestPaths {
		paths[path] = true
		maze[path[1]][path[0]] = 'O'
	}
	slog.Info("Part 2:", "tiles", len(paths))
}
