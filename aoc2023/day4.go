package main

import (
	"bufio"
	"log/slog"
	"math"
	"os"
	"regexp"
	"strings"
	"sync"
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

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go part1(file, wg)
	//go part2(file, wg)
	wg.Wait()
}

func part1(file *os.File, wg *sync.WaitGroup) {
	defer wg.Done()
	points := 0
	counts := countWinningNumbers(file)
	for _, count := range counts {
		points += int(math.Pow(2, float64(count-1)))
	}
	slog.Info("Total Points:", "points", points)
}

func countWinningNumbers(file *os.File) map[string]int {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ":")
		cardId := rexpDigits.FindString(s[0])
		s = strings.Split(s[1], "|")
		winningNums := rexpDigits.FindAllString(s[0], -1)
		myNums := rexpDigits.FindAllString(s[1], -1)
		m := make(map[string]bool)
		for _, num := range winningNums {
			m[num] = false
		}
		winningCount := 0
		for _, num := range myNums {
			if _, ok := m[num]; ok {
				winningCount++
			}
		}
		counts[cardId] = winningCount
	}
	return counts
}
