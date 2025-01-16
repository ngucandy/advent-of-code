package aoc2024

import (
	"fmt"
	"strings"
)

func init() {
	DayMap["15"] = Day15{
		example: `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`,
		directions: map[rune][2]int{
			'^': {-1, 0},
			'v': {1, 0},
			'<': {0, -1},
			'>': {0, 1},
		},
	}
}

type Day15 struct {
	example    string
	directions map[rune][2]int
}

func (d Day15) Part1(input string) {
	sections := strings.Split(input, "\n\n")

	var grid [][]rune
	var sr, sc int
	for r, line := range strings.Split(sections[0], "\n") {
		grid = append(grid, []rune(line))
		if c := strings.Index(line, "@"); c != -1 {
			sr, sc = r, c
		}
	}
	movements := strings.Join(strings.Split(sections[1], "\n"), "")

	rr, rc := sr, sc
	for _, m := range movements {
		if d.move(rr, rc, m, grid) {
			rr += d.directions[m][0]
			rc += d.directions[m][1]
		}
	}

	sum := 0
	for r, row := range grid {
		for c, ch := range row {
			if ch == 'O' {
				sum += 100*r + c
			}
		}
	}
	fmt.Println("part1", sum)
}

func (d Day15) move(r, c int, m rune, grid [][]rune) bool {
	nr, nc := r+d.directions[m][0], c+d.directions[m][1]
	if grid[nr][nc] == '#' {
		// wall
		return false
	}
	if grid[nr][nc] == '.' {
		// space is open
		grid[nr][nc], grid[r][c] = grid[r][c], '.'
		return true
	}
	// obstacle in the way so try to (recursively)
	if d.move(nr, nc, m, grid) {
		grid[nr][nc], grid[r][c] = grid[r][c], '.'
		return true
	}
	// obstacle cannot be moved
	return false
}

func (d Day15) Part2(input string) {
	sections := strings.Split(input, "\n\n")

	var grid [][]rune
	var sr, sc int
	for r, line := range strings.Split(sections[0], "\n") {
		line = strings.Replace(line, "#", "##", -1)
		line = strings.Replace(line, "O", "[]", -1)
		line = strings.Replace(line, ".", "..", -1)
		line = strings.Replace(line, "@", "@.", -1)
		grid = append(grid, []rune(line))
		if c := strings.Index(line, "@"); c != -1 {
			sr, sc = r, c
		}
	}
	movements := strings.Join(strings.Split(sections[1], "\n"), "")

	rr, rc := sr, sc
	for _, m := range movements {
		if m == '<' || m == '>' {
			// moving left or right is the same as part1
			if d.move(rr, rc, m, grid) {
				rr += d.directions[m][0]
				rc += d.directions[m][1]
			}
		} else { // moving up or down
			if d.blocked(rr, rc, m, grid) {
				continue
			}
			d.moveUpOrDown(rr, rc, m, grid)
			rr += d.directions[m][0]
			rc += d.directions[m][1]
		}
	}

	sum := 0
	for r, row := range grid {
		for c, ch := range row {
			if ch == '[' {
				sum += 100*r + c
			}
		}
	}
	fmt.Println("part2", sum)
}

func (d Day15) moveUpOrDown(r, c int, m rune, grid [][]rune) {
	// extra?
	if d.blocked(r, c, m, grid) {
		panic(fmt.Sprintf("cannot move %c because path is blocked", m))
	}
	nr := r + d.directions[m][0]

	// moving the robot only requires changing a single grid position
	if grid[r][c] == '@' {
		switch grid[nr][c] {
		// unobstructed
		case '.':
			grid[nr][c], grid[r][c] = '@', '.'
			return
		// obstructed by [ side of box
		case '[':
			d.moveUpOrDown(nr, c, m, grid)
			grid[nr][c], grid[r][c] = '@', '.'
			return
		// obstructed by ] side of box
		case ']':
			// always use coordinate of [ side of box
			d.moveUpOrDown(nr, c-1, m, grid)
			grid[nr][c], grid[r][c] = '@', '.'
			return
		default:
			panic(fmt.Sprintf("foreign object in grid[%d][%d]: %c", nr, c, grid[nr][c]))
		}
	}

	// moving a box requires changing two grid positions
	if grid[nr][c] == ']' && grid[nr][c+1] == '[' { // box is obstructed by two boxes
		// e.g., direction ^
		// [][]
		// .[].
		d.moveUpOrDown(nr, c-1, m, grid)
		d.moveUpOrDown(nr, c+1, m, grid)
	} else if grid[nr][c] == ']' && grid[nr][c+1] == '.' { // box obstructed by right half of single box
		// e.g., direction ^
		// []..
		// .[].
		d.moveUpOrDown(nr, c-1, m, grid)
	} else if grid[nr][c] == '.' && grid[nr][c+1] == '[' { // box obstructed by left half of single box
		// e.g., direction ^
		// ..[]
		// .[].
		d.moveUpOrDown(nr, c+1, m, grid)
	} else if grid[nr][c] == '[' && grid[nr][c+1] == ']' { // box is obstructed by box directly aligned
		// e.g., direction ^
		// .[].
		// .[].
		d.moveUpOrDown(nr, c, m, grid)
	} // else box is unobstructed

	grid[nr][c], grid[nr][c+1] = '[', ']'
	grid[r][c], grid[r][c+1] = '.', '.'
}

func (d Day15) blocked(r, c int, m rune, grid [][]rune) bool {
	nr, nc := r+d.directions[m][0], c+d.directions[m][1]
	if grid[nr][nc] == '#' {
		return true
	}
	if grid[nr][nc] == '.' {
		return false
	}
	if grid[nr][nc] == '[' {
		return d.blocked(nr, nc, m, grid) || d.blocked(nr, nc+1, m, grid)
	}
	// grid[nr][nc] == ']'
	return d.blocked(nr, nc, m, grid) || d.blocked(nr, nc-1, m, grid)
}
