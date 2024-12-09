package aoc2024

import (
	"reflect"
	"testing"
)

func TestCartesianProduct(t *testing.T) {
	type args[T any] struct {
		sets [][]T
	}
	type testCase[T any] struct {
		name       string
		args       args[T]
		wantResult [][]T
	}
	tests := []testCase[string]{
		{
			name: "empty",
			args: args[string]{
				sets: [][]string{},
			},
			wantResult: [][]string{{}},
		},
		{
			name: "one set",
			args: args[string]{
				sets: [][]string{{"a", "b"}},
			},
			wantResult: [][]string{{"a"}, {"b"}},
		},
		{
			name: "two sets",
			args: args[string]{
				sets: [][]string{{"a", "b"}, {"a", "b"}},
			},
			wantResult: [][]string{{"a", "a"}, {"a", "b"}, {"b", "a"}, {"b", "b"}},
		},
		{
			name: "3 sets",
			args: args[string]{
				sets: [][]string{{"a", "b"}, {"a", "b"}, {"a", "b"}},
			},
			wantResult: [][]string{
				{"a", "a", "a"},
				{"a", "a", "b"},
				{"a", "b", "a"},
				{"a", "b", "b"},
				{"b", "a", "a"},
				{"b", "a", "b"},
				{"b", "b", "a"},
				{"b", "b", "b"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := CartesianProduct(tt.args.sets); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("CartesianProduct() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestCartesianProductN(t *testing.T) {
	type args[T any] struct {
		set []T
		n   int
	}
	type testCase[T any] struct {
		name       string
		args       args[T]
		wantResult [][]T
	}
	tests := []testCase[string]{
		{
			name: "n=0",
			args: args[string]{
				set: []string{"a", "b"},
				n:   0,
			},
			wantResult: [][]string{{}},
		},
		{
			name: "n=1",
			args: args[string]{
				set: []string{"a", "b"},
				n:   1,
			},
			wantResult: [][]string{
				{"a"},
				{"b"},
			},
		},
		{
			name: "n=2",
			args: args[string]{
				set: []string{"a", "b"},
				n:   2,
			},
			wantResult: [][]string{
				{"a", "a"},
				{"a", "b"},
				{"b", "a"},
				{"b", "b"},
			},
		},
		{
			name: "n=3",
			args: args[string]{
				set: []string{"a", "b"},
				n:   3,
			},
			wantResult: [][]string{
				{"a", "a", "a"},
				{"a", "a", "b"},
				{"a", "b", "a"},
				{"a", "b", "b"},
				{"b", "a", "a"},
				{"b", "a", "b"},
				{"b", "b", "a"},
				{"b", "b", "b"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := CartesianProductN(tt.args.set, tt.args.n); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("CartesianProductN() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
