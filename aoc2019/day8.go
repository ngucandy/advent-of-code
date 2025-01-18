package aoc2019

import (
	"fmt"
	"math"
	"strings"
)

func init() {
	Days["8"] = Day8{
		`123456789012`,
		`0222112222120000`,
	}
}

type Day8 struct {
	example1, example2 string
}

func (d Day8) Part1(input string) {
	//input = d.example
	//w, h := 3, 2
	w, h := 25, 6
	layers := d.parseLayers(input, w, h)

	var minLayer [][]int
	minZeros := math.MaxInt
	for _, layer := range layers {
		zeros := 0
		for _, row := range layer {
			for _, n := range row {
				if n == 0 {
					zeros++
				}
			}
		}
		if zeros < minZeros {
			minZeros = zeros
			minLayer = layer
		}
	}

	ones := 0
	twos := 0
	for _, row := range minLayer {
		for _, n := range row {
			switch n {
			case 1:
				ones++
			case 2:
				twos++
			}
		}
	}
	fmt.Println("part1", ones*twos)
}

func (d Day8) parseLayers(input string, w int, h int) [][][]int {
	c, r := 0, 0
	var layers [][][]int
	var layer [][]int
	var row []int
	for _, ch := range strings.TrimSpace(input) {
		n := int(ch - '0')
		row = append(row, n)
		c = (c + 1) % w
		if c == 0 {
			layer = append(layer, row)
			row = make([]int, 0)
			r = (r + 1) % h
			if r == 0 {
				layers = append(layers, layer)
				layer = make([][]int, 0)
			}
		}
	}
	return layers
}

func (d Day8) Part2(input string) {
	//input = d.example2
	//w, h := 2, 2
	w, h := 25, 6
	layers := d.parseLayers(input, w, h)

	final := make([][]int, h)
	for i := range h {
		final[i] = make([]int, w)
	}

	for r := range h {
		for c := range w {
			for l := range layers {
				if layers[l][r][c] == 2 {
					continue
				}
				final[r][c] = layers[l][r][c]
				break
			}
		}
	}

	fmt.Println("part2")
	for _, row := range final {
		for _, n := range row {
			if n == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("\u2588")
			}
		}
		fmt.Println()
	}
}
