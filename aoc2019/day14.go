package aoc2019

import (
	"fmt"
	"strings"
)

func init() {
	Days["14"] = &Day14{
		eg1: `10 ORE => 10 A
1 ORE => 1 B
7 A, 1 B => 1 C
7 A, 1 C => 1 D
7 A, 1 D => 1 E
7 A, 1 E => 1 FUEL`,
		eg2: `9 ORE => 2 A
8 ORE => 3 B
7 ORE => 5 C
3 A, 4 B => 1 AB
5 B, 7 C => 1 BC
4 C, 1 A => 1 CA
2 AB, 3 BC, 4 CA => 1 FUEL`,
		eg3: `157 ORE => 5 NZVS
165 ORE => 6 DCFZ
44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL
12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ
179 ORE => 7 PSHF
177 ORE => 5 HKGWZ
7 DCFZ, 7 PSHF => 2 XJWVT
165 ORE => 2 GPVTF
3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT`,
		eg4: `2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG
17 NVRVD, 3 JNWZP => 8 VPVL
53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL
22 VJHF, 37 MNCFX => 5 FWMGM
139 ORE => 4 NVRVD
144 ORE => 7 JNWZP
5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC
5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV
145 ORE => 6 MNCFX
1 NVRVD => 8 CXFTF
1 VJHF, 6 MNCFX => 4 RFSQX
176 ORE => 6 VJHF`,
		eg5: `171 ORE => 8 CNZTR
7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL
114 ORE => 4 BHXH
14 VRPVC => 6 BMBT
6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL
6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT
15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW
13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW
5 BMBT => 4 WPTQ
189 ORE => 9 KTJDG
1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP
12 VRPVC, 27 CNZTR => 2 XDBXC
15 KTJDG, 12 BHXH => 5 XCVML
3 BHXH, 2 VRPVC => 7 MZWV
121 ORE => 7 VRPVC
7 XCVML => 6 RJRHP
5 BHXH, 4 VRPVC => 5 LTCX`,
	}
}

type Day14 struct {
	eg1, eg2, eg3, eg4, eg5 string
}

func (d Day14) Part1(input string) any {
	reactions := make(map[string][]string)
	exchanges := make(map[[2]string][2]int)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " => ")

		// output side of reaction
		var nameOut string
		var qtyOut int
		_, _ = fmt.Sscanf(parts[1], "%d %s", &qtyOut, &nameOut)

		// input side of reaction
		for _, rInput := range strings.Split(parts[0], ", ") {
			var nameIn string
			var qtyIn int
			_, _ = fmt.Sscanf(rInput, "%d %s", &qtyIn, &nameIn)
			reactions[nameOut] = append(reactions[nameOut], nameIn)
			exchanges[[2]string{nameOut, nameIn}] = [2]int{qtyOut, qtyIn}
		}
	}

	return d.oreNeeded("FUEL", 1, exchanges, reactions)
}

func (d Day14) oreNeeded(nameStart string, qtyStart int, exchanges map[[2]string][2]int, reactions map[string][]string) int {
	holding := make(map[string]int)
	holding[nameStart] = qtyStart
	q := []string{nameStart}
	for len(q) > 0 {
		out := q[0]
		q = q[1:]

		if out != nameStart && len(q) == 0 {
			// there are no more reactions possible with what we have left
			// so go through what we have left and borrow what we need
			for name, qty := range holding {
				if name == "ORE" {
					continue
				}
				if qty <= 0 {
					continue
				}
				// we still need more of `name` materials so borrow them by
				// ignoring min quantity requirements
				qtyOut := exchanges[[2]string{name, reactions[name][0]}][0]
				for _, in := range reactions[name] {
					qtyIn := exchanges[[2]string{name, in}][1]
					holding[in] += qtyIn
					q = append(q, in)
				}
				holding[name] -= qtyOut
			}
			continue
		}

		if out == "ORE" {
			continue
		}

		qtyOut := exchanges[[2]string{out, reactions[out][0]}][0]
		// if we aren't holding enough of the `out` material for this reaction,
		// then skip it in hopes that future exchanges will provide the amount
		// we're short
		if holding[out] < qtyOut {
			continue
		}

		// see how many multiples of the `out` material we have for this
		// reaction
		multiples := holding[out] / qtyOut
		for _, in := range reactions[out] {
			qtyIn := exchanges[[2]string{out, in}][1]
			holding[in] += qtyIn * multiples
			q = append(q, in)
		}
		holding[out] -= qtyOut * multiples
	}
	return holding["ORE"]
}

func (d Day14) Part2(input string) any {
	reactions := make(map[string][]string)
	exchanges := make(map[[2]string][2]int)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " => ")

		// output side of reaction
		var nameOut string
		var qtyOut int
		_, _ = fmt.Sscanf(parts[1], "%d %s", &qtyOut, &nameOut)

		// input side of reaction
		for _, rInput := range strings.Split(parts[0], ", ") {
			var nameIn string
			var qtyIn int
			_, _ = fmt.Sscanf(rInput, "%d %s", &qtyIn, &nameIn)
			reactions[nameOut] = append(reactions[nameOut], nameIn)
			exchanges[[2]string{nameOut, nameIn}] = [2]int{qtyOut, qtyIn}
		}
	}

	// binary search for ore needed at different amounts of fuel between
	// 1 and 1 trillion
	holding := 1_000_000_000_000 // 1 trillion
	fuel := holding / 2
	for fuelMin, fuelMax := 1, holding; (fuelMax-fuelMin)/2 > 0; {
		ore := d.oreNeeded("FUEL", fuel, exchanges, reactions)
		if ore == holding {
			break
		}
		if ore > holding {
			fuelMax = fuel
		} else {
			fuelMin = fuel
		}
		fuel = fuelMin + ((fuelMax - fuelMin) / 2)
	}
	return fuel
}
