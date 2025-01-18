package aoc2024

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func init() {
	Days["5"] = Day5{}
}

type Day5 struct {
}

func (d Day5) Part1(input string) {
	sections := strings.Split(input, "\n\n")
	rules := make(map[string]struct{})
	for _, line := range strings.Split(sections[0], "\n") {
		rules[line] = struct{}{}
	}

	sum := 0
update:
	for _, line := range strings.Split(sections[1], "\n") {
		pages := strings.Split(strings.TrimSpace(line), ",")
		for i, page1 := range pages {
			for _, page2 := range pages[:i] {
				if _, exists := rules[strings.Join([]string{page1, page2}, "|")]; exists {
					// page order vioaltion
					continue update
				}
			}
		}
		// no page order violation
		middle, _ := strconv.Atoi(pages[len(pages)/2])
		sum += middle
	}
	fmt.Println("part1", sum)
}

func (d Day5) Part2(input string) {
	sections := strings.Split(input, "\n\n")
	rules := make(map[string]struct{})
	for _, line := range strings.Split(sections[0], "\n") {
		rules[line] = struct{}{}
	}

	sum := 0
update:
	for _, line := range strings.Split(sections[1], "\n") {
		pages := strings.Split(strings.TrimSpace(line), ",")
		for i, page1 := range pages {
			for _, page2 := range pages[:i] {
				if _, exists := rules[strings.Join([]string{page1, page2}, "|")]; exists {
					// page order vioaltion
					slices.SortStableFunc(pages, func(a, b string) int {
						if _, ok := rules[a+"|"+b]; ok {
							return -1
						}
						if _, ok := rules[b+"|"+a]; ok {
							return 1
						}
						return 0
					})
					middle, _ := strconv.Atoi(pages[len(pages)/2])
					sum += middle
					continue update
				}
			}
		}
	}
	fmt.Println("part2", sum)
}
