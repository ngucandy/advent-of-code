package main

import (
	"bufio"
	"github.com/ngucandy/advent-of-code/internal/helpers"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	bytes, _ := os.ReadFile(infile)
	input := string(bytes)

	part1(input)
	part2(input)
}

func part1(input string) {
	defer helpers.TrackTime(time.Now(), "part1")
	total := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		secret, _ := strconv.Atoi(line)

		for range 2000 {
			secret = nextSecret(secret)
		}
		total += secret
	}
	slog.Info("Part 1:", "total", total)
}

func part2(input string) {
	defer helpers.TrackTime(time.Now(), "part2")
	times := 2000
	buyerPrices := make([][]int, 0)
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		secret, _ := strconv.Atoi(line)

		prices := make([]int, times+1)
		prices[0], _ = strconv.Atoi(line[len(line)-1:])
		for i := 1; i < times+1; i++ {
			next := nextSecret(secret)
			nextStr := strconv.Itoa(next)
			prices[i], _ = strconv.Atoi(nextStr[len(nextStr)-1:])
			secret = next
		}
		buyerPrices = append(buyerPrices, prices)
	}

	start := time.Now()
	buyerIndexes := make([]map[[4]int]int, 0)
	allIndexes := make(map[[4]int]bool)
	for _, prices := range buyerPrices {
		index := make(map[[4]int]int)
		for i := 4; i < len(prices); i++ {
			changes := [4]int{
				prices[i-3] - prices[i-4],
				prices[i-2] - prices[i-3],
				prices[i-1] - prices[i-2],
				prices[i] - prices[i-1],
			}
			if _, ok := index[changes]; ok {
				continue
			}
			index[changes] = prices[i]
			allIndexes[changes] = true
		}
		buyerIndexes = append(buyerIndexes, index)
	}
	helpers.TrackTime(start, "building indexes")

	start = time.Now()
	maxPrice := 0
	var maxIndex []int
	ch := make(chan []int)
	go func() {
		for result := range ch {
			price := result[0]
			if price > maxPrice {
				maxPrice = price
				maxIndex = result[1:]
			}
		}
	}()
	wg := sync.WaitGroup{}
	wg.Add(len(allIndexes))
	for index := range allIndexes {
		go func(idx [4]int) {
			price := 0
			for _, buyer := range buyerIndexes {
				price += buyer[index]
			}
			ch <- append([]int{price}, idx[:]...)
			wg.Done()
		}(index)
	}
	wg.Wait()
	//for index := range allIndexes {
	//	price := 0
	//	for _, buyer := range buyerIndexes {
	//		price += buyer[index]
	//	}
	//	if price > maxPrice {
	//		maxPrice = price
	//		maxIndex = index
	//	}
	//}
	helpers.TrackTime(start, "find max prices")
	slog.Info("Part 2:", "price", maxPrice, "index", maxIndex)
}

func nextSecret(secret int) int {
	next := secret << 6 // multiply by 64
	next ^= secret
	next &= 0xffffff // modulo 16777216

	secret = next
	next = secret >> 5 // divide by 32
	next ^= secret
	next &= 0xffffff // modulo 16777216

	secret = next
	next = secret << 11 // multiply by 2048
	next ^= secret
	next &= 0xffffff // modulo 16777216

	return next
}
