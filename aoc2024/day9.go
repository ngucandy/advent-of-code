package aoc2024

import (
	"fmt"
	"strings"
)

func init() {
	DayMap["9"] = Day9{
		`2333133121414131402`,
	}
}

type Day9 struct {
	example string
}

func (d Day9) Part1(input string) {
	var blocks []int
	for i, ch := range strings.TrimSpace(input) {
		n := int(ch - '0')
		var id int
		if i%2 == 1 { // free block
			id = -1
		} else { // file block
			id = i / 2
		}
		for range n {
			blocks = append(blocks, id)
		}
	}

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
	fmt.Println("part1", sum)
}

func (d Day9) Part2(input string) {
	freeBlocks := [][2]int{}
	fileBlocks := [][2]int{}
	id := 0
	for i, ch := range strings.TrimSpace(input) {
		n := int(ch - '0')
		if i%2 == 1 { // free block
			freeBlocks = append(freeBlocks, [2]int{id, n})
		} else { // file block
			fileBlocks = append(fileBlocks, [2]int{id, n})
		}
		id += n
	}

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
	fmt.Println("part2", sum)
}
