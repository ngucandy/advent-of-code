package aoc2019

import (
	"fmt"
	"github.com/ngucandy/advent-of-code/internal/helpers"
	"strings"
)

func init() {
	DayMap["10"] = Day10{
		`.#..#
.....
#####
....#
...##`,
		`......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`,
		`#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`,
		`.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`,
		`.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`,
	}
}

type Day10 struct {
	eg1, eg2, eg3, eg4, eg5 string
}

func (d Day10) Part1(input string) {
	visible := make(map[[2]int]int)
	asteroids := make([][2]int, 0)
	gh, gw := 0, 0
	for r, line := range strings.Split(input, "\n") {
		for c, ch := range line {
			if ch == '#' {
				pos := [2]int{r, c}
				visible[pos] = 0
				asteroids = append(asteroids, pos)
			}
			gw = c + 1
		}
		gh = r + 1
	}

	maxVisible := 0
	var maxPosition [2]int
	for i := 0; i < len(asteroids)-1; i++ {
		for j := i + 1; j < len(asteroids); j++ {
			a, b := asteroids[i], asteroids[j]
			dr, dc := b[0]-a[0], b[1]-a[1]
			gcd := helpers.AbsInt(helpers.GCD(dr, dc))
			if gcd == 1 {
				visible[a]++
				visible[b]++
				if visible[a] > maxVisible {
					maxVisible = visible[a]
					maxPosition = a
				}
				if visible[b] > maxVisible {
					maxVisible = visible[b]
					maxPosition = b
				}
				continue
			}
			// if gcd != 1, then check all points between a and b for asteroid
			dr, dc = dr/gcd, dc/gcd
			blocked := false
			for r, c := a[0]+dr, a[1]+dc; r >= 0 && r < gh && c >= 0 && c < gw; r, c = r+dr, c+dc {
				next := [2]int{r, c}
				if b == next {
					break
				}
				if _, exists := visible[next]; exists {
					blocked = true
					break
				}
			}
			if !blocked {
				visible[a]++
				visible[b]++
				if visible[a] > maxVisible {
					maxVisible = visible[a]
					maxPosition = a
				}
				if visible[b] > maxVisible {
					maxVisible = visible[b]
					maxPosition = b
				}
			}
		}
	}
	fmt.Println(maxPosition, maxVisible)
}

func (d Day10) Part2(input string) {
}
