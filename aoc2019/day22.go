package aoc2019

import (
	"fmt"
	"slices"
	"strings"
)

func init() {
	Days["22"] = &Day22{
		eg1: `deal with increment 7
deal into new stack
deal into new stack`,
		eg2: `cut 6
deal with increment 7
deal into new stack`,
		eg3: `deal with increment 3`,
		eg4: `deal with increment 7
deal with increment 9
cut -2`,
		eg5: `deal into new stack
cut -2
deal with increment 7
cut 8
cut -4
deal with increment 7
cut 3
deal with increment 9
deal with increment 3
cut -1`,
	}
}

type Day22 struct {
	eg1, eg2, eg3, eg4, eg5 string
}

func (d Day22) Part1(input string) any {
	//input = d.eg5
	//size := 10
	size := 10007
	var deck []int
	for i := range size {
		deck = append(deck, i)
	}
	for _, line := range strings.Split(input, "\n") {
		if line == "deal into new stack" {
			slices.Reverse(deck)
			continue
		}
		if strings.HasPrefix(line, "cut") {
			var n int
			_, _ = fmt.Sscanf(line, "cut %d", &n)
			var newDeck []int
			if n < 0 {
				newDeck = append(deck[len(deck)+n:], deck[:len(deck)+n]...)
			} else {
				newDeck = append(deck[n:], deck[:n]...)
			}
			deck = newDeck
			continue
		}
		var n int
		_, _ = fmt.Sscanf(line, "deal with increment %d", &n)
		newDeck := make([]int, len(deck))
		for i, j := 0, 0; i < len(deck); i, j = i+1, (j+n)%len(deck) {
			newDeck[j] = deck[i]
		}
		deck = newDeck
	}
	return slices.Index(deck, 2019)
}

func (d Day22) Part2(input string) any {
	return "no answer yet"
}
