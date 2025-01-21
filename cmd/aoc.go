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

	fmt.Printf("AoC %s Day %s\n", year, day)
	fn := map[int]func(string) any{
		0: aoc.Part1,
		1: aoc.Part2,
	}
	for i := range 2 {
		s := time.Now()
		fmt.Printf("part%d %v (%v)\n", i+1, fn[i](input), time.Since(s))
	}
}
