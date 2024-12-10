package aoc2023

import (
	"reflect"
	"testing"
)

func TestMapping_Map(t *testing.T) {
	type args struct {
		src int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "No mapping",
			args: args{src: 1},
			want: 1,
		},
		{
			name: "Mapper 10",
			args: args{src: 10},
			want: 20,
		},
		{
			name: "Mapper 20",
			args: args{src: 20},
			want: 10,
		},
		{
			name: "No mapping",
			args: args{src: 30},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mapper{}
			m.AddMapping(NewMapping(10, 20, 10))
			m.AddMapping(NewMapping(20, 10, 10))
			if got, _ := m.Map(tt.args.src); got != tt.want {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapper_MapRange(t *testing.T) {
	type fields struct {
		Src      string
		Dst      string
		mappings []Mapping
	}
	type args struct {
		srcStart int
		length   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   [][]int
		want1  string
	}{
		{
			name: "MapRange no overlap",
			fields: fields{
				Src: "src",
				Dst: "dst",
				mappings: []Mapping{
					NewMapping(10, 20, 10),
					NewMapping(30, 40, 10),
				},
			},
			args: args{
				srcStart: 20,
				length:   10,
			},
			want: [][]int{
				{20, 10},
			},
			want1: "dst",
		},
		{
			name: "MapRange full overlap",
			fields: fields{
				Src: "src",
				Dst: "dst",
				mappings: []Mapping{
					NewMapping(10, 20, 10),
				},
			},
			args: args{
				srcStart: 10,
				length:   10,
			},
			want: [][]int{
				{20, 10},
			},
			want1: "dst",
		},
		{
			name: "MapRange full overlap with multiple mappings",
			fields: fields{
				Src: "src",
				Dst: "dst",
				mappings: []Mapping{
					NewMapping(10, 20, 5),
					NewMapping(15, 30, 5),
				},
			},
			args: args{
				srcStart: 10,
				length:   10,
			},
			want: [][]int{
				{20, 5},
				{30, 5},
			},
			want1: "dst",
		},
		{
			name: "MapRange overlap with disjoint mappings",
			fields: fields{
				Src: "src",
				Dst: "dst",
				mappings: []Mapping{
					NewMapping(10, 110, 10),
					NewMapping(30, 130, 10),
				},
			},
			args: args{
				srcStart: 15,
				length:   20,
			},
			want: [][]int{
				{115, 5},
				{20, 10},
				{130, 5},
			},
			want1: "dst",
		},
		{
			name: "MapRange overlap with embedded disjoint mappings",
			fields: fields{
				Src: "src",
				Dst: "dst",
				mappings: []Mapping{
					NewMapping(10, 110, 10),
					NewMapping(23, 123, 5),
					NewMapping(30, 130, 10),
				},
			},
			args: args{
				srcStart: 15,
				length:   20,
			},
			want: [][]int{
				{115, 5},
				{20, 3},
				{123, 5},
				{28, 2},
				{130, 5},
			},
			want1: "dst",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Mapper{
				Src:      tt.fields.Src,
				Dst:      tt.fields.Dst,
				mappings: tt.fields.mappings,
			}
			got, got1 := m.MapRange(tt.args.srcStart, tt.args.length)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapRange() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MapRange() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
