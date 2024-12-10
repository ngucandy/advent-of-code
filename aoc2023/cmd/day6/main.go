package main

import (
	"bufio"
	"log/slog"
	"os"
	"regexp"
	"strconv"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, _ := os.Open(infile)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	rexpDigits := regexp.MustCompile(`\d+`)
	times := make([]int, 0)
	distances := make([]int, 0)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	nums := rexpDigits.FindAllString(line, -1)
	for _, num := range nums {
		n, _ := strconv.Atoi(num)
		times = append(times, n)
	}

	scanner.Scan()
	line = scanner.Text()
	nums = rexpDigits.FindAllString(line, -1)
	for _, num := range nums {
		n, _ := strconv.Atoi(num)
		distances = append(distances, n)
	}

	part1(times, distances)
	part2(times, distances)
}

func part1(times []int, distances []int) {
	ans := 1

	for i, time := range times {
		wins := 0
		distance := distances[i]
		for speed := range time {
			duration := time - speed
			if (speed * duration) > distance {
				wins++
			}
		}
		ans *= wins
	}

	slog.Info("Part 1:", "answer", ans)
}

func part2(times []int, distances []int) {
	ans := 0

	s := ""
	for _, time := range times {
		s += strconv.Itoa(time)
	}
	time, _ := strconv.Atoi(s)

	s = ""
	for _, distance := range distances {
		s += strconv.Itoa(distance)
	}
	distance, _ := strconv.Atoi(s)

	for speed := range time {
		duration := time - speed
		if (speed * duration) > distance {
			ans++
		}
	}

	slog.Info("Part 2:", "answer", ans)
}
