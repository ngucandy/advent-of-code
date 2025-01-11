package main

import (
	"fmt"
	"github.com/ngucandy/advent-of-code/internal/helpers"
	"log/slog"
	"os"
	"slices"
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
	part2(input)
}

func part1(input string) {
	modules := buildModules(input)
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

func buildModules(input string) map[string]*Module {
	modules := make(map[string]*Module)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " -> ")

		// chop off '%' or '&' for name
		name := parts[0]
		if slices.Contains([]rune{'%', '&'}, rune(name[0])) {
			name = name[1:]
		}
		module, exists := modules[name]
		if !exists {
			module = &Module{
				Name: name,
			}
			modules[name] = module
		}

		// module type
		switch parts[0][0] {
		case 'b':
			module.Type = "broadcaster"
		case '%':
			module.Type = "%"
			module.OnOff = off
			module.ReceivedPulses = make(map[string]string)
		case '&':
			module.Type = "&"
			module.ReceivedPulses = make(map[string]string)
		}

		// destinations
		module.Destinations = strings.Split(parts[1], ", ")
		for _, dest := range module.Destinations {
			if _, exists := modules[dest]; !exists {
				modules[dest] = &Module{
					Name:           dest,
					Type:           "untyped",
					ReceivedPulses: make(map[string]string),
				}
			}
		}
	}

	// map conjunction inputs
	for _, module := range modules {
		for _, dest := range module.Destinations {
			if m, exists := modules[dest]; exists {
				m.ReceivedPulses[module.Name] = low
			}
		}
	}
	return modules
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
	modules := buildModules(input)
	rx := modules["rx"]
	if len(rx.ReceivedPulses) > 1 {
		panic("rx module has more than one sender")
	}
	var rxSource string
	for k := range rx.ReceivedPulses {
		rxSource = k
	}

	var sources []string
	for feeder := range modules[rxSource].ReceivedPulses {
		sources = append(sources, feeder)
	}

	var q []Pulse
	srcHigh := make(map[string][]int)
	i := 0
Outer:
	for {
		i++
		q = append(q, Pulse{"button", low, "broadcaster"})
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			if slices.Contains(sources, p.Source) && p.Signal == high {
				srcHigh[p.Source] = append(srcHigh[p.Source], i)
				if len(srcHigh[p.Source]) == 10 {
					break Outer
				}
			}
			q = append(q, pulse(p, modules)...)
		}
	}
	lcm := 1
	for src, presses := range srcHigh {
		for j := 1; j < len(presses); j++ {
			if presses[j] != presses[0]*(j+1) {
				panic(fmt.Sprintf("bad repeating pattern; expected %d; got %d: %v", presses[0]*(j+1), presses[j], src))
			}
		}
		lcm = helpers.LCM(lcm, presses[0])
	}
	slog.Info("Part 2:", "presses", lcm)
}
