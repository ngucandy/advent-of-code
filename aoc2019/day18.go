package aoc2019

import (
	"strings"
)

func init() {
	Days["18"] = &Day18{
		eg1: `#########
#b.A.@.a#
#########`,
		eg2: `########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################`,
		eg3: `########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################`,
		eg4: `#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################`,
		eg5: `########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################`,
	}
}

type Day18 struct {
	eg1, eg2, eg3, eg4, eg5 string
}

func (d Day18) Part1(input string) any {
	var grid [][]rune
	var sr, sc, numKeys int
	for r, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
		for c, ch := range line {
			if ch == '@' {
				sr, sc = r, c
				continue
			}
			if ch >= 'a' && ch <= 'z' {
				numKeys++
				continue
			}
		}
	}

	index := func(ch rune) int {
		return int(ch - 'a')
	}

	lcOffset := 'a' - 'A'
	type position struct {
		r, c int
		keys [26]bool
	}
	type state struct {
		position
		steps     int
		collected int
	}
	visited := make(map[position]int)
	q := []state{
		{position: position{sr, sc, [26]bool{}}, steps: 0, collected: 0},
	}
	var ans int
	for len(q) > 0 {
		p, steps, collected := q[0].position, q[0].steps, q[0].collected
		q = q[1:]
		//fmt.Println(p.r, p.c, steps)

		if vsteps, exists := visited[p]; exists && vsteps <= steps {
			//fmt.Println("already visited")
			continue
		}
		visited[p] = steps

		ch := grid[p.r][p.c]
		if ch >= 'A' && ch <= 'Z' && !p.keys[index(ch+lcOffset)] {
			continue
		}
		if ch >= 'a' && ch <= 'z' && !p.keys[index(ch)] {
			p.keys[index(ch)] = true
			collected++
			if collected == numKeys {
				ans = steps
				break
			}
			visited[p] = steps
		}

		for _, dir := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			np := position{p.r + dir[0], p.c + dir[1], p.keys}
			if np.r < 0 || np.r >= len(grid) || np.c < 0 || np.c >= len(grid[np.r]) || grid[np.r][np.c] == '#' {
				continue
			}
			q = append(q, state{np, steps + 1, collected})
		}
	}

	return ans
}

func (d Day18) Part2(input string) any {
	return "no answer yet"
}
