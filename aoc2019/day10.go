package aoc2019

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/ngucandy/advent-of-code/internal/helpers"
)

func init() {
	Days["10"] = Day10{
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
		`.#....#####...#..
##...##.#####..##
##...#...#.#####.
..#.....#...###..
..#.#.....#....##`,
	}
}

type Day10 struct {
	eg1, eg2, eg3, eg4, eg5, eg6 string
}

func (d Day10) Part1(input string) any {
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
				}
				if visible[b] > maxVisible {
					maxVisible = visible[b]
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
				}
				if visible[b] > maxVisible {
					maxVisible = visible[b]
				}
			}
		}
	}
	return maxVisible
}

func (d Day10) Part2(input string) any {
	//input = d.eg5
	//s := [2]int{13, 11}
	s := [2]int{29, 26}
	asteroids := make(map[[2]int]struct{})
	gh, gw := 0, 0
	for r, line := range strings.Split(input, "\n") {
		gh = r + 1
		for c, ch := range line {
			gw = c + 1
			// skip asteroid where station is located
			if r == s[0] && c == s[1] {
				continue
			}
			if ch == '#' {
				pos := [2]int{r, c}
				asteroids[pos] = struct{}{}
			}
		}
	}

	var visible [][4]int
	for a := range asteroids {
		dr, dc := a[0]-s[0], a[1]-s[1]
		gcd := helpers.AbsInt(helpers.GCD(dr, dc))
		if gcd == 1 {
			visible = append(visible, [4]int{a[0], a[1], dr, dc})
			continue
		}
		// if gcd != 1, then check all points between s and a for asteroid
		dr, dc = dr/gcd, dc/gcd
		blocked := false
		for r, c := s[0]+dr, s[1]+dc; r >= 0 && r < gh && c >= 0 && c < gw; r, c = r+dr, c+dc {
			next := [2]int{r, c}
			if a == next {
				break
			}
			if _, exists := asteroids[next]; exists {
				blocked = true
				break
			}
		}
		if !blocked {
			visible = append(visible, [4]int{a[0], a[1], dr, dc})
		}
	}

	slices.SortFunc(visible, d.sort)
	a200 := visible[199]
	return a200[1]*100 + a200[0]
}

func (d Day10) sort(a, b [4]int) int {
	ar, br := a[0], b[0]
	ac, bc := a[1], b[1]
	adr, adc := a[2], a[3]
	bdr, bdc := b[2], b[3]

	// a or b points up or down
	if adc == 0 || bdc == 0 {
		// a up, b up
		if adc == 0 && adr < 0 && bdc == 0 && bdr < 0 {
			return cmp.Compare(-ar, -br)
		}

		// a down b down
		if adc == 0 && adr > 0 && bdc == 0 && bdr > 0 {
			return cmp.Compare(ar, br)
		}

		// a up, b not up
		if adc == 0 && adr < 0 {
			return -1
		}

		// a not up, b up
		if bdc == 0 && bdr < 0 {
			return 1
		}

		// a down; b left
		if adc == 0 && adr > 0 && bdc < 0 {
			return cmp.Compare(-ac, bc)
		}

		// a down; b right
		if adc == 0 && adr > 0 && bdc > 0 {
			return cmp.Compare(ac, -bc)
		}

		// b down; a left
		if bdc == 0 && bdr > 0 && adc < 0 {
			return cmp.Compare(ac, -bc)
		}

		// b down; a right
		if bdc == 0 && bdr > 0 && adc > 0 {
			return cmp.Compare(-ac, bc)
		}

		panic(fmt.Sprintf("unhandled comparison where a or b is vertical: a=%v; b=%v", a, b))
	}

	aslope := float64(adr) / float64(adc)
	bslope := float64(bdr) / float64(bdc)

	// a and b both point left or right
	if (adc > 0 && bdc > 0) || (adc < 0 && bdc < 0) {
		return cmp.Compare(aslope, bslope)
	}

	// a points left; b points right
	if adc < 0 && bdc > 0 {
		return cmp.Compare(ac, -bc)
	}

	// a points right; b points left
	return cmp.Compare(-ac, bc)
}
