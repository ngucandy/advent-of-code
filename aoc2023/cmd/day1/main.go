package main

import (
	"bufio"
	"log/slog"
	"os"
	"regexp"
)

var table = map[string]int{
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

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

	// watch out for these: `oneight`, `eightwo`, `twone`
	releft := regexp.MustCompile("([\\d]|one|two|three|four|five|six|seven|eight|nine).*")
	reright := regexp.MustCompile(".*([\\d]|one|two|three|four|five|six|seven|eight|nine)")

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		match := releft.FindStringSubmatch(line)
		first := match[1]
		digit, ok := table[first]
		if !ok {
			panic(first)
		}
		val := 10 * digit

		match = reright.FindStringSubmatch(line)
		last := match[1]
		digit, ok = table[last]
		if !ok {
			panic(last)
		}
		val += digit

		//slog.Info("Digits:", "line", line, "first", first, "last", last, "val", val)
		sum += val
	}
	slog.Info("Sum of calibration values:", "sum", sum)
}
