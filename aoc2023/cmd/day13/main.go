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
	total := 0
	patterns := strings.Split(input, "\n\n")
	for _, pattern := range patterns {
		var grid [][]rune
		lines := strings.Split(pattern, "\n")
		for _, line := range lines {
			grid = append(grid, []rune(line))
		}

		// look for horizontal reflection
		if found, row1, row2 := findReflection(grid); found {
			fmt.Println("horizontal reflection found", row1, row2)
			total += 100 * row2
			continue
		}
		fmt.Println("no horizontal reflection found")

		// look for vertical reflection
		tgrid := transpose(grid)
		if found, col1, col2 := findReflection(tgrid); found {
			fmt.Println("vertical reflection found", col1, col2)
			total += col2
			continue
		}
		panic("no horizontal or vertical reflection found")
	}

	slog.Info("Part 1:", "total", total)
}

func findReflection(grid [][]rune) (bool, int, int) {
	for row := 1; row < len(grid); row++ {
		if string(grid[row]) == string(grid[row-1]) {
			reflection := true
			for prevRow, nextRow := row-2, row+1; prevRow >= 0 && nextRow < len(grid); prevRow, nextRow = prevRow-1, nextRow+1 {
				if string(grid[prevRow]) == string(grid[nextRow]) {
					continue
				}
				reflection = false
				break
			}
			if reflection {
				return true, row - 1, row
			}
		}
	}
	return false, 0, 0
}

func transpose(grid [][]rune) [][]rune {
	tgrid := make([][]rune, len(grid[0]))
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			tgrid[col] = append(tgrid[col], grid[row][col])
		}
	}
	return tgrid
}

func part2(input string) {

}
