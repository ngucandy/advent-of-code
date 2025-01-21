package aoc2024

import (
	"strings"
)

func init() {
	Days["25"] = Day25{
		eg1: `#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####`,
	}
}

type Day25 struct {
	eg1, eg2 string
}

func (d Day25) Part1(input string) any {
	var locks, keys [][5]int
	for _, section := range strings.Split(input, "\n\n") {
		heights := [5]int{-1, -1, -1, -1, -1}
		for _, line := range strings.Split(section, "\n") {
			for i, char := range line {
				if char == '#' {
					heights[i]++
				}
			}
		}
		if section[:5] == "#####" {
			locks = append(locks, heights)
		} else {
			keys = append(keys, heights)
		}
	}

	count := 0
	for _, lock := range locks {
	KeyLoop:
		for _, key := range keys {
			for i := range key {
				if lock[i]+key[i] > 5 {
					continue KeyLoop
				}
			}
			count++
		}
	}
	return count
}

func (d Day25) Part2(_ string) any {
	return "done!"
}
