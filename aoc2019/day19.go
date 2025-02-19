package aoc2019

func init() {
	Days["19"] = &Day19{}
}

type Day19 struct {
	eg1, eg2 string
}

func (d Day19) Part1(input string) any {
	program := ParseIntcodeProgram(input)
	count := 0
	n := 50
	for y := range n {
		for x := range n {
			c := NewIntcodeComputer(program, []int{x, y})
			for len(c.output) == 0 {
				c.Step()
			}
			if c.output[0] == 1 {
				count++
			}
		}
	}
	return count
}

func (d Day19) Part2(input string) any {
	program := ParseIntcodeProgram(input)
	ans := 0
	x, y := 0, 50
	for {
		c := NewIntcodeComputer(program, []int{x, y})
		for len(c.output) == 0 {
			c.Step()
		}
		if c.output[0] == 0 {
			x++
			continue
		}
		c = NewIntcodeComputer(program, []int{x + 99, y})
		for len(c.output) == 0 {
			c.Step()
		}
		if c.output[0] == 0 {
			y++
			continue
		}
		c = NewIntcodeComputer(program, []int{x, y + 99})
		for len(c.output) == 0 {
			c.Step()
		}
		if c.output[0] == 0 {
			x++
			continue
		}
		ans = x*10000 + y
		break
	}
	return ans
}
