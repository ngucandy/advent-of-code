package main

import (
	"fmt"
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
	var grid [][]rune
	var s [2]int
	startPipe := '-'
	lines := strings.Split(input, "\n")
	for row, line := range lines {
		grid = append(grid, []rune(line))
		if strings.Contains(line, "S") {
			s = [2]int{row, strings.Index(line, "S")}
			grid[s[0]][s[1]] = startPipe
		}
	}
	q := [][3]int{{s[0], s[1], 0}}
	dist := make(map[[2]int]int)
	maxDist := 0
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		if d, ok := dist[[2]int{current[0], current[1]}]; ok && current[2] > d {
			continue
		}
		dist[[2]int{current[0], current[1]}] = current[2]
		if current[2] > maxDist {
			maxDist = current[2]
		}
		switch grid[current[0]][current[1]] {
		case '|':
			q = append(q, [3]int{current[0] - 1, current[1], current[2] + 1})
			q = append(q, [3]int{current[0] + 1, current[1], current[2] + 1})
		case '-':
			q = append(q, [3]int{current[0], current[1] - 1, current[2] + 1})
			q = append(q, [3]int{current[0], current[1] + 1, current[2] + 1})
		case 'L':
			q = append(q, [3]int{current[0] - 1, current[1], current[2] + 1})
			q = append(q, [3]int{current[0], current[1] + 1, current[2] + 1})
		case 'J':
			q = append(q, [3]int{current[0] - 1, current[1], current[2] + 1})
			q = append(q, [3]int{current[0], current[1] - 1, current[2] + 1})
		case '7':
			q = append(q, [3]int{current[0], current[1] - 1, current[2] + 1})
			q = append(q, [3]int{current[0] + 1, current[1], current[2] + 1})
		case 'F':
			q = append(q, [3]int{current[0], current[1] + 1, current[2] + 1})
			q = append(q, [3]int{current[0] + 1, current[1], current[2] + 1})
		default:
			panic(fmt.Sprintf("bad pipe: %v; %s", current, grid[current[0]][current[1]]))
		}
	}

	slog.Info("Part 1:", "max", maxDist)
}

func part2(input string) {

}
