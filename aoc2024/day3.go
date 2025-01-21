package aoc2024

import (
	"regexp"
	"strconv"
	"strings"
)

func init() {
	Days["3"] = Day3{}
}

type Day3 struct {
}

func (d Day3) Part1(input string) any {
	sum := 0
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	for _, line := range strings.Split(input, "\n") {
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			n1, _ := strconv.Atoi(match[1])
			n2, _ := strconv.Atoi(match[2])
			sum += n1 * n2
		}
	}
	return sum
}

func (d Day3) Part2(input string) any {
	sum := 0
	enabled := true
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	for _, line := range strings.Split(input, "\n") {
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			if strings.HasPrefix(match[0], "don") {
				enabled = false
				continue
			}
			if strings.HasPrefix(match[0], "do(") {
				enabled = true
				continue
			}
			if enabled {
				n1, _ := strconv.Atoi(match[1])
				n2, _ := strconv.Atoi(match[2])
				sum += n1 * n2
				continue
			}
		}
	}
	return sum
}
