package aoc2025

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func init() {
	Days["2"] = Day2{
		example: "11-22,95-115,998-1012,1188511880-1188511890,222220-222224," +
			"1698522-1698528,446443-446449,38593856-38593862,565653-565659," +
			"824824821-824824827,2121212118-2121212124",
	}
}

type Day2 struct {
	example string
}

func (d Day2) Part1(input string) any {
	sum := 0
	for _, r := range strings.Split(input, ",") {
		parts := strings.Split(r, "-")
		// if number of digits in start and end range are the same and it's
		// odd then this range cannot contain a pattern that repeats twice
		if len(parts[0]) == len(parts[1]) && len(parts[0])%2 == 1 {
			continue
		}
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		for i := start; i <= end; i++ {
			if d.isInvalid1(i) {
				sum += i
			}
		}
	}
	return sum
}

func (d Day2) isInvalid1(n int) bool {
	digits := d.countDigits(n)
	if digits%2 == 1 {
		return false
	}
	// find the power of 10 that will isolate left/right halves of digits of n
	divisor := int(math.Pow(10, float64(digits/2)))
	return n/divisor == n%divisor
}

func (d Day2) countDigits(n int) int {
	x := 10
	var digits int
	for digits = 1; n/x > 0; digits++ {
		x *= 10
	}
	return digits
}

func (d Day2) Part2(input string) any {
	sum := 0
	for _, r := range strings.Split(input, ",") {
		parts := strings.Split(r, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		for i := start; i <= end; i++ {
			if d.isInvalid2(i) {
				sum += i
			}
		}
	}
	return sum
}

func (d Day2) isInvalid2(n int) bool {
	num := fmt.Sprintf("%d", n)
	// iterate through possible patterns of num, starting with the first digit,
	// then first and second, then first, second and third, etc...
	for i, length := 1, len(num); i <= length/2; i++ {
		// skip pattern lengths that are not a multiple of the total length
		if length%i > 0 {
			continue
		}
		if strings.Repeat(num[:i], length/i) == num {
			return true
		}
	}
	return false
}
