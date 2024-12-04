package aoc2023

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
		if src >= r.srcStart && src < r.srcStart+r.length {
			return r.dstStart + (src - r.srcStart), m.Dst
		}
	}
	return src, m.Dst
}

func (m *Mapper) AddMapping(mapping Mapping) {
	m.mappings = append(m.mappings, mapping)
}

type Mapping struct {
	srcStart int
	dstStart int
	length   int
}

func NewMapping(srcStart int, dstStart int, length int) Mapping {
	return Mapping{srcStart: srcStart, dstStart: dstStart, length: length}
}
