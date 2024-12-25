package main

import (
	"bufio"
	"fmt"
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
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
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
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	connectionsPart := parts[1]

	operations := make(map[string][3]string)
	operationsMap := make(map[[3]string]string)
	zs := make([]string, 0)
	re := regexp.MustCompile(`(...) ([^\s]+) (...) -> (...)`)
	scanner := bufio.NewScanner(strings.NewReader(strings.TrimSpace(connectionsPart)))
	for scanner.Scan() {
		line := scanner.Text()
		parts = re.FindStringSubmatch(line)
		operations[parts[4]] = [3]string{parts[2], parts[1], parts[3]}
		if strings.HasPrefix(parts[4], "z") {
			zs = append(zs, parts[4])
		}
		var key [3]string
		if parts[1] < parts[3] {
			key = [3]string{parts[2], parts[1], parts[3]}
		} else {
			key = [3]string{parts[2], parts[3], parts[1]}
		}
		operationsMap[key] = parts[4]
	}
	slices.Sort(zs)
	slices.Reverse(zs)
	fmt.Println(operationsMap)

	adders := []Adder{
		Adder{
			ai: "x00",
			bi: "y00",
			so: operationsMap[[3]string{"XOR", "x00", "y00"}],
			co: operationsMap[[3]string{"AND", "x00", "y00"}],
		},
	}
	fmt.Printf("%+v\n", adders)
	for i := 1; i < 45; i++ {
		x := fmt.Sprintf("x%02d", i)
		y := fmt.Sprintf("y%02d", i)
		z := fmt.Sprintf("z%02d", i)
		adder := Adder{
			ai: x,
			bi: y,
			ci: adders[i-1].co,
			so: z,
			x1: operationsMap[[3]string{"XOR", x, y}],
			a1: operationsMap[[3]string{"AND", x, y}],
		}
		var key [3]string
		if adder.ci < adder.x1 {
			key = [3]string{"XOR", adder.ci, adder.x1}
		} else {
			key = [3]string{"XOR", adder.x1, adder.ci}
		}
		if operationsMap[key] != adder.so {
			fmt.Printf("%+v\n", adder)
			panic(fmt.Sprintf("mismatch: %v == %v; should be %v", key, operationsMap[key], adder.so))
		}

		if adder.x1 < adder.ci {
			key = [3]string{"AND", adder.x1, adder.ci}
		} else {
			key = [3]string{"AND", adder.ci, adder.x1}
		}
		adder.a2 = operationsMap[key]

		if adder.a1 < adder.a2 {
			key = [3]string{"OR", adder.a1, adder.a2}
		} else {
			key = [3]string{"OR", adder.a2, adder.a1}
		}
		adder.co = operationsMap[key]
		adders = append(adders, adder)

		fmt.Printf("%+v\n", adder)
	}

	// x09 XOR y09 -> qwf
	// x09 AND y09 -> cnk
	//
	// y14 AND x14 -> z14
	// ndq XOR rkm -> vhm
	//
	// snv OR jgq -> z27
	// kqw XOR kqj -> mps
	//
	// trn AND gpm -> z39
	// gpm XOR trn -> msq
	output := []string{"qwf", "cnk", "z14", "vhm", "z27", "mps", "z39", "msq"}
	slices.Sort(output)
	slog.Info("Part 2:", "output", strings.Join(output, ","))
}

type Adder struct {
	ai string
	bi string
	ci string
	so string // ci ^ x1
	co string // a1 | a2
	x1 string // a ^ b
	a1 string // a & b
	a2 string // x1 & ci
}
