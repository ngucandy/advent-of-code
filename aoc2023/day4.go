package main

import (
	"bufio"
	"log/slog"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
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

	counts := countWinningNumbers(file)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go part1(counts, wg)
	go part2(counts, wg)
	wg.Wait()
}

func part2(counts map[int]int, wg *sync.WaitGroup) {
	defer wg.Done()
	cards := make(map[int]int)
	var ids []int
	for id := range counts {
		cards[id] = 1
		ids = append(ids, id)
	}
	sort.Ints(ids)
	for _, id := range ids {
		for i := 1; i <= counts[id]; i++ {
			cards[id+i] += cards[id]
		}
	}
	sum := 0
	for _, cardCount := range cards {
		sum += cardCount
	}
	slog.Info("Part 2:", "total cards", sum)
}

func part1(counts map[int]int, wg *sync.WaitGroup) {
	defer wg.Done()
	points := 0
	for _, count := range counts {
		points += int(math.Pow(2, float64(count-1)))
	}
	slog.Info("Part 1:", "total points", points)
}

func countWinningNumbers(file *os.File) map[int]int {
	counts := make(map[int]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Split(line, ":")
		cardId, err := strconv.Atoi(rexpDigits.FindString(s[0]))
		if err != nil {
			panic(err)
		}
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
