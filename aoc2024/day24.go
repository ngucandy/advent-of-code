package aoc2024

import (
	"fmt"
	"slices"
	"strings"
)

func init() {
	DayMap["24"] = Day24{}
}

type Day24 struct {
	eg1, eg2 string
}

func (d Day24) Part1(input string) {
	sections := strings.Split(input, "\n\n")
	suppliers := make(map[string]Supplier)
	for _, line := range strings.Split(sections[0], "\n") {
		// e.g., x00: 1
		var name string
		var val int
		_, _ = fmt.Sscanf(line, "%s %d", &name, &val)
		name = name[:len(name)-1] // chop off trailing ':'
		suppliers[name] = NewSwitch(val, name)
	}

	// read in all gates
	var gates []*Gate
	var zs []string
	for _, line := range strings.Split(sections[1], "\n") {
		// e.g., y06 AND x06 -> pqj
		var operand1, operand2, op, target string
		_, _ = fmt.Sscanf(line, "%s %s %s -> %s", &operand1, &op, &operand2, &target)
		gate := NewGate(op, target, operand1, operand2, -1)
		gates = append(gates, gate)
		suppliers[target] = gate
		if target[0] == 'z' {
			zs = append(zs, target)
		}
	}

	// connect the gates
	for _, gate := range gates {
		for _, missing := range gate.MissingConnections() {
			gate.ConnectInput(suppliers[missing])
		}
	}

	slices.Sort(zs)
	slices.Reverse(zs)
	output := 0
	for _, z := range zs {
		output <<= 1
		s := suppliers[z]
		o := s.Supply()
		output |= o
	}
	fmt.Println("part1", output)
}

func (d Day24) Part1Recursive(input string) {
	sections := strings.Split(input, "\n\n")
	values := make(map[string]int)
	for _, line := range strings.Split(sections[0], "\n") {
		// e.g., x00: 1
		var key string
		var val int
		_, _ = fmt.Sscanf(line, "%s %d", &key, &val)
		key = key[:len(key)-1] // chop off trailing ':'
		values[key] = val
	}

	connections := make(map[string][3]string)
	var zs []string
	for _, line := range strings.Split(sections[1], "\n") {
		// e.g., y06 AND x06 -> pqj
		var operand1, operand2, op, target string
		_, _ = fmt.Sscanf(line, "%s %s %s -> %s", &operand1, &op, &operand2, &target)
		connections[target] = [3]string{op, operand1, operand2}
		if target[0] == 'z' {
			zs = append(zs, target)
		}
	}

	slices.Sort(zs)
	slices.Reverse(zs)
	output := 0
	for _, z := range zs {
		output <<= 1
		values[z] = d.computeValue(z, values, connections)
		output |= values[z]
	}
	fmt.Println("part1 recursive", output)
}

func (d Day24) computeValue(name string, values map[string]int, connections map[string][3]string) int {
	if value, ok := values[name]; ok {
		return value
	}

	op, op1name, op2name := connections[name][0], connections[name][1], connections[name][2]

	op1value := d.computeValue(op1name, values, connections)
	values[op1name] = op1value
	op2value := d.computeValue(op2name, values, connections)
	values[op2name] = op2value

	switch op {
	case "AND":
		return op1value & op2value
	case "OR":
		return op1value | op2value
	case "XOR":
		return op1value ^ op2value
	default:
		panic(fmt.Sprintf("unknown operation: %s", op))
	}
}

func (d Day24) Part2(input string) {
	sections := strings.Split(input, "\n\n")

	// mapping of [input1, input2] -> [[gate, output]...]
	mapping := make(map[[2]string][][2]string)
	var zs []string
	for _, line := range strings.Split(sections[1], "\n") {
		// e.g., y06 AND x06 -> pqj
		parts := strings.Split(line, " ")
		input1, gate, input2, output := parts[0], parts[1], parts[2], parts[4]
		var key [2]string
		if input1 < input2 {
			key = [2]string{input1, input2}
		} else {
			key = [2]string{input2, input1}
		}
		mapping[key] = append(mapping[key], [2]string{gate, output})

		if output[0] == 'z' {
			zs = append(zs, output)
		}
	}

	slices.Sort(zs)
	var crossed []string
	var adders []Adder
	// lowest output bit uses a half adder
	var adder Adder
	adder = &HalfAdder{a: "x00", b: "y00"}
	adder.Assemble("", mapping)
	adders = append(adders, adder)
	prev := adder
	// the rest of the bits use a full adder except for the highest bit which
	// is just the carry output from the last full adder
	for _, z := range zs[1 : len(zs)-1] {
		a, b := "x"+z[1:], "y"+z[1:]
		adder = &FullAdder{a: a, b: b}
		crossed = append(adder.Assemble(prev.Carry(), mapping), crossed...)
		adders = append(adders, adder)
		prev = adder
	}
	slices.Sort(crossed)
	fmt.Println("part2", strings.Join(crossed, ","))
}

