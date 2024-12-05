package main

import (
	"bufio"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"
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

	rexpOrderRule := regexp.MustCompile(`\d\d\|\d\d`)
	mapRules := make(map[string]bool)
	var pageUpdates [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if rexpOrderRule.MatchString(line) {
			mapRules[line] = true
			continue
		}
		pages := strings.Split(line, ",")
		pageUpdates = append(pageUpdates, pages)
	}

	part1(mapRules, pageUpdates)
}

func part1(rules map[string]bool, updates [][]string) {
	sum := 0
UpdatesLoop:
	for _, update := range updates {
		slog.Info("Processing update:", "update", update)
		for i := range update {
			for a := range i {
				r := update[i] + "|" + update[a]
				if _, ok := rules[r]; ok {
					slog.Info("Violation:", "rule", r)
					continue UpdatesLoop
				}
			}
		}
		n, err := strconv.Atoi(update[len(update)/2])
		if err != nil {
			panic(err)
		}
		slog.Info("Correct:", "middle", n)
		sum += n
	}
	slog.Info("Part 1:", "sum", sum)
}
