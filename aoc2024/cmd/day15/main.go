package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"regexp"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, _ := os.Open(infile)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	rexpMovement := regexp.MustCompile(`[<>v^]`)
	rexpRobot := regexp.MustCompile(`@`)
	scanner := bufio.NewScanner(file)
	movements := ""
	start := [2]int{}
	grid := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if rexpMovement.MatchString(line) {
			movements += line
			continue
		}
		grid = append(grid, []rune(line))
		if rexpRobot.MatchString(line) {
			start[0] = rexpRobot.FindStringIndex(line)[1] - 1
			start[1] = len(grid) - 1
		}
	}

	gridCopy := make([][]rune, len(grid))
	for i := range grid {
		gridCopy[i] = make([]rune, len(grid[i]))
		copy(gridCopy[i], grid[i])
	}
	part1(gridCopy, movements, start)
	part2(grid, movements, start)
}

func part2(grid [][]rune, movements string, start [2]int) {
	grid = expandGrid(grid)
	start[0] = start[0] * 2
	printGrid(grid)
	x := start[0]
	y := start[1]
	for _, dir := range movements {
		if dir == '<' || dir == '>' {
			if move(x, y, dir, grid) {
				x += directions[dir][0]
				y += directions[dir][1]
			}
		} else {
			if isBlocked(x, y, dir, grid) {
				continue
			}
			moveUpOrDown(x, y, dir, grid)
			x += directions[dir][0]
			y += directions[dir][1]
		}
	}
	printGrid(grid)

	sum := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '[' {
				if grid[row][col+1] != ']' {
					printGrid(grid)
					panic(fmt.Sprint(col, row))
				}
				sum += 100*row + col
			}
		}
	}
	slog.Info("Part 2:", "sum", sum)
}

func moveUpOrDown(x int, y int, dir rune, grid [][]rune) {
	nextY := y + directions[dir][1]
	// move the robot
	if grid[y][x] == '@' {
		switch grid[nextY][x] {
		// unobstructed
		case '.':
			grid[nextY][x] = '@'
			grid[y][x] = '.'
			return
		// obstructed by [-side of box
		case '[':
			moveUpOrDown(x, nextY, dir, grid)
			grid[nextY][x] = '@'
			grid[y][x] = '.'
			return
		// obstructed by ]-side of box
		case ']':
			// always use coordinate of [-side of box
			moveUpOrDown(x-1, nextY, dir, grid)
			grid[nextY][x] = '@'
			grid[y][x] = '.'
			return
		default:
			panic(fmt.Sprint(x, y, dir, grid))
		}
	}
	// move a box

	// box is unobstructed
	// e.g., direction ^
	// ....
	//  []
	if grid[nextY][x] == '.' && grid[nextY][x+1] == '.' {
		grid[nextY][x] = '['
		grid[nextY][x+1] = ']'
		grid[y][x] = '.'
		grid[y][x+1] = '.'
		return
	}

	// box is obstructed by two boxes
	// e.g., direction ^
	// [][]
	// .[].
	if grid[nextY][x] == ']' && grid[nextY][x+1] == '[' {
		moveUpOrDown(x-1, nextY, dir, grid)
		moveUpOrDown(x+1, nextY, dir, grid)
		grid[nextY][x] = '['
		grid[nextY][x+1] = ']'
		grid[y][x] = '.'
		grid[y][x+1] = '.'
		return
	}

	// box obstructed by right half of one box
	// e.g., direction ^
	// []..
	// .[].
	if grid[nextY][x] == ']' && grid[nextY][x+1] == '.' {
		moveUpOrDown(x-1, nextY, dir, grid)
		grid[nextY][x] = '['
		grid[nextY][x+1] = ']'
		grid[y][x] = '.'
		grid[y][x+1] = '.'
		return
	}

	// box obstructed by left half of one box
	// e.g., direction ^
	// ..[]
	// .[].
	if grid[nextY][x] == '.' && grid[nextY][x+1] == '[' {
		moveUpOrDown(x+1, nextY, dir, grid)
		grid[nextY][x] = '['
		grid[nextY][x+1] = ']'
		grid[y][x] = '.'
		grid[y][x+1] = '.'
		return
	}

	// box is obstructed by box directly aligned
	// e.g., direction ^
	// .[].
	// .[].
	moveUpOrDown(x, nextY, dir, grid)
	grid[nextY][x] = '['
	grid[nextY][x+1] = ']'
	grid[y][x] = '.'
	grid[y][x+1] = '.'
}

func isBlocked(x int, y int, dir rune, grid [][]rune) bool {
	nextX := x + directions[dir][0]
	nextY := y + directions[dir][1]
	if grid[nextY][nextX] == '#' {
		return true
	}
	if grid[nextY][nextX] == '.' {
		return false
	}
	if grid[nextY][nextX] == '[' {
		return isBlocked(nextX, nextY, dir, grid) || isBlocked(nextX+1, nextY, dir, grid)
	}
	return isBlocked(nextX, nextY, dir, grid) || isBlocked(nextX-1, nextY, dir, grid)
}

func expandGrid(grid [][]rune) [][]rune {
	newGrid := make([][]rune, 0, len(grid))
	for _, row := range grid {
		newRow := make([]rune, 0, len(row)*2)
		for _, c := range row {
			switch c {
			case '#':
				newRow = append(newRow, '#', '#')
			case 'O':
				newRow = append(newRow, '[', ']')
			case '@':
				newRow = append(newRow, '@', '.')
			default:
				newRow = append(newRow, '.', '.')
			}
		}
		newGrid = append(newGrid, newRow)
	}
	return newGrid
}

func part1(grid [][]rune, movements string, start [2]int) {
	printGrid(grid)
	x := start[0]
	y := start[1]
	for _, dir := range movements {
		if move(x, y, dir, grid) {
			x += directions[dir][0]
			y += directions[dir][1]
		}
	}
	printGrid(grid)

	sum := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 'O' {
				sum += 100*row + col
			}
		}
	}
	slog.Info("Part 1:", "sum", sum)
}
func gridToString(grid [][]rune) string {
	ret := ""
	for row := range grid {
		ret += string(grid[row]) + "\n"
	}
	return ret
}

func printGrid(grid [][]rune) {
	fmt.Println(gridToString(grid))
}

var directions = map[rune][2]int{
	'^': {0, -1},
	'v': {0, 1},
	'<': {-1, 0},
	'>': {1, 0},
}

func move(x int, y int, dir rune, grid [][]rune) bool {
	nextX := x + directions[dir][0]
	nextY := y + directions[dir][1]
	if grid[nextY][nextX] == '#' {
		return false
	}
	if grid[nextY][nextX] == '.' {
		grid[nextY][nextX] = grid[y][x]
		grid[y][x] = '.'
		return true
	}
	// try to move obstacle (recursively)
	if move(nextX, nextY, dir, grid) {
		grid[nextY][nextX] = grid[y][x]
		grid[y][x] = '.'
		return true
	}
	// obstacle cannot be moved
	return false
}
