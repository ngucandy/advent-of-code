package aoc2024

import (
	"fmt"
	"strings"
)

func init() {
	Days["8"] = Day8{
		`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`,
	}
}

type Day8 struct {
	example string
}

func (d Day8) Part1(input string) {
	var w, h int
	frequencies := make(map[rune][][2]int)
	for r, line := range strings.Split(input, "\n") {
		h = max(h, r+1)
		for c, ch := range line {
			w = max(w, c+1)
			if ch == '.' {
				continue
			}
			frequencies[ch] = append(frequencies[ch], [2]int{r, c})
		}
	}

	antinodes := make(map[[2]int]struct{})
	for _, antennas := range frequencies {
		for i, a1 := range antennas[:len(antennas)-1] {
			for _, a2 := range antennas[i+1:] {
				dr, dc := a2[0]-a1[0], a2[1]-a1[1]
				nr, nc := a2[0]+dr, a2[1]+dc
				if nr >= 0 && nr < h && nc >= 0 && nc < w {
					antinodes[[2]int{nr, nc}] = struct{}{}
				}
				nr, nc = a1[0]-dr, a1[1]-dc
				if nr >= 0 && nr < h && nc >= 0 && nc < w {
					antinodes[[2]int{nr, nc}] = struct{}{}
				}
			}
		}
	}
	fmt.Println("part1", len(antinodes))
}

func (d Day8) Part2(input string) {
	var w, h int
	frequencies := make(map[rune][][2]int)
	for r, line := range strings.Split(input, "\n") {
		h = max(h, r+1)
		for c, ch := range line {
			w = max(w, c+1)
			if ch == '.' {
				continue
			}
			frequencies[ch] = append(frequencies[ch], [2]int{r, c})
		}
	}

	antinodes := make(map[[2]int]struct{})
	for _, antennas := range frequencies {
		for i, a1 := range antennas[:len(antennas)-1] {
			for _, a2 := range antennas[i+1:] {
				dr, dc := a2[0]-a1[0], a2[1]-a1[1]
				for _, dir := range [][2]int{{dr, dc}, {-dr, -dc}} {
					for _, antenna := range [][2]int{a1, a2} {
						for nr, nc := antenna[0]+dir[0], antenna[1]+dir[1]; nr >= 0 && nr < h && nc >= 0 && nc < w; nr, nc = nr+dir[0], nc+dir[1] {
							antinodes[[2]int{nr, nc}] = struct{}{}
						}
					}
				}
			}
		}
	}
	fmt.Println("part2", len(antinodes))
}
