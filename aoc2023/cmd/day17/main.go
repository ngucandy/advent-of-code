package main

import (
	"container/heap"
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
	heap.Push(pq, queue.Item{0, 0, 0, 0, 0, 0})
	seen := make(map[[5]int]bool)
	for pq.Len() > 0 {
		item := heap.Pop(pq).(queue.Item)
		hl, r, c, dr, dc, steps := item[0], item[1], item[2], item[3], item[4], item[5]

		if r == len(grid)-1 && c == len(grid[0])-1 {
			slog.Info("Part 1:", "heat loss", hl)
			break
		}

		key := [5]int{r, c, dr, dc, steps}
		if seen[key] {
			continue
		}
		seen[key] = true

		up := [2]int{-1, 0}
		down := [2]int{1, 0}
		left := [2]int{0, -1}
		right := [2]int{0, 1}

		for _, dir := range [][2]int{up, down, left, right} {
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
				heap.Push(pq, queue.Item{hl + grid[nr][nc], nr, nc, dir[0], dir[1], steps + 1})
			} else {
				heap.Push(pq, queue.Item{hl + grid[nr][nc], nr, nc, dir[0], dir[1], 1})
			}
		}
	}
}

func part2(input string) {

}
