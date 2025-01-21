package aoc2024

import (
	"slices"
	"strings"
)

func init() {
	Days["23"] = Day23{}
}

type Day23 struct {
	eg1, eg2 string
}

func (d Day23) Part1(input string) any {
	connections := make(map[string][]string)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "-")
		c1, c2 := parts[0], parts[1]
		connections[c1] = append(connections[c1], c2)
		connections[c2] = append(connections[c2], c1)
	}

	triples := make(map[[3]string]struct{})
	for c1, conns := range connections {
		for i, c2 := range conns[:len(conns)-1] {
			for _, c3 := range conns[i+1:] {
				// only looking for sets of 3 where at least 1 starts with 't'
				// need to ensure c2 is connected to c3 too
				if c1[0] != 't' && c2[0] != 't' && c3[0] != 't' || !slices.Contains(connections[c2], c3) {
					continue
				}
				triple := []string{c1, c2, c3}
				slices.Sort(triple)
				triples[[3]string{triple[0], triple[1], triple[2]}] = struct{}{}
			}
		}
	}

	return len(triples)
}

func (d Day23) Part2(input string) any {
	connections := make(map[string][]string)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "-")
		c1, c2 := parts[0], parts[1]
		connections[c1] = append(connections[c1], c2)
		connections[c2] = append(connections[c2], c1)
	}

	maxSize := 0
	var maxSet []string
	for c1 := range connections {
		connectedSet := d.getConnected(c1, []string{c1}, connections)
		if len(connectedSet) > maxSize {
			maxSize = len(connectedSet)
			maxSet = connectedSet
		}
	}

	slices.Sort(maxSet)
	return strings.Join(maxSet, ",")
}

func (d Day23) getConnected(c1 string, set []string, connections map[string][]string) []string {
	// loop through all of c1's connections
	for _, c2 := range connections[c1] {
		if slices.Contains(set, c2) {
			// we already processed that c1 and c2 are connected
			continue
		}
		skip := false
		// ensure c2 is connected to all of c1's connections
		for _, c3 := range set {
			if !slices.Contains(connections[c3], c2) {
				// c2 is not connected to one of c1's connections
				skip = true
				break
			}
		}
		if skip {
			continue
		}
		// c2 and c1 share a connection
		// add c2 to the set and recursively look through c2's connections
		return d.getConnected(c2, append(set, c2), connections)
	}
	return set
}
