package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strconv"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, _ := os.Open(infile)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	// e.g., p=92,72 v=-49,-72
	rexpNums := regexp.MustCompile(`-?\d+`)
	robots := [][2][2]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := rexpNums.FindAllString(line, -1)
		px, _ := strconv.Atoi(nums[0])
		py, _ := strconv.Atoi(nums[1])
		vx, _ := strconv.Atoi(nums[2])
		vy, _ := strconv.Atoi(nums[3])
		robot := [2][2]int{{px, py}, {vx, vy}}
		robots = append(robots, robot)
	}

	part1(robots)
	part2(robots)
}

func part1(robots [][2][2]int) {
	rows := 103
	cols := 101
	seconds := 100
	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0
	for _, robot := range robots {
		startx := robot[0][0]
		starty := robot[0][1]
		velx := robot[1][0]
		vely := robot[1][1]
		negxvel := false
		negyvel := false

		if velx < 0 {
			startx = cols - startx - 1
			velx = -velx
			negxvel = true
		}
		if vely < 0 {
			starty = rows - starty - 1
			vely = -vely
			negyvel = true
		}
		endx := (velx*seconds + startx) % cols
		endy := (vely*seconds + starty) % rows
		if negxvel {
			endx = cols - endx - 1
		}
		if negyvel {
			endy = rows - endy - 1
		}

		if endx == cols/2 || endy == rows/2 {
			continue
		}
		if endx < cols/2 && endy < rows/2 {
			q1++
			continue
		}
		if endx < cols/2 && endy > rows/2 {
			q2++
			continue
		}
		if endx > cols/2 && endy < rows/2 {
			q3++
			continue
		}
		if endx > cols/2 && endy > rows/2 {
			q4++
			continue
		}
	}
	slog.Info("Part 1:", "safety factor", q1*q2*q3*q4, "q1", q1, "q2", q2, "q3", q3, "q4", q4)
}

func part2(robots [][2][2]int) {
	rows := 103
	cols := 101
	seconds := 0
	for {
		seconds++
		graph := make([][]string, rows)
		for row := 0; row < rows; row++ {
			graph[row] = make([]string, cols)
			for col := 0; col < cols; col++ {
				graph[row][col] = "."
			}
		}
		cluster := false
		for _, robot := range robots {
			startx := robot[0][0]
			starty := robot[0][1]
			velx := robot[1][0]
			vely := robot[1][1]
			negxvel := false
			negyvel := false

			if velx < 0 {
				startx = cols - startx - 1
				velx = -velx
				negxvel = true
			}
			if vely < 0 {
				starty = rows - starty - 1
				vely = -vely
				negyvel = true
			}
			endx := (velx*seconds + startx) % cols
			endy := (vely*seconds + starty) % rows
			if negxvel {
				endx = cols - endx - 1
			}
			if negyvel {
				endy = rows - endy - 1
			}
			graph[endy][endx] = "#"

			// look for a robot surrounded by robots on all sides
			localcluster := true
			for _, dir := range [][2]int{{-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}} {
				neighborx := endx + dir[0]
				neighbory := endy + dir[1]
				if neighborx < 0 || neighborx >= cols || neighbory < 0 || neighbory >= rows {
					localcluster = false
					break
				}
				if graph[neighbory][neighborx] != "#" {
					localcluster = false
					break
				}
			}
			cluster = cluster || localcluster
		}

		if cluster {
			for _, row := range graph {
				fmt.Println(row)
			}
			break
		}
	}
	slog.Info("Part 2:", "seconds", seconds)
}
