package main

import (
	"log/slog"
	"os"
	"strings"
)

const (
	testInput1 = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

	testInput2 = `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`

	high = "high"
	low  = "low"
	on   = "on"
	off  = "off"
)

type Module struct {
	Name           string
	Type           string
	Destinations   []string
	OnOff          string
	ReceivedPulses map[string]string
}

type Pulse struct {
	Source      string
	Signal      string
	Destination string
}

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	bytes, _ := os.ReadFile(infile)
	input := strings.ReplaceAll(string(bytes), "\r\n", "\n")

	part1(testInput1)
	part1(testInput2)
	part1(input)
	part2(testInput1)
	part2(testInput2)
	part2(input)
}

func part1(input string) {
	modules := make(map[string]*Module)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " -> ")
		switch parts[0][0] {
		case 'b':
			modules[parts[0]] = &Module{
				Name:         "broadcaster",
				Type:         "broadcaster",
				Destinations: strings.Split(parts[1], ", "),
			}
		case '%':
			modules[parts[0][1:]] = &Module{
				Name:         parts[0][1:],
				Type:         "%",
				Destinations: strings.Split(parts[1], ", "),
				OnOff:        off,
			}
		case '&':
			modules[parts[0][1:]] = &Module{
				Name:           parts[0][1:],
				Type:           "&",
				Destinations:   strings.Split(parts[1], ", "),
				ReceivedPulses: make(map[string]string),
			}
		}
	}
	// map conjunction inputs
	for _, module := range modules {
		for _, dest := range module.Destinations {
			if m, exists := modules[dest]; exists && m.Type == "&" {
				m.ReceivedPulses[module.Name] = low
			}
		}
	}
	lc, hc := 0, 0
	var q []Pulse
	for range 1000 {
		q = append(q, Pulse{"button", low, "broadcaster"})
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			if p.Signal == high {
				hc++
			} else {
				lc++
			}
			q = append(q, pulse(p, modules)...)
		}
	}
	slog.Info("Part 1:", "low", lc, "high", hc, "product", lc*hc)
}

func pulse(p Pulse, modules map[string]*Module) []Pulse {
	m, ok := modules[p.Destination]
	if !ok { // untyped module
		return nil
	}
	var next []Pulse
	switch m.Type {
	case "broadcaster":
		for _, nextDest := range m.Destinations {
			next = append(next, Pulse{m.Name, low, nextDest})
		}
	// Flip-flop modules (prefix %) are either on or off; they are initially off.
	// If a flip-flop module receives a high pulse, it is ignored and nothing happens.
	// However, if a flip-flop module receives a low pulse, it flips between on and off.
	// If it was off, it turns on and sends a high pulse. If it was on, it turns off
	// and sends a low pulse.
	case "%":
		if p.Signal == high {
			break
		}
		var nextSignal string
		if m.OnOff == on {
			nextSignal = low
			m.OnOff = off
		} else {
			nextSignal = high
			m.OnOff = on
		}
		for _, nextDest := range m.Destinations {
			next = append(next, Pulse{m.Name, nextSignal, nextDest})
		}
	// Conjunction modules (prefix &) remember the type of the most recent pulse received
	// from each of their connected input modules; they initially default to remembering
	// a low pulse for each input. When a pulse is received, the conjunction module first
	// updates its memory for that input. Then, if it remembers high pulses for all inputs,
	// it sends a low pulse; otherwise, it sends a high pulse.
	case "&":
		m.ReceivedPulses[p.Source] = p.Signal
		nextSignal := low
		for _, signal := range m.ReceivedPulses {
			if signal == low {
				nextSignal = high
				break
			}
		}
		for _, nextDest := range m.Destinations {
			next = append(next, Pulse{m.Name, nextSignal, nextDest})
		}
	}
	return next
}

func part2(input string) {

}
