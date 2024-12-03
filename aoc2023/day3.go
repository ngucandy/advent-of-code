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
	rexpGear   = regexp.MustCompile(`\*`)
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

	schematic, parts, gears := buildSchematic(file)

	sum := 0
	for _, part := range parts {
		if isSymbolAdjacent(part, schematic) {
			sum += part.val
		}
	}
	slog.Info("Sum of all part numbers:", "sum", sum)

	sum = 0
	for _, gear := range gears {
		sum += computeRatio(gear, schematic)
	}
	slog.Info("Sum of all gear ratios:", "sum", sum)
}

func buildSchematic(file *os.File) ([][]int, []Part, []Gear) {
	var schematic [][]int
	var parts []Part
	var gears []Gear
	y := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))

		symbolMatches := rexpSymbol.FindAllStringSubmatchIndex(line, -1)
		for _, match := range symbolMatches {
			row[match[0]] = -1
		}

		gearMatches := rexpGear.FindAllStringSubmatchIndex(line, -1)
		for _, match := range gearMatches {
			g := Gear{
				x: match[0],
				y: y,
			}
			gears = append(gears, g)
		}

		partMatches := rexpPart.FindAllStringSubmatchIndex(line, -1)
		for _, match := range partMatches {
			n, err := strconv.Atoi(line[match[0]:match[1]])
			if err != nil {
				panic(err)
			}
			for i := range match[1] - match[0] {
				row[match[0]+i] = n
			}
			p := Part{
				x:   match[0],
				xx:  match[1],
				y:   y,
				val: n,
			}
			parts = append(parts, p)
		}

		schematic = append(schematic, row)
		y++
	}
	return schematic, parts, gears
}

func isSymbolAdjacent(part Part, schematic [][]int) bool {
	for j := max(0, part.y-1); j <= min(len(schematic)-1, part.y+1); j++ {
		for i := max(0, part.x-1); i <= min(len(schematic[j])-1, part.xx); i++ {
			if schematic[j][i] == -1 {
				return true
			}
		}
	}
	return false
}

func computeRatio(gear Gear, schematic [][]int) int {
	ratio := 0
	adjParts := getAdjacentParts(gear, schematic)
	if len(adjParts) == 2 {
		ratio = adjParts[0] * adjParts[1]
	}
	return ratio
}

func getAdjacentParts(gear Gear, schematic [][]int) []int {
	m := make(map[int]bool)
	for j := max(0, gear.y-1); j <= min(len(schematic)-1, gear.y+1); j++ {
		for i := max(0, gear.x-1); i <= min(len(schematic[j])-1, gear.x+1); i++ {
			if schematic[j][i] > 0 {
				m[schematic[j][i]] = true
			}
		}
	}

	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

type Part struct {
	x   int
	xx  int
	y   int
	val int
}

type Gear struct {
	x int
	y int
}
