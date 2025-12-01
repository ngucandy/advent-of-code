package aoc2025

import (
	"fmt"
	"strings"
)

func init() {
	Days["1"] = Day1{
		example: `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`,
	}
}

type Day1 struct {
	example string
}

func (d Day1) Part1(input string) any {
	size := 100
	count := 0
	dial := 50
	var rotation string
	var distance int
	for _, line := range strings.Split(input, "\n") {
		_, err := fmt.Sscanf(line, "%1s%d", &rotation, &distance)
		if err != nil {
			fmt.Printf("error scanning input data: %v\n", err)
			break
		}
		if rotation == "L" {
			distance = -distance + size
		}
		dial = (dial + distance) % size
		if dial == 0 {
			count++
		}
	}
	return count
}

func (d Day1) Part2(input string) any {
	size := 100
	count := 0
	dial := 50
	var rotation string
	var distance int
	for _, line := range strings.Split(input, "\n") {
		_, err := fmt.Sscanf(line, "%1s%d", &rotation, &distance)
		if err != nil {
			fmt.Printf("error scanning input data: %v\n", err)
			break
		}

		// increase count for every full rotation
		if distance > size {
			count += distance / size
			distance %= size
		}

		if rotation == "L" {
			if dial > 0 && distance >= dial {
				count++
			}
			distance = -distance + size
		} else {
			if distance >= size-dial {
				count++
			}
		}
		dial = (dial + distance) % size
		fmt.Println(line, dial, count)
	}
	return count
}
