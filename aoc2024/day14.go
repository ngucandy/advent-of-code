package aoc2024

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/ngucandy/advent-of-code/internal/helpers"
)

func init() {
	Days["14"] = Day14{}
}

type Day14 struct {
	example string
}

func (d Day14) Part1(input string) any {
	var robots [][2][2]int
	for _, line := range strings.Split(input, "\n") {
		// e.g., p=92,72 v=-49,-72
		var px, py, vx, vy int
		_, _ = fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		robots = append(robots, [2][2]int{{px, py}, {vx, vy}})
	}

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

		endx := (((velx*seconds + startx) % cols) + cols) % cols
		endy := (((vely*seconds + starty) % rows) + rows) % rows

		if endx == cols/2 || endy == rows/2 {
			continue
		}
		if endx < cols/2 && endy < rows/2 {
			q1++
			continue
		}
		if endx < cols/2 {
			q2++
			continue
		}
		if endy < rows/2 {
			q3++
			continue
		}
		q4++
	}
	return q1 * q2 * q3 * q4
}
func (d Day14) Part2(input string) any {
	var robots [][2][2]int
	for _, line := range strings.Split(input, "\n") {
		// e.g., p=92,72 v=-49,-72
		var px, py, vx, vy int
		_, _ = fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		robots = append(robots, [2][2]int{{px, py}, {vx, vy}})
	}

	rows := 103
	cols := 101
	seconds := 0
	for {
		seconds++
		graph := make([][]rune, rows)
		for row := 0; row < rows; row++ {
			graph[row] = make([]rune, cols)
			for col := 0; col < cols; col++ {
				graph[row][col] = ' '
			}
		}
		cluster := false
		for _, robot := range robots {
			startx := robot[0][0]
			starty := robot[0][1]
			velx := robot[1][0]
			vely := robot[1][1]

			endx := (((velx*seconds + startx) % cols) + cols) % cols
			endy := (((vely*seconds + starty) % rows) + rows) % rows
			graph[endy][endx] = '\u2588'

			// look for a robot surrounded by robots on all sides
			localcluster := true
			for _, dir := range [][2]int{{-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}} {
				neighborx := endx + dir[0]
				neighbory := endy + dir[1]
				if neighborx < 0 || neighborx >= cols || neighbory < 0 || neighbory >= rows {
					localcluster = false
					break
				}
				if graph[neighbory][neighborx] != '\u2588' {
					localcluster = false
					break
				}
			}
			cluster = cluster || localcluster
		}
		if cluster {
			helpers.PrintGrid(graph)
			d.makeImage(rows, cols, graph, seconds)
			break
		}
	}
	d.animate(rows, cols, robots, seconds)
	return seconds
}

func (d Day14) makeImage(rows, cols int, graph [][]rune, frame int) {
	scale := 6
	img := image.NewRGBA(image.Rect(0, 0, rows*scale, cols*scale))
	c := color.RGBA{0, 255, 0, 0xff}
	for row := 0; row < len(graph); row++ {
		for col := 0; col < len(graph[row]); col++ {
			if graph[row][col] == '\u2588' {
				x := col * scale
				y := row * scale
				r := image.Rect(x, y, x+scale, y+scale)
				draw.Draw(img, r, &image.Uniform{c}, image.Point{0, 0}, draw.Src)
			}
		}
	}
	cwd, _ := os.Getwd()
	f, err := os.Create(filepath.Join(cwd, "aoc2024", "day14", fmt.Sprintf("%04d.png", frame)))
	if err != nil {
		panic(err)
	}
	_ = png.Encode(f, img)
}

func (d Day14) animate(rows, cols int, robots [][2][2]int, simStop int) {
	scale := 10
	width := cols * scale
	height := rows * scale
	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	rl.InitWindow(int32(width), int32(height+50), "Advent of Code 2024 - Day 14")
	defer rl.CloseWindow()
	fpsTarget := int32(60)
	rl.SetTargetFPS(fpsTarget)

	simSeconds := simStop - 10 // simulation seconds
	frames := 0
	fpsSim := 60 // frames per simulation seconds
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		if simSeconds < simStop || (simSeconds == simStop && frames%fpsSim == 0) {
			rl.ClearBackground(rl.RayWhite)
			rl.DrawText(fmt.Sprintf("simulation time: %ds (stopping in %d)", simSeconds, simStop-simSeconds), 5, int32(height+20), 20, rl.Red)
			for _, robot := range robots {
				x, y := float32(robot[0][0]*scale), float32(robot[0][1]*scale)   // initial position
				vx, vy := float32(robot[1][0]*scale), float32(robot[1][1]*scale) // velocity
				dx, dy := vx/float32(fpsSim), vy/float32(fpsSim)
				nx := int32(((int(x+vx*float32(simSeconds)+dx*float32(frames%fpsSim)) % width) + width) % width)
				ny := int32(((int(y+vy*float32(simSeconds)+dy*float32(frames%fpsSim)) % height) + height) % height)

				//rl.DrawRectangle(nx, ny, int32(scale), int32(scale), rl.Green)
				//rl.DrawCircle(nx+(int32(scale)/2), ny+(int32(scale)/2), float32(scale/2)+2, rl.Green)

				v1 := rl.NewVector2(float32(nx+int32(scale/2)), float32(ny))
				v2 := rl.NewVector2(float32(nx), float32(ny+int32(scale)))
				v3 := rl.NewVector2(float32(nx+int32(scale)), float32(ny+int32(scale)))
				rl.DrawTriangle(v1, v2, v3, rl.Green)
			}
		}
		rl.EndDrawing()

		frames++
		if frames%fpsSim == 0 {
			simSeconds++
		}
		// reduce simulation speed for the final 3 seconds
		if simSeconds == simStop-3 && frames%fpsSim == 0 {
			fpsTarget /= 4
			rl.SetTargetFPS(fpsTarget)
		}
	}
}
