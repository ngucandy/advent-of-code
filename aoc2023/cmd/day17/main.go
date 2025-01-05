package main

import (
	"github.com/ngucandy/advent-of-code/internal/queue"
	"log/slog"
	"os"
	"strings"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	bytes, _ := os.ReadFile(infile)
	input := strings.ReplaceAll(string(bytes), "\r\n", "\n")

	part1(input)
	part2(input)

}

var (
	up         = [2]int{-1, 0}
	down       = [2]int{1, 0}
	left       = [2]int{0, -1}
	right      = [2]int{0, 1}
	directions = [][2]int{up, down, left, right}
)

func part1(input string) {
	var grid [][]int
	for _, line := range strings.Split(input, "\n") {
		var row []int
		for _, ch := range line {
			row = append(row, int(ch-'0'))
		}
		grid = append(grid, row)
	}

	pq := queue.NewPriorityQueue()
	queue.Push(pq, queue.Item{0, 0, 0, 0, 0, 0})
	seen := make(map[[5]int]bool)
	for pq.Len() > 0 {
		item := queue.Pop(pq)
		hl, r, c, dr, dc, steps := item[0], item[1], item[2], item[3], item[4], item[5]

		// check if we've reached destination
		if r == len(grid)-1 && c == len(grid[0])-1 {
			slog.Info("Part 1:", "heat loss", hl)
			break
		}

		key := [5]int{r, c, dr, dc, steps}
		if seen[key] {
			continue
		}
		seen[key] = true

		for _, dir := range directions {
			nr := r + dir[0]
			nc := c + dir[1]
			// check bounds
			if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) {
				continue
			}
			// cannot reverse direction
			if (dr != 0 && dir[0] == -dr) || (dc != 0 && dir[1] == -dc) {
				continue
			}
			// check neighbor in the same direction
			if dir[0] == dr && dir[1] == dc {
				// cannot move more than 3 steps in same direction
				if steps == 3 {
					continue
				}
				queue.Push(pq, queue.Item{hl + grid[nr][nc], nr, nc, dir[0], dir[1], steps + 1})
			} else {
				queue.Push(pq, queue.Item{hl + grid[nr][nc], nr, nc, dir[0], dir[1], 1})
			}
		}
	}
}

func part2(input string) {
	var grid [][]int
	for _, line := range strings.Split(input, "\n") {
		var row []int
		for _, ch := range line {
			row = append(row, int(ch-'0'))
		}
		grid = append(grid, row)
	}

	pq := queue.NewPriorityQueue()
	queue.Push(pq, queue.Item{0, 0, 0, 0, 0, 0})
	seen := make(map[[5]int]bool)
	for pq.Len() > 0 {
		item := queue.Pop(pq)
		hl, r, c, dr, dc, steps := item[0], item[1], item[2], item[3], item[4], item[5]

		// check if we've reached destination with at least 4 steps in the same direction
		if r == len(grid)-1 && c == len(grid[0])-1 && steps >= 4 {
			slog.Info("Part 2:", "heat loss", hl)
			break
		}

		key := [5]int{r, c, dr, dc, steps}
		if seen[key] {
			continue
		}
		seen[key] = true

		for _, dir := range directions {
			nr := r + dir[0]
			nc := c + dir[1]
			// check bounds
			if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) {
				continue
			}
			// cannot reverse direction
			if (dr != 0 && dir[0] == -dr) || (dc != 0 && dir[1] == -dc) {
				continue
			}
			// ensure we move at least 4 steps in the same direction
			// special case at beginning when steps == 0
			if steps < 4 && steps > 0 {
				if dir[0] == dr && dir[1] == dc {
					queue.Push(pq, queue.Item{hl + grid[nr][nc], nr, nc, dir[0], dir[1], steps + 1})
				}
				continue
			}
			// check neighbor in the same direction
			if dir[0] == dr && dir[1] == dc {
				// cannot move more than 10 steps in same direction
				if steps == 10 {
					continue
				}
				queue.Push(pq, queue.Item{hl + grid[nr][nc], nr, nc, dir[0], dir[1], steps + 1})
			} else {
				queue.Push(pq, queue.Item{hl + grid[nr][nc], nr, nc, dir[0], dir[1], 1})
			}
		}
	}
}
