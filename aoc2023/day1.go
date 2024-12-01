package main

import (
	"bufio"
	"log/slog"
	"os"
	"regexp"
	"strconv"
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
	re := regexp.MustCompile("[0-9]")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		digits := re.FindAllString(line, -1)
		digit, err := strconv.ParseInt(digits[0], 10, 32)
		if err != nil {
			panic(err)
		}
		val := 10 * int(digit)
		digit, err = strconv.ParseInt(digits[len(digits)-1], 10, 32)
		if err != nil {
			panic(err)
		}
		val += int(digit)
		sum += val
	}
	slog.Info("Sum of calibration values:", "sum", sum)
}
