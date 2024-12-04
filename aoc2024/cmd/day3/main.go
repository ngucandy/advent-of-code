package main

import (
	"bufio"
	"log/slog"
	"os"
	"regexp"
	"strconv"
)

var (
	rexpMulDoDont       = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	rexpMulInstruction  = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	rexpDoInstruction   = regexp.MustCompile(`do\(\)`)
	rexpDontInstruction = regexp.MustCompile(`don't\(\)`)
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

	sum := 0
	enabled := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matchedInstructions := rexpMulDoDont.FindAllString(line, -1)
		slog.Info("Matched line:", "line", line, "matchedInstructions", matchedInstructions)
		for _, instruction := range matchedInstructions {
			if rexpDoInstruction.MatchString(instruction) {
				enabled = true
				continue
			}
			if rexpDontInstruction.MatchString(instruction) {
				enabled = false
				continue
			}
			if !enabled {
				slog.Info("Skipping instruction:", "instruction", instruction)
				continue
			}
			match := rexpMulInstruction.FindStringSubmatch(instruction)
			slog.Info("Mul instruction:", "instruction", instruction)
			x, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(match[2])
			if err != nil {
				panic(err)
			}
			sum += x * y
		}
	}
	slog.Info("Sum of mul instructions:", "sum", sum)
}
