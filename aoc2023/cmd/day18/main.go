package main

import (
	"fmt"
	"github.com/ngucandy/advent-of-code/internal/helpers"
	"log/slog"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const (
	testInput = `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	bytes, _ := os.ReadFile(infile)
	input := strings.ReplaceAll(string(bytes), "\r\n", "\n")

	part1(testInput)
	part1(input)
	part2(testInput)
	part2(input)
}

func part1(input string) {
	// tr, tc are row and column for a point on the trench
	grid, trenchSize, tr, tc := buildGrid(input)

	// for each location around the trench point, do a bfs flood-fill;
	// if the flood fill reaches any edge of the grid, then it must be
	// outside of the trench area; if we can visit every node in an area
	// without hitting a grid edge, then it must be bounded by the trench
TrenchNeighbors:
	for _, dir := range [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}} {
		nr := tr + dir[0]
		nc := tc + dir[1]
		if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) {
			continue
		}
		if grid[nr][nc] == '#' {
			continue
		}
		q := [][2]int{{nr, nc}}
		seen := make(map[[2]int]bool)
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			r := cur[0]
			c := cur[1]
			// if we've reached a grid edge, then we're outside the trench
			if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
				continue TrenchNeighbors
			}
			if grid[r][c] == '#' {
				continue
			}
			if seen[cur] {
				continue
			}
			seen[cur] = true
			for _, d := range [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
				q = append(q, [2]int{r + d[0], c + d[1]})
			}
		}
		// we've visited every node in the q without hitting any edge of the grid
		slog.Info("Part 1:", "capacity", trenchSize+len(seen))
	}
}

func buildGrid(input string) (grid [][]rune, trenches, r, c int) {
	grid = [][]rune{{'#'}}
	r, c = 0, 0
	trenches = 0
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, " ")
		dir := parts[0]
		n, _ := strconv.Atoi(parts[1])
		trenches += n
		switch dir {
		case "R":
			if c+n >= len(grid[0]) {
				for row := range grid {
					expand := c + n - (len(grid[row]) - 1)
					grid[row] = append(grid[row], slices.Repeat([]rune{'.'}, expand)...)
				}
			}
			for range n {
				grid[r][c+1] = '#'
				c++
			}
		case "L":
			if c-n < 0 {
				expand := 0 - c + n
				for row := range grid {
					grid[row] = append(slices.Repeat([]rune{'.'}, expand), grid[row]...)
				}
				c += expand
			}
			for range n {
				grid[r][c-1] = '#'
				c--
			}
		case "U":
			if r-n < 0 {
				expand := 0 - r + n
				for range expand {
					grid = append([][]rune{slices.Repeat([]rune{'.'}, len(grid[0]))}, grid...)
				}
				r += expand
			}
			for range n {
				grid[r-1][c] = '#'
				r--
			}
		case "D":
			if r+n >= len(grid) {
				expand := r + n - (len(grid) - 1)
				for range expand {
					grid = append(grid, slices.Repeat([]rune{'.'}, len(grid[0])))
				}
			}
			for range n {
				grid[r+1][c] = '#'
				r++
			}
		}
	}
	return grid, trenches, r, c
}

func part2(input string) {
	// run through input and save corner locations
	re := regexp.MustCompile(`#(.....)(.)`)
	r, c := 0, 0
	corners := [][2]int{{r, c}}
	// also track min and max row to compute final grid height
	minRow, maxRow := 0, 0
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		match := re.FindStringSubmatch(line)
		n, _ := strconv.ParseInt(match[1], 16, 64)
		switch match[2] {
		case "0": // right
			c += int(n)
		case "1": // down
			r += int(n)
			maxRow = max(maxRow, r)
		case "2": // left
			c -= int(n)
		case "3": // up
			r -= int(n)
			minRow = min(minRow, r)
		}
		corners = append(corners, [2]int{r, c})
	}

	// find a corner at the "top" (minRow) of the grid
	c1Index := slices.IndexFunc(corners, func(c [2]int) bool {
		return c[0] == minRow
	})
	// if the next corner is to the right, then we were moving clockwise
	// around the perimeter when it was added
	c2Index := (c1Index + 1) % len(corners)
	clockwise := corners[c1Index][1] < corners[c2Index][1]
	fmt.Println("moving clockwise?", clockwise)
	area := 0
	perimeter := 0
	gridHeight := maxRow - minRow
	fmt.Println("grid height", gridHeight)

	// compute the area by adding or subtracting the area of the rectangle
	// formed from two corners that form a horizontal line; corners that
	// form a vertical line will have an area of 0; ranging over the corners
	// in creation order ensures we are looking at corners that are connected
	// either by a vertical or horizontal line
	for i, n := 0, len(corners); i < n; i++ {
		c1 := corners[(i+c1Index)%n]
		c2 := corners[(i+c1Index+1)%n]
		if clockwise {
			area += (gridHeight - c1[0]) * (c2[1] - c1[1])
		} else {
			area += (gridHeight - c1[0]) * (c1[1] - c2[1])
		}
		perimeter += helpers.AbsInt((c2[1] - c1[1]) + (c2[0] - c1[0]))
	}
	// area calculation includes half of the perimeter, so need to add back the other half
	// +1 because int division rounds down
	perimeter = (perimeter / 2) + 1
	slog.Info("Part 2:", "capacity", area+perimeter)
}