type Adder interface {
	Carry() string
	Assemble(string, map[[2]string][][2]string) []string
}

type FullAdder struct {
	// inputs
	a, b, c string
	// output gates
	sum, carry string
	// intermediate gates
	xo1      string
	an1, an2 string
}

func (f *FullAdder) Assemble(c string, m map[[2]string][][2]string) []string {
	f.c = c

	// ensures keys are sorted
	mk := func(a, b string) [2]string {
		if a < b {
			return [2]string{a, b}
		}
		return [2]string{b, a}
	}

	var crossed []string

	// a full adder consists of 5 gates: 2 XOR, 2 AND, 1 OR
	// gates are identified by their output name
	// starting with the inputs, look for the gates in the mapping

	// a and b are inputs to two gates
	// a XOR b -> f.xo1
	// a AND b -> f.an1
	key := mk(f.a, f.b)
	targets, exists := m[key]
	for _, target := range targets {
		if target[0] == "AND" {
			f.an1 = target[1]
		} else if target[0] == "XOR" {
			f.xo1 = target[1]
		} else {
			fmt.Println(key, targets)
			panic(f)
		}
	}

	// c and xo1 (from above) are inputs to two gates
	// c XOR xo1 -> f.sum
	// c AND xo1 -> f.an2
	key = mk(f.xo1, f.c)
	targets, exists = m[key]
	if !exists {
		// xo1 and an1 outputs crossed
		// xo1 cannot be crossed with a, b or c
		crossed = append(crossed, f.xo1, f.an1)
		key = mk(f.an1, f.c)
		targets, exists = m[key]
		if !exists {
			fmt.Println(key, targets)
			panic(f)
		}
		f.xo1, f.an1 = f.an1, f.xo1
	}
	for _, target := range targets {
		if target[0] == "AND" {
			f.an2 = target[1]
		} else if target[0] == "XOR" {
			f.sum = target[1]
		} else {
			fmt.Println(key, targets)
			panic(f)
		}
	}

	// an1 and an2 are inputs to one gate
	// an1 OR an2 -> f.carry
	key = mk(f.an1, f.an2)
	targets, exists = m[key]
	if !exists {
		// either an1 or an2 is crossed with sum
		// an1 crossed with xo1 would have been caught earlier
		// an2 cannot be crossed with xo1 since xo1 is an input to an2
		key = mk(f.sum, f.an1)
		targets, exists = m[key]
		if exists {
			// sum and an2 outputs are crossed
			crossed = append(crossed, f.sum, f.an2)
			f.sum, f.an2 = f.an2, f.sum
		} else {
			key = mk(f.sum, f.an2)
			targets, exists = m[key]
			if exists {
				// sum and an1 outputs are crossed
				crossed = append(crossed, f.sum, f.an1)
				f.sum, f.an1 = f.an1, f.sum
			} else {
				fmt.Println(key, targets)
				panic(f)
			}
		}
	}
	for _, target := range targets {
		if target[0] == "OR" {
			f.carry = target[1]
		} else {
			fmt.Println(key, targets)
			panic(f)
		}
	}

	if f.sum[0] != 'z' {
		// sum and carry outputs are crossed
		crossed = append(crossed, f.sum, f.carry)
		f.sum, f.carry = f.carry, f.sum
	}

	return crossed
}

func (f *FullAdder) Carry() string {
	return f.carry
}

func (f *FullAdder) String() string {
	return fmt.Sprintf("fulladder(xo1:%3s XOR %3s -> %3s, "+
		"sum:%3s XOR %3s -> %3s, "+
		"an1:%3s AND %3s -> %3s, "+
		"an2:%3s AND %3s -> %3s, "+
		"carry:%3s OR %3s -> %3s)",
		f.a, f.b, f.xo1, f.xo1, f.c, f.sum, f.a, f.b, f.an1, f.xo1, f.c, f.an2, f.an1, f.an2, f.carry)
}

