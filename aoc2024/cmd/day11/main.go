package main

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ngucandy/advent-of-code/internal/helpers"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, _ := os.Open(infile)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	input := strings.Split(line, " ")

	part1(input)
	part2(input)
}

func part1(input []string) {
	defer helpers.TrackTime(time.Now(), "Part 1")
	stones := input

	for range 25 {
		next := make([]string, 0, len(stones))
		for _, stone := range stones {
			if stone == "0" {
				next = append(next, "1")
				continue
			}
			if len(stone)%2 == 0 {
				next = append(next, stone[:len(stone)/2])
				n, _ := strconv.Atoi(stone[len(stone)/2:])
				next = append(next, strconv.Itoa(n))
				continue
			}
			n, _ := strconv.Atoi(stone)
			next = append(next, strconv.Itoa(n*2024))
		}
		stones = next
	}

	slog.Info("Part 1:", "stones", len(stones))
}

func part2(input []string) {
	defer helpers.TrackTime(time.Now(), "Part 2")

	stones := make([]int, 0, len(input))
	for _, stone := range input {
		n, _ := strconv.Atoi(stone)
		stones = append(stones, n)
	}

	count := 0
	cache := make(map[[2]int]int)
	for _, stone := range stones {
		count += blink(stone, 75, cache)
	}
	slog.Info("Part 2:", "stones", count)
}

func blink(stone int, times int, cache map[[2]int]int) int {
	if times == 0 {
		return 1
	}
	if count, ok := cache[[2]int{stone, times}]; ok {
		return count
	}

	if stone == 0 {
		cache[[2]int{stone, times}] = blink(1, times-1, cache)
		return cache[[2]int{stone, times}]
	}
	s := strconv.Itoa(stone)
	if len(s)%2 == 0 {
		n1, _ := strconv.Atoi(s[:len(s)/2])
		n2, _ := strconv.Atoi(s[len(s)/2:])
		cache[[2]int{stone, times}] = blink(n1, times-1, cache) + blink(n2, times-1, cache)
		return cache[[2]int{stone, times}]
	}
	cache[[2]int{stone, times}] = blink(stone*2024, times-1, cache)
	return cache[[2]int{stone, times}]
}
