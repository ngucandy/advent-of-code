package main

import (
	"encoding/csv"
	"io"
	"log/slog"
	"os"
	"slices"
	"strconv"
)

const (
	MIN_DIFF = 1
	MAX_DIFF = 3
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

	safe1, safe2 := countSafeReports(file)
	slog.Info("Report count part 1:", "safe", safe1)
	slog.Info("Report count part 2:", "safe", safe2)
}

func countSafeReports(file *os.File) (int, int) {
	r := csv.NewReader(file)
	r.Comma = ' '
	r.FieldsPerRecord = -1

	safe1 := 0
	safe2 := 0
	for {
		report, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		// part 1
		if isSafe(report) {
			safe1++
			safe2++
			continue
		}

		// part 2
		// brute force, create new slices with the ith element missing
		for i := range len(report) {
			if isSafe(slices.Concat(report[:i], report[i+1:])) {
				safe2++
				break
			}
		}
	}
	return safe1, safe2
}

func isSafe(report []string) bool {
	var levelDiffs []int
	for i, v := range report[1:] {
		n, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			panic(err)
		}
		level := int(n)

		nPrev, err := strconv.ParseInt(report[i], 10, 32)
		if err != nil {
			panic(err)
		}
		levelPrev := int(nPrev)

		levelDiffs = append(levelDiffs, level-levelPrev)
	}

	ascCount := 0
	for _, v := range levelDiffs {
		if v > 0 {
			ascCount++
		} else {
			v = -v
		}
		if v < MIN_DIFF || v > MAX_DIFF {
			slog.Info("Level difference out of range:", "report", report, "diffs", levelDiffs)
			return false
		}
	}
	if ascCount != 0 && ascCount != len(levelDiffs) {
		slog.Info("Levels not all ascending or descending:", "report", report, "diffs", levelDiffs, "count", ascCount)
		return false
	}
	slog.Info("Levels are safe:", "report", report, "diffs", levelDiffs)
	return true
}
