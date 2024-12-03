package main

import (
	"bufio"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	reDelimReveals = regexp.MustCompile("[:;]")
	reDelimCubes   = regexp.MustCompile("[,]")
	reGame         = regexp.MustCompile("Game ([\\d]+)")
	reBlue         = regexp.MustCompile("([\\d]+) blue")
	reRed          = regexp.MustCompile("([\\d]+) red")
	reGreen        = regexp.MustCompile("([\\d]+) green")

	allCubes = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			slog.Error("Error closing file:", "error", err)
		}
	}(file)

	sumPossible := 0
	sumPower := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		reveals := reDelimReveals.Split(line, -1)

		match := reGame.FindStringSubmatch(reveals[0])
		id, err := strconv.ParseInt(match[1], 10, 32)
		if err != nil {
			panic(err)
		}
		slog.Info("Game:", "ID", id)

		if isGamePossible(reveals) {
			slog.Info("Possible game")
			sumPossible += int(id)
		} else {
			slog.Info("Impossible game")
		}

		power := computePower(line)
		slog.Info("Game power:", "power", power)
		sumPower += power
	}
	slog.Info("Sum of possible game IDs:", "sum", sumPossible)
	slog.Info("Sum of game powers:", "sum", sumPower)
}

func isGamePossible(reveals []string) bool {
	for _, reveal := range reveals[1:] {
		reveal = strings.TrimSpace(reveal)
		cubes := reDelimCubes.Split(reveal, -1)
		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			split := strings.Split(cube, " ")
			n, err := strconv.ParseInt(split[0], 10, 32)
			if err != nil {
				panic(err)
			}
			maxCubes, ok := allCubes[split[1]]
			if !ok {
				panic(reveals)
			}
			if maxCubes-int(n) < 0 {
				slog.Info("Impossible game:", "cube", cube)
				return false
			}
		}
	}
	return true
}

func computePower(line string) int {
	matches := reBlue.FindAllStringSubmatch(line, -1)
	maxBlue := findMax(matches)
	matches = reGreen.FindAllStringSubmatch(line, -1)
	maxGreen := findMax(matches)
	matches = reRed.FindAllStringSubmatch(line, -1)
	maxRed := findMax(matches)
	return maxRed * maxGreen * maxBlue
}

func findMax(matches [][]string) int {
	maxN := 0
	for _, match := range matches {
		c, err := strconv.ParseInt(match[1], 10, 32)
		if err != nil {
			panic(err)
		}
		n := int(c)
		if n > maxN {
			maxN = n
		}
	}
	return maxN
}
