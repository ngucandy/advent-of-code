package aoc2019

import (
	"strings"
)

func init() {
	Days["20"] = &Day20{
		eg1: `         A           
         A           
  #######.#########  
  #######.........#  
  #######.#######.#  
  #######.#######.#  
  #######.#######.#  
  #####  B    ###.#  
BC...##  C    ###.#  
  ##.##       ###.#  
  ##...DE  F  ###.#  
  #####    G  ###.#  
  #########.#####.#  
DE..#######...###.#  
  #.#########.###.#  
FG..#########.....#  
  ###########.#####  
             Z       
             Z       `,
		eg2: `                   A               
                   A               
  #################.#############  
  #.#...#...................#.#.#  
  #.#.#.###.###.###.#########.#.#  
  #.#.#.......#...#.....#.#.#...#  
  #.#########.###.#####.#.#.###.#  
  #.............#.#.....#.......#  
  ###.###########.###.#####.#.#.#  
  #.....#        A   C    #.#.#.#  
  #######        S   P    #####.#  
  #.#...#                 #......VT
  #.#.#.#                 #.#####  
  #...#.#               YN....#.#  
  #.###.#                 #####.#  
DI....#.#                 #.....#  
  #####.#                 #.###.#  
ZZ......#               QG....#..AS
  ###.###                 #######  
JO..#.#.#                 #.....#  
  #.#.#.#                 ###.#.#  
  #...#..DI             BU....#..LF
  #####.#                 #.#####  
YN......#               VT..#....QG
  #.###.#                 #.###.#  
  #.#...#                 #.....#  
  ###.###    J L     J    #.#.###  
  #.....#    O F     P    #.#...#  
  #.###.#####.#.#####.#####.###.#  
  #...#.#.#...#.....#.....#.#...#  
  #.#####.###.###.#.#.#########.#  
  #...#.#.....#...#.#.#.#.....#.#  
  #.###.#####.###.###.#.#.#######  
  #.#.........#...#.............#  
  #########.###.###.#############  
           B   J   C               
           U   P   P               `,
	}
}

type Day20 struct {
	eg1, eg2 string
}

func (d Day20) Part1(input string) any {
	grid, sr, sc, portals := d.parseInput(input)

	var ans int
	q := [][3]int{{sr, sc, 0}}
	seen := make(map[[2]int]bool)
	for len(q) > 0 {
		r, c, steps := q[0][0], q[0][1], q[0][2]
		q = q[1:]

		if grid[r][c] == '$' {
			ans = steps
			break
		}

		if seen[[2]int{r, c}] {
			continue
		}
		seen[[2]int{r, c}] = true

		if grid[r][c] == '%' {
			// add other end of portal
			p := portals[[2]int{r, c}]
			q = append(q, [3]int{p[0], p[1], steps + 1})
		}

		for _, dir := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			nr, nc := r+dir[0], c+dir[1]
			nch := grid[nr][nc]
			if nch == '#' || (nch >= 'A' && nch <= 'Z') {
				continue
			}
			q = append(q, [3]int{nr, nc, steps + 1})
		}
	}
	return ans
}

func (d Day20) parseInput(input string) ([][]rune, int, int, map[[2]int][2]int) {
	var grid [][]rune
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}
	// find portals, start and end
	var sr, sc int
	portalMap := make(map[string][][2]int)
	for r, row := range grid {
		for c, ch := range row {
			if ch != '.' {
				continue
			}
			for _, dir := range [][]int{{-2, 0, -1, 0}, {1, 0, 2, 0}, {0, -2, 0, -1}, {0, 1, 0, 2}} {
				ch1, ch2 := grid[r+dir[0]][c+dir[1]], grid[r+dir[2]][c+dir[3]]
				key := string([]rune{ch1, ch2})
				if key == "AA" {
					sr, sc = r, c
					grid[r][c] = '^'
					continue
				}
				if key == "ZZ" {
					grid[r][c] = '$'
					continue
				}
				if ch1 >= 'A' && ch1 <= 'Z' && ch2 >= 'A' && ch2 <= 'Z' {
					portalMap[key] = append(portalMap[key], [2]int{r, c})
					grid[r][c] = '%'
				}
			}
		}
	}
	portals := make(map[[2]int][2]int)
	for _, p := range portalMap {
		portals[p[0]] = p[1]
		portals[p[1]] = p[0]
	}
	return grid, sr, sc, portals
}

func (d Day20) Part2(input string) any {
	return "no answer yet"
}
