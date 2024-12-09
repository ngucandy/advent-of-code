package main

import (
	"bufio"
	"log/slog"
	"os"
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, _ := os.Open(infile)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var input string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = scanner.Text()
	}

	blocks := []int{}
	for i, r := range input {
		n := int(r - '0')
		id := i / 2
		if i%2 == 1 {
			id = -1
		}
		for range n {
			blocks = append(blocks, id)
		}
	}

	slog.Info("Mapped blocks:", "input", input, "blocks", blocks)

	part1(blocks)
}

func part1(blocks []int) {
	sum := 0

	for i, j := 0, len(blocks)-1; i < j-1; j-- {
		if blocks[j] < 0 {
			continue
		}
		for _, id := range blocks[i:] {
			if id < 0 {
				break
			}
			i++
		}
		if i >= j {
			break
		}
		//slog.Info("Moving block:", "id", blocks[j], "i", i, "j", j)
		blocks[i] = blocks[j]
		blocks[j] = -1
	}
	//slog.Info("Finished moving:", "blocks", blocks)

	for i, id := range blocks {
		if id < 0 {
			continue
		}
		sum += i * id
	}
	slog.Info("Part 1:", "sum", sum)
}
