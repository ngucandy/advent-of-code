package main

import (
	"bufio"
	"log/slog"
	"os"
	"regexp"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, _ := os.Open(infile)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	rexpFrequency := regexp.MustCompile(`[a-zA-Z\d]`)
	index := make(map[rune][][]int)
	grid := [][]rune{}
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
		locations := rexpFrequency.FindAllStringIndex(line, -1)
		for _, location := range locations {
			frequency := rune(line[location[0]])
			if _, ok := index[frequency]; !ok {
				index[frequency] = make([][]int, 0)
			}
			index[frequency] = append(index[frequency], []int{location[0], y})
		}
		y++
	}

	part1(grid, index)
}

func part1(grid [][]rune, index map[rune][][]int) {
	total := 0

	antinodes := make(map[[2]int]bool)
	for _, nodes := range index {
		for i, node1 := range nodes[:len(nodes)-1] {
			for _, node2 := range nodes[i+1:] {
				dx, dy := subtract(node1, node2)
				antinode1 := add(node1, -dx, -dy)
				if !isOutOfBounds(antinode1, grid) {
					antinodes[antinode1] = true
				}
				antinode2 := add(node2, dx, dy)
				if !isOutOfBounds(antinode2, grid) {
					antinodes[antinode2] = true
				}
				slog.Info("Antinodes:", "node1", node1, "node2", node2, "antinode1", antinode1, "antinode2", antinode2)
			}
		}
	}

	for range antinodes {
		total++
	}
	slog.Info("Part 1:", "total", total)
}

func subtract(point1 []int, point2 []int) (int, int) {
	return point2[0] - point1[0], point2[1] - point1[1]
}

func add(point []int, dx, dy int) [2]int {
	return [2]int{point[0] + dx, point[1] + dy}
}

func isOutOfBounds(point [2]int, grid [][]rune) bool {
	return point[0] < 0 || point[0] >= len(grid[0]) || point[1] < 0 || point[1] >= len(grid)
}
