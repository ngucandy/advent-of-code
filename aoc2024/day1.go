package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	infile := os.Args[1]
	file, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	list1, list2 := parseLists(file)

	printTotalDistance(list1, list2)
	printSimilarityScore(list1, list2)
}

func printSimilarityScore(list1 []int, list2 []int) {
	list2index := make([]int, list2[len(list2)-1]+1, list2[len(list2)-1]+1)
	for _, n := range list2 {
		list2index[n]++
	}

	score := 0
	for _, n := range list1 {
		score += n * list2index[n]
	}
	fmt.Println(score)
}

func printTotalDistance(list1 []int, list2 []int) {
	sum := 0
	for i, n := range list1 {
		sum += absInt(n - list2[i])
	}

	fmt.Println(sum)
}

func parseLists(file *os.File) ([]int, []int) {
	delimre := regexp.MustCompile("\\s+")
	var list1, list2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := delimre.Split(line, -1)
		if n, err := strconv.ParseInt(tokens[0], 10, 32); err == nil {
			list1 = append(list1, int(n))
		}
		if n, err := strconv.ParseInt(tokens[1], 10, 32); err == nil {
			list2 = append(list2, int(n))
		}
	}
	sort.Ints(list1)
	sort.Ints(list2)
	return list1, list2
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
