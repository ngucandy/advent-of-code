package main

import (
	"bufio"
	"log/slog"
	"os"
	"regexp"
	"strings"
)

var (
	rexpDigits = regexp.MustCompile(`\d+`)
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
		s := strings.Split(line, ":")
		s = strings.Split(s[1], "|")
		winningNums := rexpDigits.FindAllString(s[0], -1)
		myNums := rexpDigits.FindAllString(s[1], -1)
		slog.Info("Numbers:", "winning", winningNums, "mine", myNums)
		m := make(map[string]bool)
		for _, num := range winningNums {
			m[num] = false
		}
		for _, num := range myNums {
			if _, ok := m[num]; ok {
				m[num] = true
			}
		}
		slog.Info("Matched numbers:", "map", m)

		points := 0
		for _, v := range m {
			if v {
				if points == 0 {
					points = 1
				} else {
					points <<= 1
				}
			}
		}
		slog.Info("Points:", "points", points)
		sum += points
	}
	slog.Info("Total Points:", "sum", sum)
}
