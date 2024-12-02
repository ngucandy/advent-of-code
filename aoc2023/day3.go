package main

import (
	"bufio"
	"log/slog"
	"os"
	"regexp"
	"strconv"
)

var (
	rexpPart   = regexp.MustCompile(`\d+`)
	rexpSymbol = regexp.MustCompile(`[^.\d]`)
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			slog.Error("Error closing file:", "error", err)
		}
	}(file)

	var schematic [][]int
	var parts []Part
	y := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))

		symbolMatches := rexpSymbol.FindAllStringSubmatchIndex(line, -1)
		for _, match := range symbolMatches {
			row[match[0]] = 1
		}

		partMatches := rexpPart.FindAllStringSubmatchIndex(line, -1)
		for _, match := range partMatches {
			for i := range match[1] - match[0] {
				row[match[0]+i] = 2
			}
			n, err := strconv.ParseInt(line[match[0]:match[1]], 10, 32)
			if err != nil {
				panic(err)
			}
			p := Part{
				x:   match[0],
				xx:  match[1],
				y:   y,
				val: int(n),
			}
			parts = append(parts, p)
		}

		schematic = append(schematic, row)
		y++
	}

	sum := 0
	for _, part := range parts {
		if isSymbolAdjacent(part, schematic) {
			sum += part.val
		}
	}
	slog.Info("Sum of all part numbers:", "sum", sum)
}

func isSymbolAdjacent(part Part, schematic [][]int) bool {
	for j := max(0, part.y-1); j <= min(len(schematic)-1, part.y+1); j++ {
		for i := max(0, part.x-1); i <= min(len(schematic[j])-1, part.xx); i++ {
			if schematic[j][i] == 1 {
				return true
			}
		}
	}
	return false
}

type Part struct {
	x   int
	xx  int
	y   int
	val int
}
