package main

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
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
	total := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		secret, _ := strconv.Atoi(line)

		for range 2000 {
			secret = nextSecret(secret)
		}
		total += secret
	}
	slog.Info("Part 1:", "total", total)
}

func part2(input string) {

}

func nextSecret(secret int) int {
	next := secret << 6 // multiply by 64
	next ^= secret
	next &= 0xffffff // modulo 16777216

	secret = next
	next = secret >> 5
	next ^= secret
	next &= 0xffffff // modulo 16777216

	secret = next
	next = secret << 11 // multiply by 2048
	next ^= secret
	next &= 0xffffff // modulo 16777216

	return next
}
