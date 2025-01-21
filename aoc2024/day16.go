package aoc2024

import (
	"github.com/ngucandy/advent-of-code/internal/queue"
	"math"
	"slices"
	"strings"
)

type Day16State struct {
	cost  int
	r     int
	c     int
	dir   int
	nodes [][2]int
}

func (s Day16State) Cost() int {
	return s.cost
}

func init() {
	Days["16"] = Day16{
		eg1: `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`,
		eg2: `#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`,
		directions: [][2]int{
			{0, 1},  // east
			{1, 0},  // south
			{0, -1}, // west
			{-1, 0}, // north
		},
	}
}

type Day16 struct {
	eg1, eg2   string
	directions [][2]int
}

func (d Day16) Part1(input string) any {
	var maze [][]rune
	var sr, sc int
	sdir := 0 // east
	for r, line := range strings.Split(input, "\n") {
		maze = append(maze, []rune(line))
		if c := strings.Index(line, "S"); c != -1 {
			sr, sc = r, c
		}
	}

	// use dijkstra to find ANY path with the lowest cost
	pq := queue.PQ[Day16State]{}
	pq.Push(Day16State{0, sr, sc, sdir, [][2]int{}})
	seen := make(map[[3]int]bool)
	for pq.Len() > 0 {
		st := pq.Pop()
		cost := st.cost
		r, c := st.r, st.c
		dir := st.dir

		if maze[r][c] == 'E' {
			return cost
		}
		if seen[[3]int{r, c, dir}] {
			continue
		}
		seen[[3]int{r, c, dir}] = true

		ldir := (len(d.directions) + dir - 1) % len(d.directions)
		rdir := (dir + 1) % len(d.directions)

		forward := Day16State{cost + 1, r + d.directions[dir][0], c + d.directions[dir][1], dir, st.nodes}
		left := Day16State{cost + 1000, r, c, ldir, st.nodes}
		right := Day16State{cost + 1000, r, c, rdir, st.nodes}
		if maze[forward.r][forward.c] != '#' {
			pq.Push(forward)
		}
		pq.Push(left)
		pq.Push(right)
	}
	return nil
}
func (d Day16) Part2(input string) any {
	var maze [][]rune
	var sr, sc int
	sdir := 0 // east
	for r, line := range strings.Split(input, "\n") {
		maze = append(maze, []rune(line))
		if c := strings.Index(line, "S"); c != -1 {
			sr, sc = r, c
		}
	}

	// use dijkstra to find ALL paths with the lowest cost
	pq := queue.PQ[Day16State]{}
	pq.Push(Day16State{0, sr, sc, sdir, [][2]int{}})
	seen := make(map[[3]int]int)
	minCost := math.MaxInt
	var paths [][2]int
	for pq.Len() > 0 {
		st := pq.Pop()
		cost := st.cost
		r, c := st.r, st.c
		dir := st.dir

		if maze[r][c] == 'E' {
			if cost > minCost {
				// found a worst path
				continue
			}
			if cost < minCost {
				// found a better path
				minCost = cost
				paths = make([][2]int, 0)
			}
			// keep track of visited nodes for all best paths
			paths = append(paths, st.nodes...)
			// include end location
			paths = append(paths, [2]int{r, c})
		}
		if seenCost, exists := seen[[3]int{r, c, dir}]; exists && seenCost < cost {
			continue
		}
		seen[[3]int{r, c, dir}] = cost

		ldir := (len(d.directions) + dir - 1) % len(d.directions)
		rdir := (dir + 1) % len(d.directions)

		forward := Day16State{cost + 1, r + d.directions[dir][0], c + d.directions[dir][1], dir, append(st.nodes, [2]int{r, c})}
		left := Day16State{cost + 1000, r, c, ldir, slices.Clone(st.nodes)}
		right := Day16State{cost + 1000, r, c, rdir, slices.Clone(st.nodes)}
		if maze[forward.r][forward.c] != '#' {
			pq.Push(forward)
		}
		pq.Push(left)
		pq.Push(right)
	}

	// filter out duplicate locations from all best paths
	tiles := make(map[[2]int]struct{})
	for _, path := range paths {
		tiles[path] = struct{}{}
	}
	return len(tiles)
}
