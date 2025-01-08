package main

import (
	"fmt"
	"log/slog"
	"os"
	"regexp"
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
	re := regexp.MustCompile(`^(.+)([<>])(.+):(.+)$`)
	r := &Rule{}
	match := re.FindStringSubmatch(raw)
	if match == nil {
		r.Destination = raw
		return r
	}
	r.Operand1 = match[1]
	r.Operation = match[2]
	r.Operand2, _ = strconv.Atoi(match[3])
	r.Destination = match[4]
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

	workflows := make(map[string][]*Rule)
	reWorkflow := regexp.MustCompile(`^(.+){(.+)}$`)
	for _, line := range strings.Split(inputParts[0], "\n") {
		match := reWorkflow.FindStringSubmatch(line)
		rules := strings.Split(match[2], ",")
		for _, rule := range rules {
			r := NewRule(rule)
			workflows[match[1]] = append(workflows[match[1]], r)
		}
	}

	//for name, rules := range workflows {
	//	s := fmt.Sprintf("%s:", name)
	//	for _, rule := range rules {
	//		s = fmt.Sprintf("%s %v", s, rule)
	//	}
	//	fmt.Println(s)
	//}

	sum := 0
	rejected := 0
	for _, line := range strings.Split(inputParts[1], "\n") {
		var x, m, a, s int
		fmt.Sscanf(line, "{x=%d,m=%d,a=%d,s=%d}", &x, &m, &a, &s)
		part := map[string]int{
			"x": x,
			"m": m,
			"a": a,
			"s": s,
		}
		//fmt.Println(part)
		q := []string{"in"}
		for len(q) > 0 {
			workflow := q[0]
			q = q[1:]
			//fmt.Println(workflow)
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

func part2(input string) {

}
