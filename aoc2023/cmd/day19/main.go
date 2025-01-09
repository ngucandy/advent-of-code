package main

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

const testInput = `px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	bytes, _ := os.ReadFile(infile)
	input := strings.ReplaceAll(string(bytes), "\r\n", "\n")

	part1(testInput)
	part1(input)
	part2(testInput)
	part2(input)
}

type Rule struct {
	Operand1    string
	Operation   string
	Operand2    int
	Destination string
}

func NewRule(raw string) *Rule {
	r := &Rule{}
	// e.g., s>2770:qs
	parts := strings.Split(raw, ":")
	if len(parts) == 1 {
		r.Destination = raw
		return r
	}
	r.Operand1 = string(parts[0][0])
	r.Operation = string(parts[0][1])
	r.Operand2, _ = strconv.Atoi(parts[0][2:])
	r.Destination = parts[1]
	return r
}

func (r *Rule) Apply(part map[string]int) (string, bool) {
	if len(r.Operation) == 0 {
		return r.Destination, true
	}
	partVal := part[r.Operand1]
	switch r.Operation {
	case "<":
		if partVal < r.Operand2 {
			return r.Destination, true
		}
	case ">":
		if partVal > r.Operand2 {
			return r.Destination, true
		}
	}
	return "", false
}

func part1(input string) {
	inputParts := strings.Split(input, "\n\n")

	workflows := buildWorkflows(inputParts[0])

	sum := 0
	rejected := 0
	for _, line := range strings.Split(inputParts[1], "\n") {
		var x, m, a, s int
		_, _ = fmt.Sscanf(line, "{x=%d,m=%d,a=%d,s=%d}", &x, &m, &a, &s)
		part := map[string]int{
			"x": x,
			"m": m,
			"a": a,
			"s": s,
		}
		q := []string{"in"}
		for len(q) > 0 {
			workflow := q[0]
			q = q[1:]
			if workflow == "A" {
				sum += part["x"] + part["m"] + part["a"] + part["s"]
				break
			}
			if workflow == "R" {
				rejected++
				break
			}
			rules := workflows[workflow]
			for _, rule := range rules {
				next, success := rule.Apply(part)
				if success {
					q = append(q, next)
					break
				}
			}
		}
	}

	slog.Info("Part 1:", "sum", sum)
}

func buildWorkflows(input string) map[string][]*Rule {
	workflows := make(map[string][]*Rule)
	for _, line := range strings.Split(input, "\n") {
		// e.g., qqz{s>2770:qs,m<1801:hdj,R}
		parts := strings.Split(line[:len(line)-1], "{")
		name := parts[0]
		rules := strings.Split(parts[1], ",")
		for _, rule := range rules {
			r := NewRule(rule)
			workflows[name] = append(workflows[name], r)
		}
	}
	return workflows
}

func part2(input string) {
	inputParts := strings.Split(input, "\n\n")

	workflows := buildWorkflows(inputParts[0])
	var paths [][]string

	q := [][]string{{"0", "in"}}
	for len(q) > 0 {
		n, _ := strconv.Atoi(q[0][0])
		workflow := q[0][1]
		path := q[0][2:]
		q = q[1:]

		if workflow == "A" {
			paths = append(paths, append(path, strconv.Itoa(n), "A"))
		}

		if workflow == "R" {
			continue
		}

		for i, rule := range workflows[workflow] {
			newpath := append(path, strconv.Itoa(n), workflow)
			next := append([]string{strconv.Itoa(i), rule.Destination}, newpath...)
			q = append(q, next)
		}
	}

	total := 0
	for _, path := range paths {
		part := map[string][]int{
			"x": {1, 4000},
			"m": {1, 4000},
			"a": {1, 4000},
			"s": {1, 4000},
		}
		for i := 1; i < len(path)-1; i += 2 {
			workflow := path[i]
			n, _ := strconv.Atoi(path[i+1])
			next := path[i+2]

			rules := workflows[workflow]
			for j, rule := range rules {
				if rule.Destination == next && n == j {
					if len(rule.Operation) > 0 {
						if rule.Operation == "<" {
							part[rule.Operand1][1] = min(part[rule.Operand1][1], rule.Operand2-1)
						} else { // >
							part[rule.Operand1][0] = max(part[rule.Operand1][0], rule.Operand2+1)
						}
					}
					break
				} else {
					if rule.Operation == "<" { // >=
						part[rule.Operand1][0] = max(part[rule.Operand1][0], rule.Operand2)
					} else { // <=
						part[rule.Operand1][1] = min(part[rule.Operand1][1], rule.Operand2)
					}
				}
			}
		}
		combinations := (part["x"][1] - part["x"][0] + 1) *
			(part["m"][1] - part["m"][0] + 1) *
			(part["a"][1] - part["a"][0] + 1) *
			(part["s"][1] - part["s"][0] + 1)
		total += combinations
	}
	slog.Info("Part 2:", "total combinations", total)
}
