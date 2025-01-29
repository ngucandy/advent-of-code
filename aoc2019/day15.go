package aoc2019

import (
	"fmt"
)

func init() {
	Days["15"] = &Day15{}
}

type Day15 struct {
	eg1, eg2 string
}

func (d Day15) Part1(input string) any {
	program := ParseIntcodeProgram(input)
	computer := NewIntcodeComputer(program, []int{})

	north, south, west, east := 1, 2, 3, 4
	directions := map[int][2]int{
		north: {-1, 0},
		south: {1, 0},
		west:  {0, -1},
		east:  {0, 1},
	}

	// bfs
	type state struct {
		r, c  int
		steps int
		ic    *IntcodeComputer
	}
	var steps int
	q := []state{{0, 0, 0, computer}}
	seen := make(map[[2]int]bool)
queue:
	for len(q) > 0 {
		s := q[0]
		q = q[1:]

		for cmd, dir := range directions {
			nr, nc := s.r+dir[0], s.c+dir[1]
			if seen[[2]int{nr, nc}] {
				continue
			}
			seen[[2]int{nr, nc}] = true

			nic := s.ic.Clone()
			out := nic.Run(cmd)
			switch out {
			case 0:
				// hit a wall
				continue
			case 1:
				q = append(q, state{nr, nc, s.steps + 1, nic})
				continue
			case 2:
				// reached destination
				steps = s.steps + 1
				break queue
			default:
				panic(fmt.Sprintf("unknown status: %d", out))
			}
		}
	}
	return steps
}

func (d Day15) Part2(input string) any {
	return "no answer yet"
}
