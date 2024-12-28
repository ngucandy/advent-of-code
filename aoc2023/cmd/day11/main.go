package main

import (
	"log/slog"
	"os"
	"regexp"
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
	re := regexp.MustCompile(`#`)
	var grid [][]rune
	var galaxies [][2]int
	galaxyRows := make(map[int]int)
	galaxyCols := make(map[int]int)
	lines := strings.Split(input, "\n")
	for row, line := range lines {
		grid = append(grid, []rune(line))
		if re.MatchString(line) {
			matches := re.FindAllStringIndex(line, -1)
			for _, match := range matches {
				galaxies = append(galaxies, [2]int{row, match[0]})
				galaxyRows[row]++
				galaxyCols[match[0]]++
			}
		}
	}

	total := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			row1 := galaxies[i][0]
			col1 := galaxies[i][1]
			row2 := galaxies[j][0]
			col2 := galaxies[j][1]
			distance := (max(row1, row2) - min(row1, row2)) + (max(col1, col2) - min(col1, col2))
			for row := min(row1, row2) + 1; row < max(row1, row2); row++ {
				if _, exists := galaxyRows[row]; !exists {
					distance++
				}
			}
			for col := min(col1, col2) + 1; col < max(col1, col2); col++ {
				if _, exists := galaxyCols[col]; !exists {
					distance++
				}
			}
			total += distance
		}
	}
	slog.Info("Part 1:", "total", total)
}

func part2(input string) {

}
