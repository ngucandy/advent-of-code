package main

import (
	"bufio"
	"log/slog"
	"os"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, _ := os.Open(infile)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	m := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, 0, len(line))
		for _, c := range line {
			row = append(row, int(c-'0'))
		}
		m = append(m, row)
	}

	trailHeads := [][2]int{}
	trailEnds := [][2]int{}
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			switch m[y][x] {
			case 0:
				trailHeads = append(trailHeads, [2]int{x, y})
			case 9:
				trailEnds = append(trailEnds, [2]int{x, y})
			}
		}
	}

	part1(m, trailHeads, trailEnds)
	part2(m, trailHeads, trailEnds)
}

func part1(tmap [][]int, trailHeads [][2]int, trailEnds [][2]int) {
	total := 0

	for _, trailHead := range trailHeads {
		for _, trailEnd := range trailEnds {
			if isReachable(trailHead, trailEnd, tmap) {
				total++
			}
		}
	}
	slog.Info("Part 1:", "total", total)
}

func part2(tmap [][]int, trailHeads [][2]int, trailEnds [][2]int) {
	total := 0

	for _, trailHead := range trailHeads {
		for _, trailEnd := range trailEnds {
			total += countPaths(trailHead, trailEnd, tmap)
		}
	}
	slog.Info("Part 2:", "total", total)

}

var dirs = [4][2]int{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func isReachable(start [2]int, end [2]int, tmap [][]int) bool {
	if start == end {
		return true
	}
	for _, dir := range dirs {
		next := [2]int{start[0] + dir[0], start[1] + dir[1]}
		if next[0] < 0 || next[0] >= len(tmap[0]) || next[1] < 0 || next[1] >= len(tmap) {
			continue
		}
		if tmap[next[1]][next[0]]-tmap[start[1]][start[0]] != 1 {
			continue
		}
		if isReachable(next, end, tmap) {
			return true
		}
	}
	return false
}

func countPaths(start [2]int, end [2]int, tmap [][]int) int {
	if start == end {
		return 1
	}
	n := 0
	for _, dir := range dirs {
		next := [2]int{start[0] + dir[0], start[1] + dir[1]}
		if next[0] < 0 || next[0] >= len(tmap[0]) || next[1] < 0 || next[1] >= len(tmap) {
			continue
		}
		if tmap[next[1]][next[0]]-tmap[start[1]][start[0]] != 1 {
			continue
		}
		n += countPaths(next, end, tmap)
	}
	return n
}
