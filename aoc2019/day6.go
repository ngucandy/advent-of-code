package aoc2019

import (
	"strings"
)

func init() {
	Days["6"] = Day6{
		`COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`,
		`COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`,
	}
}

type Day6 struct {
	example1 string
	example2 string
}

func (d Day6) Part1(input string) any {
	orbits := make(map[string]string)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ")")
		orbits[parts[1]] = parts[0]
	}

	total := 0
	cache := make(map[string]int)
	for orbiter := range orbits {
		total += d.countOrbits(orbiter, orbits, cache)
	}
	return total
}

type Day6Tuple struct {
	name string
	cost int
}

func (d Day6) Part2(input string) any {
	orbits := make(map[string][]string)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ")")
		orbits[parts[0]] = append(orbits[parts[0]], parts[1])
		orbits[parts[1]] = append(orbits[parts[1]], parts[0])
	}

	q := []Day6Tuple{{"YOU", 0}}
	seen := make(map[string]int)
	for len(q) > 0 {
		t := q[0]
		q = q[1:]

		if t.name == "SAN" {
			return t.cost - 2
		}

		if seenCost, found := seen[t.name]; found && seenCost < t.cost {
			continue
		}
		seen[t.name] = t.cost

		for _, neighbor := range orbits[t.name] {
			q = append(q, Day6Tuple{neighbor, t.cost + 1})
		}
	}
	return nil
}

func (d Day6) countOrbits(orbiter string, orbits map[string]string, cache map[string]int) int {
	orbitee, exists := orbits[orbiter]
	if !exists {
		return 0
	}

	if count, exists := cache[orbiter]; exists {
		return count
	}

	orbitCount := 1 + d.countOrbits(orbitee, orbits, cache)
	cache[orbiter] = orbitCount
	return orbitCount
}
