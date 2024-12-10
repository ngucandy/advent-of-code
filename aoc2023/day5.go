package aoc2023

import (
	"slices"
)

type Mapper struct {
	Src      string
	Dst      string
	mappings []Mapping
}

func NewMapper(src string, dst string) *Mapper {
	return &Mapper{Src: src, Dst: dst}
}

func (m *Mapper) Map(src int) (int, string) {
	for _, r := range m.mappings {
		if src >= r.srcStart && src < r.srcEnd {
			return r.dstStart + (src - r.srcStart), m.Dst
		}
	}
	return src, m.Dst
}

func (m *Mapper) MapRange(srcStart int, length int) ([][]int, string) {
	var dstRanges [][]int

	// chop up src range into multiple ranges
	srcEnd := srcStart + length
	currentStart := srcStart
	var srcMappings []Mapping
	slices.SortFunc(m.mappings, func(a, b Mapping) int {
		if a.srcEnd < b.srcEnd {
			return -1
		}
		if a.srcEnd > b.srcEnd {
			return 1
		}
		return 0
	})
	for _, r := range m.mappings {
		if r.srcEnd <= srcStart || r.srcStart >= srcEnd {
			continue
		}
		if r.srcEnd <= srcEnd {
			if currentStart < r.srcStart {
				srcMapping := NewMapping(currentStart, currentStart, r.srcStart-currentStart)
				srcMappings = append(srcMappings, srcMapping)
				currentStart = r.srcStart
			}
			srcMapping := NewMapping(currentStart, r.dstStart+(currentStart-r.srcStart), r.srcEnd-currentStart)
			srcMappings = append(srcMappings, srcMapping)
			currentStart = r.srcEnd
			continue
		}
		if r.srcStart < srcEnd {
			if currentStart < r.srcStart {
				srcMapping := NewMapping(currentStart, currentStart, r.srcStart-currentStart)
				srcMappings = append(srcMappings, srcMapping)
				currentStart = r.srcStart
			}
			srcMapping := NewMapping(currentStart, r.dstStart, srcEnd-currentStart)
			srcMappings = append(srcMappings, srcMapping)
			currentStart = srcEnd
			continue
		}
	}
	// deal with remainder of original source range
	if currentStart < srcEnd {
		srcMapping := NewMapping(currentStart, currentStart, srcEnd-currentStart)
		srcMappings = append(srcMappings, srcMapping)
	}

	for _, srcMapping := range srcMappings {
		dstRange := []int{srcMapping.dstStart, srcMapping.dstEnd - srcMapping.dstStart}
		dstRanges = append(dstRanges, dstRange)
	}

	return dstRanges, m.Dst
}

func (m *Mapper) AddMapping(mapping Mapping) {
	m.mappings = append(m.mappings, mapping)
}

type Mapping struct {
	srcStart int
	srcEnd   int
	dstStart int
	dstEnd   int
}

func NewMapping(srcStart int, dstStart int, length int) Mapping {
	return Mapping{
		srcStart: srcStart,
		srcEnd:   srcStart + length,
		dstStart: dstStart,
		dstEnd:   dstStart + length,
	}
}
