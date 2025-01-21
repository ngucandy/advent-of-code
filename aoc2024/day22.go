package aoc2024

import (
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ngucandy/advent-of-code/internal/helpers"
)

func init() {
	Days["22"] = Day22{}
}

type Day22 struct {
	eg1, eg2 string
}

func (d Day22) Part1(input string) any {
	defer helpers.TrackTime(time.Now())
	total := 0
	for _, line := range strings.Split(input, "\n") {
		secret, _ := strconv.Atoi(line)

		for range 2000 {
			secret = d.nextSecret(secret)
		}
		total += secret
	}
	return total
}

func (d Day22) Part2(input string) any {
	defer helpers.TrackTime(time.Now())
	times := 2000
	buyerPrices := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		secret, _ := strconv.Atoi(line)

		prices := make([]int, times+1)
		prices[0], _ = strconv.Atoi(line[len(line)-1:])
		for i := 1; i < times+1; i++ {
			next := d.nextSecret(secret)
			nextStr := strconv.Itoa(next)
			prices[i], _ = strconv.Atoi(nextStr[len(nextStr)-1:])
			secret = next
		}
		buyerPrices = append(buyerPrices, prices)
	}

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

	maxPrice := 0
	ch := make(chan []int)
	go func() {
		for result := range ch {
			price := result[0]
			if price > maxPrice {
				maxPrice = price
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
	return maxPrice
}

func (d Day22) nextSecret(secret int) int {
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
