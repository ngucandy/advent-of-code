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

	freeBlocks := [][2]int{}
	fileBlocks := [][2]int{}
	blocks := []int{}
	ii := 0
	for i, r := range input {
		n := int(r - '0')
		var id int
		if i%2 == 1 { // free block
			id = -1
			freeBlocks = append(freeBlocks, [2]int{ii, n})
		} else { // file block
			id = i / 2
			fileBlocks = append(fileBlocks, [2]int{ii, n})
		}
		for range n {
			blocks = append(blocks, id)
			ii++
		}
	}

	part1(blocks)
	part2(fileBlocks, freeBlocks)
}

func part2(fileBlocks [][2]int, freeBlocks [][2]int) {
	sum := 0

	for id := len(fileBlocks) - 1; id > 0; id-- {
		// find free block big enough to hold file
		for i, freeBlock := range freeBlocks {
			if freeBlock[0] > fileBlocks[id][0] { // free block isn't left of file
				break
			}
			if freeBlock[1] < fileBlocks[id][1] { // too small
				continue
			}
			if freeBlock[1] == fileBlocks[id][1] { // exact fit
				fileBlocks[id] = freeBlock
				freeBlocks = append(freeBlocks[:i], freeBlocks[i+1:]...)
				break
			}
			// free block is bigger than file
			fileBlocks[id][0] = freeBlock[0]
			freeBlocks[i][0] += fileBlocks[id][1]
			freeBlocks[i][1] -= fileBlocks[id][1]
		}
	}

	for id, fileBlock := range fileBlocks {
		for i := range fileBlock[1] {
			sum += id * (fileBlock[0] + i)
		}
	}
	slog.Info("Part 2:", "sum", sum)
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
		blocks[i] = blocks[j]
		blocks[j] = -1
	}

	for i, id := range blocks {
		if id < 0 {
			continue
		}
		sum += i * id
	}
	slog.Info("Part 1:", "sum", sum)
}
