package aoc2024

import (
	"reflect"
	"testing"
)

func TestDay21_Part1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"379A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d21 := Days["21"].(*Day21)
			d := Day21{
				numpad:     d21.numpad,
				npButtons:  d21.npButtons,
				dirpad:     d21.dirpad,
				dpButtons:  d21.dpButtons,
				directions: d21.directions,
				cache:      d21.cache,
			}
			d.Part1(tt.name)
		})
	}
}

func BenchmarkDay21_Part1(b *testing.B) {
	tests := []struct {
		name string
	}{
		{"379A"},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			d21 := Days["21"].(*Day21)
			d := Day21{
				numpad:     d21.numpad,
				npButtons:  d21.npButtons,
				dirpad:     d21.dirpad,
				dpButtons:  d21.dpButtons,
				directions: d21.directions,
				cache:      d21.cache,
			}
			d.Part1(tt.name)
		})
	}
}

func TestDay21_paths(t *testing.T) {
	d21 := Days["21"].(*Day21)
	type args struct {
		s       rune
		e       rune
		grid    [][]rune
		buttons map[rune][2]int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"A->2", args{'A', '2', d21.numpad, d21.npButtons}, []string{"<^A", "^<A"}},
		{"A->0", args{'A', '0', d21.numpad, d21.npButtons}, []string{"<A"}},
		{"A->3", args{'A', '3', d21.numpad, d21.npButtons}, []string{"^A"}},
		{"A->7", args{'A', '7', d21.numpad, d21.npButtons}, []string{"<^<^^A", "<^^<^A", "<^^^<A", "^<<^^A", "^<^<^A", "^<^^<A", "^^<<^A", "^^<^<A", "^^^<<A"}},
		{"1->9", args{'1', '9', d21.numpad, d21.npButtons}, []string{">>^^A", ">^>^A", ">^^>A", "^>>^A", "^>^>A", "^^>>A"}},
		{"A->^", args{'A', '^', d21.dirpad, d21.dpButtons}, []string{"<A"}},
		{"A->>", args{'A', '>', d21.dirpad, d21.dpButtons}, []string{"vA"}},
		{"A->v", args{'A', 'v', d21.dirpad, d21.dpButtons}, []string{"<vA", "v<A"}},
		{"A-><", args{'A', '<', d21.dirpad, d21.dpButtons}, []string{"<v<A", "v<<A"}},
		{"<->A", args{'<', 'A', d21.dirpad, d21.dpButtons}, []string{">>^A", ">^>A"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Day21{
				directions: d21.directions,
				cache:      d21.cache,
			}
			got := d.paths(tt.args.s, tt.args.e, tt.args.grid, tt.args.buttons)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("paths() = %v, want %v", got, tt.want)
			}
			if cache := d.cache[[2]rune{tt.args.s, tt.args.e}]; !reflect.DeepEqual(got, cache) {
				t.Errorf("paths() = %v, cache %v", got, cache)
			}
		})
	}
}
