package aoc2019

import (
	"slices"
	"strings"

	"github.com/ngucandy/advent-of-code/internal/helpers"
)

func init() {
	Days["16"] = &Day16{
		eg1: `12345678`,
		eg2: `80871224585914546619083218645595`,
	}
}

type Day16 struct {
	eg1, eg2 string
}

func (d Day16) Part1(input string) any {
	var signal []int
	for _, ch := range strings.TrimSpace(input) {
		signal = append(signal, int(ch-'0'))
	}
	base := []int{0, 1, 0, -1}
	var patterns [][]int
	for i := range signal {
		var pattern []int
		for _, n := range base {
			pattern = append(pattern, slices.Repeat([]int{n}, i+1)...)
		}
		patterns = append(patterns, pattern)
	}

	var output []int
	for range 100 {
		output = make([]int, len(signal))
		for o := range signal {
			n := 0
			for i, j := 0, 1; i < len(signal); i, j = i+1, (j+1)%len(patterns[o]) {
				n += signal[i] * patterns[o][j]
			}
			output[o] = helpers.AbsInt(n) % 10
		}
		signal = output
	}

	return helpers.Join(output[:8], "")
}

func (d Day16) Part2(input string) any {
	input = d.eg2
	var signal []int
	for range 4 {
		for _, ch := range strings.TrimSpace(input) {
			signal = append(signal, int(ch-'0'))
		}
	}

	var output []int
	for range 1 {
		output = make([]int, len(signal))
		// loop output digits
		for o := range signal {
			n := 0
			p := 1
			// Signal digits up to `o` will always be multiplied by 0, so we
			// can skip them when computing `n`. Signal digits from `o` to `o+1*2`
			for i := o; i < len(signal); i += (o + 1) * 2 {
				for j := 0; j < o+1 && i+j < len(signal); j++ {
					//fmt.Printf("%d*%d + ", signal[i+j], p)
					n += signal[i+j] * p
				}
				p = -p
			}
			//fmt.Println()
			output[o] = helpers.AbsInt(n) % 10
		}
		signal = output
	}

	return strings.Join([]string{helpers.Join(output[:len(output)/2], ""), helpers.Join(output[len(output)/2:], "")}, " ")
}
