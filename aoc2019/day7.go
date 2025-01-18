package aoc2019

import (
	"fmt"
	"github.com/ngucandy/advent-of-code/internal/helpers"
	"slices"
	"strconv"
	"strings"
)

func init() {
	Days["7"] = Day7{
		`3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0`,
		`3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0`,
		`3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0`,
		`3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5`,
		`3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10`,
	}
}

type Day7 struct {
	example1, example2, example3, example4, example5 string
}

func (d Day7) Part1(input string) {
	var memory []int
	for _, code := range strings.Split(strings.TrimSpace(input), ",") {
		n, _ := strconv.Atoi(code)
		memory = append(memory, n)
	}

	maxSignal := 0
	var maxPhases []int
	phaseCombos := helpers.CartesianProductN([]int{0, 1, 2, 3, 4}, 5)
PhaseCombo:
	for _, phases := range phaseCombos {
		// skip combos with repeat phases
		counts := make([]int, len(phases))
		for _, phase := range phases {
			counts[phase]++
		}
		for _, count := range counts {
			if count > 1 {
				continue PhaseCombo
			}
		}

		signal := 0
		for _, phase := range phases {
			c := NewIntcodeComputer(slices.Clone(memory), []int{phase, signal})
			for c.Step() {
			}
			signal = c.output[0]
		}
		if signal > maxSignal {
			maxSignal = signal
			maxPhases = phases
		}
	}
	fmt.Println("part1", maxSignal, maxPhases)
}

func (d Day7) Part2(input string) {
	var memory []int
	for _, code := range strings.Split(strings.TrimSpace(input), ",") {
		n, _ := strconv.Atoi(code)
		memory = append(memory, n)
	}

	maxSignal := 0
	var maxPhases []int
	phaseCombos := helpers.CartesianProductN([]int{5, 6, 7, 8, 9}, 5)
PhaseCombo:
	for _, phases := range phaseCombos {
		// skip combos with repeat phases
		counts := make(map[int]int)
		for _, phase := range phases {
			counts[phase]++
		}
		for _, count := range counts {
			if count > 1 {
				continue PhaseCombo
			}
		}

		var computers []*IntcodeComputer
		for _, phase := range phases {
			computers = append(computers, NewIntcodeComputer(slices.Clone(memory), []int{phase}))
		}
		i := 0
		signal := 0
		computers[i].input = append(computers[i].input, 0)
		for {
			halted := !computers[i].Step()
			if halted {
				// if last computer halted, we're done
				if i == len(computers)-1 {
					break
				} else {
					i = (i + 1) % len(computers)
					continue
				}
			}

			// if step didn't produce output, continue running on same computer
			if len(computers[i].output) == 0 {
				continue
			}

			// if step produced output, send it to next computer
			output := computers[i].output[0]
			computers[i].output = computers[i].output[1:]
			if i == len(computers)-1 {
				signal = output
			}
			ni := (i + 1) % len(computers)
			computers[ni].input = append(computers[ni].input, output)
			i = ni
		}
		if signal > maxSignal {
			maxSignal = signal
			maxPhases = phases
		}
	}
	fmt.Println("part2", maxSignal, maxPhases)
}
