package aoc2024

import (
	"strconv"
	"strings"
)

func init() {
	Days["11"] = Day11{}
}

type Day11 struct {
	example string
}

func (d Day11) Part1(input string) any {
	var stones []int
	for _, stone := range strings.Split(strings.TrimSpace(input), " ") {
		n, _ := strconv.Atoi(stone)
		stones = append(stones, n)
	}
	count := 0
	cache := make(map[[2]int]int)
	for _, stone := range stones {
		count += blink(stone, 25, cache)
	}
	return count
}

func (d Day11) Part2(input string) any {
	var stones []int
	for _, stone := range strings.Split(strings.TrimSpace(input), " ") {
		n, _ := strconv.Atoi(stone)
		stones = append(stones, n)
	}
	count := 0
	cache := make(map[[2]int]int)
	for _, stone := range stones {
		count += blink(stone, 75, cache)
	}
	return count
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
