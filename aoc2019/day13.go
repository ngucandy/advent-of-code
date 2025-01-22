package aoc2019

import (
	"fmt"
	"github.com/ngucandy/advent-of-code/internal/helpers"
	"slices"
)

func init() {
	Days["13"] = &Day13{}
}

type Day13 struct {
	eg1, eg2 string
}

func (d Day13) Part1(input string) any {
	comp := NewIntcodeComputer(ParseIntcodeProgram(input), nil)
	for comp.Step() {
	}
	var tiles [5][][2]int
	var mr, mc int
	for chunk := range slices.Chunk(comp.output, 3) {
		x, y, t := chunk[0], chunk[1], chunk[2]
		mr = max(mr, y)
		mc = max(mc, x)
		if t == 0 {
			continue
		}
		tiles[t] = append(tiles[t], [2]int{y, x})
	}

	var grid [][]rune
	for range mr + 1 {
		grid = append(grid, slices.Repeat([]rune{' '}, mc+1))
	}
	for t, positions := range tiles {
		for _, pos := range positions {
			r, c := pos[0], pos[1]
			switch t {
			case 1:
				grid[r][c] = '#'
			case 2:
				grid[r][c] = '\u2588'
			case 3:
				grid[r][c] = '='
			case 4:
				grid[r][c] = '*'
			}
		}
	}
	helpers.PrintGrid(grid)
	return len(tiles[2])
}

func (d Day13) Part2(input string) any {
	memory := ParseIntcodeProgram(input)
	memory[0] = 2 // free play
	// extend the paddle wall-to-wall
	for i := range memory[:len(memory)-2] {
		if memory[i] != 0 || memory[i+1] != 3 || memory[i+2] != 0 {
			continue
		}
		// found paddle area of memory
		// extend paddle left
		for j := i; memory[j] != 1; j-- {
			memory[j] = 3
		}
		// extend paddle right
		for j := i + 2; memory[j] != 1; j++ {
			memory[j] = 3
		}
		break
	}
	var score int
	blocks := make(map[[2]int]struct{})
	comp := NewIntcodeComputer(memory, slices.Repeat([]int{0}, 1_000_000))
	initialized := false
	for comp.Step() {
		if len(comp.output) < 3 {
			continue
		}
		output := comp.output[:3]
		comp.output = comp.output[3:]
		if output[0] == -1 { // update score
			// first time we see a score is after all tiles have been setup
			initialized = true
			score = output[2]
			continue
		}
		switch output[2] {
		case 0: // empty tile
			// ignore if game is still initializing
			if !initialized {
				continue
			}
			// empty tile after game is initialized could be a block breaking
			// or ball moving
			key := [2]int{output[0], output[1]}
			if _, exists := blocks[key]; !exists {
				// ignore if it's not a block breaking
				continue
			}
			delete(blocks, key)
		case 2: // block tile
			// should only see these when the game is initializing
			if initialized {
				panic(fmt.Sprintf("unexpected block tile after initialization: %v", output))
			}
			key := [2]int{output[0], output[1]}
			blocks[key] = struct{}{}
		}
	}
	return score
}
