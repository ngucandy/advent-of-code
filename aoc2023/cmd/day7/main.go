package main

import (
	"bufio"
	"cmp"
	"log/slog"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	cardRanks = map[rune]int{
		'2': 0,
		'3': 1,
		'4': 2,
		'5': 3,
		'6': 4,
		'7': 5,
		'8': 6,
		'9': 7,
		'T': 8,
		'J': 9,
		'Q': 10,
		'K': 11,
		'A': 12,
	}

	cardRanksJoker = map[rune]int{
		'J': 0,
		'2': 1,
		'3': 2,
		'4': 3,
		'5': 4,
		'6': 5,
		'7': 6,
		'8': 7,
		'9': 8,
		'T': 9,
		'Q': 10,
		'K': 11,
		'A': 12,
	}
)

type hand struct {
	cards           string
	classification  int
	classificationJ int
}

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, _ := os.Open(infile)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	bids := make(map[string]int)
	hands := make([]hand, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, " ")
		bid, _ := strconv.Atoi(parts[1])
		bids[parts[0]] = bid

		h := hand{
			cards:           parts[0],
			classification:  classify(parts[0]),
			classificationJ: classifyWithJoker(parts[0]),
		}
		hands = append(hands, h)
	}

	part1(hands, bids)
	part2(hands, bids)

}

func part1(hands []hand, bids map[string]int) {
	total := 0
	slices.SortFunc(hands, cmpHands)
	for i, h := range hands {
		rank := i + 1
		total += rank * bids[h.cards]
	}
	slog.Info("Part 1:", "total", total)
}

func cmpHands(hand1 hand, hand2 hand) int {
	if hand1.classification < hand2.classification {
		return -1
	}
	if hand1.classification > hand2.classification {
		return 1
	}

	cards1 := []rune(hand1.cards)
	cards2 := []rune(hand2.cards)
	for i := range len(cards1) {
		if cards1[i] == cards2[i] {
			continue
		}
		return cmp.Compare(cardRanks[cards1[i]], cardRanks[cards2[i]])
	}
	panic("both hand are the same: " + hand1.cards + "; " + hand2.cards)
}

func classify(cards string) int {
	counts := make(map[rune]int)
	for _, card := range cards {
		counts[card]++
	}
	switch len(counts) {
	case 5: // high card
		return 0
	case 1: // five of a kind
		return 6
	case 4: // one pair
		return 1
	case 2: // four of a kind or full house
		for _, count := range counts {
			if count == 4 {
				return 5
			} // 4 of a kind
		}
		return 4 // full house
	case 3: // two pairs or three of a kind
		for _, count := range counts {
			if count == 3 {
				return 3
			} // 3 of a kind
		}
		return 2 // 2 pair
	default:
		panic("can't classify: " + cards)
	}
}

func part2(hands []hand, bids map[string]int) {
	total := 0
	slices.SortFunc(hands, cmpHandsJ)
	for i, h := range hands {
		rank := i + 1
		total += rank * bids[h.cards]
	}
	slog.Info("Part 2:", "total", total)

}

func cmpHandsJ(hand1 hand, hand2 hand) int {
	if hand1.classificationJ < hand2.classificationJ {
		return -1
	}
	if hand1.classificationJ > hand2.classificationJ {
		return 1
	}

	cards1 := []rune(hand1.cards)
	cards2 := []rune(hand2.cards)
	for i := range len(cards1) {
		if cards1[i] == cards2[i] {
			continue
		}
		return cmp.Compare(cardRanksJoker[cards1[i]], cardRanksJoker[cards2[i]])
	}
	panic("both hand are the same: " + hand1.cards + "; " + hand2.cards)
}

func classifyWithJoker(cards string) int {
	counts := make(map[rune]int)
	for _, card := range cards {
		counts[card]++
	}

	classification := classify(cards)

	switch counts['J'] {
	case 0:
		return classification
	case 1:
		switch classification {
		case 0: // joker turns high card into pair
			return 1
		case 1: // joker turns 1 pair into 3 of a kind
			return 3
		case 2: // joker turns 2 pairs into a full house
			return 4
		case 3: // joker turns 3 of a kind into 4 of a kind
			return 5
		case 5: // joker turns 4 of a kind into 5 of a kind
			return 6
		}
	case 2:
		switch classification {
		case 1: // pair of jokers turns into 3 of a kind
			return 3
		case 2: // pair of jokers and second pair turns into 4 of a kind
			return 5
		case 4: // pair of jokers and 3 of a kind turns into 5 of a kind
			return 6
		}
	case 3:
		switch classification {
		case 3: // 3 jokers turn into 4 of a kind
			return 5
		case 4: // full house with 3 jokers
			return 6
		}
	case 4:
		switch classification {
		case 5: // 4 jokers
			return 6
		}
	case 5:
		return 6
	}
	panic("can't classify: " + cards)
}
