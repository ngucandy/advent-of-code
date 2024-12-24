package main

import (
	"bufio"
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
	input := string(bytes)

	part1(input)
	part2(input)
}

func part1(input string) {
	parts := strings.Split(strings.TrimSpace(strings.ReplaceAll(input, "\r\n", "\n")), "\n\n")
	valuesPart := parts[0]
	connectionsPart := parts[1]

	wireValues := make(map[string]int)
	scanner := bufio.NewScanner(strings.NewReader(strings.TrimSpace(valuesPart)))
	for scanner.Scan() {
		line := scanner.Text()
		parts = strings.Split(strings.TrimSpace(line), ": ")
		n, _ := strconv.Atoi(parts[1])
		wireValues[parts[0]] = n
	}

	operations := make(map[string][3]string)
	zs := make([]string, 0)
	re := regexp.MustCompile(`(...) ([^\s]+) (...) -> (...)`)
	scanner = bufio.NewScanner(strings.NewReader(strings.TrimSpace(connectionsPart)))
	for scanner.Scan() {
		line := scanner.Text()
		parts = re.FindStringSubmatch(line)
		operations[parts[4]] = [3]string{parts[2], parts[1], parts[3]}
		if strings.HasPrefix(parts[4], "z") {
			zs = append(zs, parts[4])
		}
	}
	slices.Sort(zs)
	slices.Reverse(zs)

	output := 0
	for _, z := range zs {
		output <<= 1
		wireValues[z] = compute(z, wireValues, operations)
		output |= wireValues[z]
	}
	slog.Info("Part 1:", "output", output)
}

func compute(wire string, wireValues map[string]int, operations map[string][3]string) int {
	if value, ok := wireValues[wire]; ok {
		return value
	}

	op := operations[wire][0]
	operand1 := operations[wire][1]
	operand2 := operations[wire][2]

	wireValues[operand1] = compute(operand1, wireValues, operations)
	wireValues[operand2] = compute(operand2, wireValues, operations)
	switch op {
	case "AND":
		return wireValues[operand1] & wireValues[operand2]
	case "OR":
		return wireValues[operand1] | wireValues[operand2]
	case "XOR":
		return wireValues[operand1] ^ wireValues[operand2]
	}

	panic("unknown operation: " + op)
}

func part2(input string) {

}
