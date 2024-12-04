package aoc2023

import "testing"

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
