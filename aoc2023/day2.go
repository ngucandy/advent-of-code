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

	sum := 0
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
			sum += int(id)
		} else {
			slog.Info("Impossible game")
		}
	}

	slog.Info("Sum of possible game IDs:", "sum", sum)
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
