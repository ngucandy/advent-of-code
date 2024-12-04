package main

import (
	"bufio"
	"log/slog"
	"math"
	"os"
	"regexp"
	"strconv"

	"github.com/ngucandy/advent-of-code/aoc2023"
)

var (
	rexpDigits = regexp.MustCompile(`\d+`)
	rexpSeeds  = regexp.MustCompile(`^seeds:`)
	rexpMap    = regexp.MustCompile(` map:`)
	rexpMapper = regexp.MustCompile(`(.*)-to-(.*) `)
)

func main() {
	infile := os.Args[1]
	slog.Info("Reading input file:", "name", infile)
	file, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			slog.Error("Error closing file:", "error", err)
		}
	}(file)

	var seeds []int
	var mapper *aoc2023.Mapper
	mappers := make(map[string]*aoc2023.Mapper)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		if rexpSeeds.MatchString(line) {
			nums := rexpDigits.FindAllString(line, -1)
			for _, num := range nums {
				n, err := strconv.Atoi(num)
				if err != nil {
					panic(err)
				}
				seeds = append(seeds, n)
			}
			slog.Info("Seeds:", "seeds", seeds)
			continue
		}

		if rexpMap.MatchString(line) {
			matches := rexpMapper.FindStringSubmatch(line)
			mapper = aoc2023.NewMapper(matches[1], matches[2])
			mappers[mapper.Src] = mapper
			slog.Info("Mapper:", "mapper", mapper)
			continue
		}

		nums := rexpDigits.FindAllString(line, -1)
		dstStart, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		srcStart, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		length, err := strconv.Atoi(nums[2])
		if err != nil {
			panic(err)
		}
		if mapper == nil {
			panic("mapper not initialized")
		}
		mapping := aoc2023.NewMapping(srcStart, dstStart, length)
		mapper.AddMapping(mapping)
	}

	minDstVal := math.MaxInt
	for _, seed := range seeds {
		dstVal, dst := deepMap(seed, "seed", mappers)
		minDstVal = min(dstVal, minDstVal)
		slog.Info("Mapped:", "src", "seed", "srcVal", seed, "dst", dst, "dstVal", dstVal)
	}
	slog.Info("Part 1:", "smallest location", minDstVal)
}

func deepMap(srcVal int, src string, mappers map[string]*aoc2023.Mapper) (int, string) {
	if mapper, ok := mappers[src]; ok {
		nextVal, nextSrc := mapper.Map(srcVal)
		return deepMap(nextVal, nextSrc, mappers)
	}
	return srcVal, src
}
