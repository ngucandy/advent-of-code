package aoc2024

import (
	"slices"
	"strconv"
	"strings"

	"github.com/ngucandy/advent-of-code/internal/helpers"
)

func init() {
	Days["2"] = Day2{}
}

type Day2 struct {
}

func (d Day2) Part1(input string) any {
	count := 0
	for _, line := range strings.Split(input, "\n") {
		var report []int
		for _, part := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(part)
			report = append(report, n)
		}
		if d.safe(report) {
			count++
		}
	}
	return count
}

func (d Day2) Part2(input string) any {
	count := 0
report:
	for _, line := range strings.Split(input, "\n") {
		var report []int
		for _, part := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(part)
			report = append(report, n)
		}
		if d.safe(report) {
			count++
			continue
		}
		for i := range report {
			if d.safe(slices.Concat(report[:i], report[i+1:])) {
				count++
				continue report
			}
		}
	}
	return count
}

func (d Day2) safe(report []int) bool {
	var diffs []int
	for i := 1; i < len(report); i++ {
		diffs = append(diffs, report[i]-report[i-1])
	}
	ascending := 0
	for _, diff := range diffs {
		switch helpers.AbsInt(diff) {
		case 1, 2, 3:
			break
		default:
			return false
		}
		if diff > 0 {
			ascending++
		}
	}
	return ascending == 0 || ascending == len(diffs)
}
