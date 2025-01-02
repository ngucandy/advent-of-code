package main

import (
	"log/slog"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	bytes, _ := os.ReadFile(infile)
	input := strings.ReplaceAll(string(bytes), "\r\n", "\n")

	part1(input)
	part2(input)
}

func part1(input string) {
	steps := strings.Split(input, ",")
	sum := 0
	for _, step := range steps {
		sum += hash(step)
	}
	slog.Info("Part 1:", "sum", sum)
}

func hash(s string) int {
	current := 0
	for _, ch := range s {
		current += int(ch)
		current *= 17
		current %= 256
	}
	return current
}

func part2(input string) {
	re := regexp.MustCompile(`(.+)([=-])(\d*)`)
	steps := strings.Split(input, ",")
	boxes := make([][][2]string, 256)

	for _, step := range steps {
		matches := re.FindStringSubmatch(step)
		label := matches[1]
		op := matches[2]
		box := hash(label)
		indexFunc := func(lens [2]string) bool {
			return lens[0] == label
		}
		if op == "=" {
			fl := matches[3]
			lens := [2]string{label, fl}
			if i := slices.IndexFunc(boxes[box], indexFunc); i != -1 {
				boxes[box][i] = lens
			} else {
				boxes[box] = append(boxes[box], lens)
			}
		} else { // op == "-"
			boxes[box] = slices.DeleteFunc(boxes[box], indexFunc)
		}
	}

	power := 0
	for i, box := range boxes {
		for slot, lens := range box {
			fl, _ := strconv.Atoi(lens[1])
			power += (i + 1) * (slot + 1) * fl
		}
	}
	slog.Info("Part 2:", "power", power)
}
