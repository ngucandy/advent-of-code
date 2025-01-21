package main

import (
	"fmt"
	"github.com/ngucandy/advent-of-code/aoc2019"
	"github.com/ngucandy/advent-of-code/aoc2024"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type AocDay interface {
	Part1(string) any
	Part2(string) any
}

func main() {
	year := os.Args[1]
	day := os.Args[2]

	var m map[string]interface{}
	switch year {
	case "2019":
		m = aoc2019.Days
	case "2024":
		m = aoc2024.Days
	default:
		panic("unsupported aoc year: " + year)
	}

	aoc, exists := m[day].(AocDay)
	if !exists {
		panic(fmt.Sprintf("unsupported day for aoc%s: %s", year, day))
	}

	cwd, _ := os.Getwd()
	inputLoc := filepath.Join(cwd, "aoc"+year, "inputs", "day"+day+".txt")
	inputBytes, err := os.ReadFile(inputLoc)
	if err != nil {
		panic(err)
	}
	input := strings.ReplaceAll(string(inputBytes), "\r\n", "\n")
	s := time.Now()
	fmt.Println("part1", aoc.Part1(input), time.Since(s))
	s = time.Now()
	fmt.Println("part2", aoc.Part2(input), time.Since(s))
}
