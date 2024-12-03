package main

import (
	"bufio"
	"log/slog"
	"os"
	"regexp"
	"strconv"
)

var (
	rexpMulInstruction = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
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
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		muls := rexpMulInstruction.FindAllString(line, -1)
		slog.Info("Matched line:", "line", line, "muls", muls)
		for _, mul := range muls {
			match := rexpMulInstruction.FindStringSubmatch(mul)
			slog.Info("Mul instruction:", "mul", mul)
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
