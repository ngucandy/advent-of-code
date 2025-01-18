package aoc2024

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ngucandy/advent-of-code/internal/helpers"
)

func init() {
	Days["1"] = Day1{}
}

type Day1 struct {
}

func (d Day1) Part1(input string) {
	list1, list2 := d.buildLists(input)
	sum := 0
	for i, n := range list1 {
		sum += helpers.AbsInt(n - list2[i])
	}

	fmt.Println("part1", sum)
}

func (d Day1) Part2(input string) {
	list1, list2 := d.buildLists(input)
	list2index := make([]int, list2[len(list2)-1]+1, list2[len(list2)-1]+1)
	for _, n := range list2 {
		list2index[n]++
	}

	score := 0
	for _, n := range list1 {
		score += n * list2index[n]
	}
	fmt.Println("part2", score)
}

func (d Day1) buildLists(input string) ([]int, []int) {
	var list1, list2 []int
	for _, line := range strings.Split(input, "\n") {
		var n1, n2 int
		_, _ = fmt.Sscanf(line, "%d   %d", &n1, &n2)
		list1, list2 = append(list1, n1), append(list2, n2)
	}
	slices.Sort(list1)
	slices.Sort(list2)
	return list1, list2
}