type HalfAdder struct {
	a, b       string
	sum, carry string
}

func (h *HalfAdder) Assemble(_ string, m map[[2]string][][2]string) []string {
	targets := m[[2]string{h.a, h.b}]

	for _, target := range targets {
		if target[0] == "AND" {
			h.carry = target[1]
		} else if target[0] == "XOR" {
			h.sum = target[1]
		}
	}

	var crossed []string
	if h.sum[0] != 'z' {
		// sum and carry outputs are crossed
		crossed = append(crossed, h.sum, h.carry)
		h.sum, h.carry = h.carry, h.sum
	}
	return crossed
}

func (h *HalfAdder) Carry() string {
	return h.carry
}

func (h *HalfAdder) String() string {
	return fmt.Sprintf("halfadder(sum:%3s XOR %3s -> %3s, "+
		"carry:%3s AND %3s -> %3s)",
		h.a, h.b, h.sum, h.a, h.b, h.carry)
}

type Supplier interface {
	Supply() int
	Name() string
	ConnectOutput(Supplier)
}

type Switch struct {
	on       int
	n        string
	outConns map[Supplier]struct{}
}

func NewSwitch(on int, name string) *Switch {
	return &Switch{on: on, n: name, outConns: make(map[Supplier]struct{})}
}

func (s *Switch) ConnectOutput(t Supplier) {
	s.outConns[t] = struct{}{}
}

func (s *Switch) Name() string {
	return s.n
}

func (s *Switch) Supply() int {
	return s.on
}

func (s *Switch) String() string {
	return fmt.Sprintf("switch(n:%s, on:%d, outConns:%v)", s.n, s.on, s.outputNames())
}

func (s *Switch) outputNames() []string {
	var names []string
	for c := range s.outConns {
		names = append(names, c.Name())
	}
	return names
}

type Gate struct {
	t        string
	n        string
	o        int
	inConns  map[string]Supplier
	outConns map[Supplier]struct{}
}

func NewGate(t string, n string, ina string, inb string, o int) *Gate {
	return &Gate{
		t: t,
		n: n,
		o: o,
		inConns: map[string]Supplier{
			ina: nil,
			inb: nil,
		},
		outConns: make(map[Supplier]struct{}),
	}
}

func (g *Gate) MissingConnections() []string {
	var missing []string
	for name, conn := range g.inConns {
		if conn != nil {
			continue
		}
		missing = append(missing, name)
	}
	return missing
}

func (g *Gate) ConnectInput(t Supplier) {
	g.inConns[t.Name()] = t
	t.ConnectOutput(g)
}

func (g *Gate) ConnectOutput(t Supplier) {
	g.outConns[t] = struct{}{}
}

func (g *Gate) Name() string {
	return g.n
}

func (g *Gate) Supply() int {
	if g.o >= 0 {
		return g.o
	}
	if len(g.inConns) != 2 {
		panic(fmt.Errorf("%s is misisng input connections: %v", g.n, g.inConns))
	}
	var inputs []int
	for name, conn := range g.inConns {
		if conn == nil {
			panic(fmt.Errorf("%s has no connection for input: %s", g.n, name))
		}
		inputs = append(inputs, conn.Supply())
	}

	var o int
	switch g.t {
	case "AND":
		o = inputs[0] & inputs[1]
	case "OR":
		o = inputs[0] | inputs[1]
	case "XOR":
		o = inputs[0] ^ inputs[1]
	default:
		panic(fmt.Sprintf("%s has an unknown gate type: %s", g.n, g.t))
	}
	g.o = o
	return o
}

func (g *Gate) String() string {
	return fmt.Sprintf("gate(n:%s, t:%s, o:%d, inConns:%v, outConns:%v)", g.n, g.t, g.o, g.inputNames(), g.outputNames())
}

func (g *Gate) inputNames() []string {
	var names []string
	for name := range g.inConns {
		names = append(names, name)
	}
	return names
}

func (g *Gate) outputNames() []string {
	var names []string
	for s := range g.outConns {
		names = append(names, s.Name())
	}
	return names
}
