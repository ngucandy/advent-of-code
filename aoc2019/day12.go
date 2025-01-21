package aoc2019

import (
	"cmp"
	"github.com/ngucandy/advent-of-code/internal/helpers"
	"slices"
	"strconv"
	"strings"
)

func init() {
	Days["12"] = &Day12{
		eg1: `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`,
		eg2: `<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>`,
	}
}

type Day12 struct {
	eg1, eg2 string
}

func (d Day12) Part1(input string) any {
	var positions, velocities [][3]int
	for _, line := range strings.Split(input, "\n") {
		var position, velocity [3]int
		parts := strings.Split(line[1:len(line)-1], ", ")
		position[0], _ = strconv.Atoi(parts[0][2:])
		position[1], _ = strconv.Atoi(parts[1][2:])
		position[2], _ = strconv.Atoi(parts[2][2:])
		positions = append(positions, position)
		velocities = append(velocities, velocity)
	}

	for range 1000 {
		// apply gravity
		for i, pos1 := range positions[:len(positions)-1] {
			for j, pos2 := range positions[i+1:] {
				velocities[i][0] += cmp.Compare(pos2[0], pos1[0])
				velocities[i+j+1][0] += cmp.Compare(pos1[0], pos2[0])
				velocities[i][1] += cmp.Compare(pos2[1], pos1[1])
				velocities[i+j+1][1] += cmp.Compare(pos1[1], pos2[1])
				velocities[i][2] += cmp.Compare(pos2[2], pos1[2])
				velocities[i+j+1][2] += cmp.Compare(pos1[2], pos2[2])
			}
		}

		// apply velocity
		for i := range positions {
			positions[i][0] += velocities[i][0]
			positions[i][1] += velocities[i][1]
			positions[i][2] += velocities[i][2]
		}
	}

	sum := 0
	for i := range positions {
		px, py, pz := helpers.AbsInt(positions[i][0]), helpers.AbsInt(positions[i][1]), helpers.AbsInt(positions[i][2])
		vx, vy, vz := helpers.AbsInt(velocities[i][0]), helpers.AbsInt(velocities[i][1]), helpers.AbsInt(velocities[i][2])
		sum += (px + py + pz) * (vx + vy + vz)
	}
	return sum
}

func (d Day12) Part2(input string) any {
	var positions, velocities, oPositions, oVelocities [][]int
	for _, line := range strings.Split(input, "\n") {
		position := make([]int, 3)
		velocity := make([]int, 3)
		parts := strings.Split(line[1:len(line)-1], ", ")
		position[0], _ = strconv.Atoi(parts[0][2:])
		position[1], _ = strconv.Atoi(parts[1][2:])
		position[2], _ = strconv.Atoi(parts[2][2:])
		positions = append(positions, position)
		velocities = append(velocities, velocity)
		oPositions = append(oPositions, slices.Clone(position))
		oVelocities = append(oVelocities, slices.Clone(velocity))
	}

	columnEquals := func(c int, a, b [][]int) bool {
		for r := range a {
			if a[r][c] != b[r][c] {
				return false
			}
		}
		return true
	}

	// Look for independent cycles of x, y and z across all positions and
	// velocities. In other words, when all x values (index 0) return to
	// their original values, that counts as a cycle for x. Same applies
	// for y and z. Look for 2 cycles before stopping.
	var cycles [3][]int
step:
	for n := 1; len(cycles[0]) < 2 || len(cycles[1]) < 2 || len(cycles[2]) < 2; n++ {
		// apply gravity
		for i, pos1 := range positions[:len(positions)-1] {
			for j, pos2 := range positions[i+1:] {
				velocities[i][0] += cmp.Compare(pos2[0], pos1[0])
				velocities[i+j+1][0] += cmp.Compare(pos1[0], pos2[0])
				velocities[i][1] += cmp.Compare(pos2[1], pos1[1])
				velocities[i+j+1][1] += cmp.Compare(pos1[1], pos2[1])
				velocities[i][2] += cmp.Compare(pos2[2], pos1[2])
				velocities[i+j+1][2] += cmp.Compare(pos1[2], pos2[2])
			}
		}

		// apply velocity
		for i := range positions {
			positions[i][0] += velocities[i][0]
			positions[i][1] += velocities[i][1]
			positions[i][2] += velocities[i][2]
		}

		for c := range cycles {
			if len(cycles[c]) >= 2 {
				continue
			}
			if !columnEquals(c, oPositions, positions) || !columnEquals(c, oVelocities, velocities) {
				continue step
			}
			cycles[c] = append(cycles[c], n)
		}
	}

	// use the difference between the first and second cycles as the step
	// value for repeating the cycle
	return helpers.LCM(cycles[0][1]-cycles[0][0], cycles[1][1]-cycles[1][0], cycles[2][1]-cycles[2][0])
}
