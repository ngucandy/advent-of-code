package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, _ := os.Open(infile)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	rexpRegister := regexp.MustCompile(`Register (.): ([\d]+)`)
	rexpProgram := regexp.MustCompile(`\d`)
	register := make(map[string]int)
	program := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if rexpRegister.MatchString(line) {
			matches := rexpRegister.FindStringSubmatch(line)
			val, _ := strconv.Atoi(matches[2])
			register[matches[1]] = val
			continue
		}
		matches := rexpProgram.FindAllString(line, -1)
		for _, match := range matches {
			val, _ := strconv.Atoi(match)
			program = append(program, val)
		}
	}

	part1(program, register)
	part2(program, register)
}

const (
	A = "A"
	B = "B"
	C = "C"
)

var (
	opMap = map[int]func(int, map[string]int, int) (int, string){
		0: adv,
		1: bxl,
		2: bst,
		3: jnz,
		4: bxc,
		5: out,
		6: bdv,
		7: cdv,
	}
)

func part1(program []int, register map[string]int) {
	ip := 0
	output := []string{}
	for ip < len(program) {
		nextIp, s := opMap[program[ip]](program[ip+1], register, ip)
		if len(s) > 0 {
			output = append(output, s)
		}
		ip = nextIp
	}
	slog.Info("Part 1: ", "output", strings.Join(output, ","))
}

func part2(program []int, _ map[string]int) {
	minA := 1_200_000_000
	maxA := minA + 100_000_000
	wg := &sync.WaitGroup{}
	wg.Add(maxA - minA)
	for a := minA; a < maxA; a++ {
		go func(val int) {
			ip := 0
			i := 0
			output := []int{}
			register := make(map[string]int)
			register[A] = val
			register[B] = 0
			register[C] = 0
			for ip < len(program) {
				nextIp, s := opMap[program[ip]](program[ip+1], register, ip)
				if len(s) > 0 {
					n, _ := strconv.Atoi(s)
					if n != program[i] {
						break
					}
					output = append(output, n)
					i++
					if i == len(program) {
						slog.Info("Part 2: ", "output", output, "a", val)
						//a = maxA
						break
					}
				}
				ip = nextIp
			}
			wg.Done()
		}(a)
	}
	wg.Wait()
}

func adv(operand int, register map[string]int, ip int) (int, string) {
	numerator := register[A]
	denominator := math.Pow(2, float64(comboOperand(operand, register)))
	register[A] = int(float64(numerator) / denominator)
	return ip + 2, ""
}

func bxl(operand int, register map[string]int, ip int) (int, string) {
	register[B] ^= operand
	return ip + 2, ""
}

func bst(operand int, register map[string]int, ip int) (int, string) {
	register[B] = comboOperand(operand, register) % 8
	return ip + 2, ""
}

func jnz(operand int, register map[string]int, ip int) (int, string) {
	if register[A] == 0 {
		return ip + 2, ""
	}
	return operand, ""
}

func bxc(_ int, register map[string]int, ip int) (int, string) {
	register[B] ^= register[C]
	return ip + 2, ""
}

func out(operand int, register map[string]int, ip int) (int, string) {
	return ip + 2, strconv.Itoa(comboOperand(operand, register) % 8)
}

func bdv(operand int, register map[string]int, ip int) (int, string) {
	numerator := register[A]
	denominator := math.Pow(2, float64(comboOperand(operand, register)))
	register[B] = int(float64(numerator) / denominator)
	return ip + 2, ""
}

func cdv(operand int, register map[string]int, ip int) (int, string) {
	numerator := register[A]
	denominator := math.Pow(2, float64(comboOperand(operand, register)))
	register[C] = int(float64(numerator) / denominator)
	return ip + 2, ""
}

func comboOperand(operand int, register map[string]int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return register[A]
	case 5:
		return register[B]
	case 6:
		return register[C]
	default:
		panic(fmt.Sprint("Invalid combo operand: ", operand))
	}
}
