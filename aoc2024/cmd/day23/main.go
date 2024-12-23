package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strings"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	bytes, _ := os.ReadFile(infile)
	input := string(bytes)

	part1(input)
	part2(input)
}

func part1(input string) {
	neighbors := make(map[string][]string)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		var left, right string
		_, _ = fmt.Sscanf(line, "%2s-%2s", &left, &right)
		neighbors[left] = append(neighbors[left], right)
		neighbors[right] = append(neighbors[right], left)
	}

	for node := range neighbors {
		slices.Sort(neighbors[node])
		neighbors[node] = slices.Compact(neighbors[node])
	}

	triples := make(map[[3]string]bool)
	for node1 := range neighbors {
		for _, node2 := range neighbors[node1] {
			if node1 == node2 {
				continue
			}
			for _, node3 := range neighbors[node2] {
				if node2 == node3 {
					continue
				}
				if slices.Contains(neighbors[node1], node3) {
					tuple := []string{node1, node2, node3}
					startsWithT := func(s string) bool {
						return strings.HasPrefix(s, "t")
					}
					if slices.ContainsFunc(tuple, startsWithT) {
						slices.Sort(tuple)
						triples[[3]string{tuple[0], tuple[1], tuple[2]}] = true
					}
				}
			}
		}
	}
	slog.Info("Part 1:", "triples", len(triples))
}

func part2(input string) {
}
