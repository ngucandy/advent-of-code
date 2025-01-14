package aoc2019

import "testing"

func TestDay10_sort(t *testing.T) {
	r, c := 2, 2
	type fields struct {
		eg1 string
		eg2 string
		eg3 string
		eg4 string
		eg5 string
		eg6 string
	}
	type args struct {
		a [4]int
		b [4]int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "a up, b up, a<b",
			fields: fields{},
			args: args{
				a: [4]int{r - 1, c, -1, 0},
				b: [4]int{r - 2, c, -1, 0},
			},
			want: -1,
		},
		{
			name:   "a up, b up, b<a",
			fields: fields{},
			args: args{
				a: [4]int{r - 2, c, -1, 0},
				b: [4]int{r - 1, c, -1, 0},
			},
			want: 1,
		},
		{
			name:   "a down, b down, a<b",
			fields: fields{},
			args: args{
				a: [4]int{r + 1, c, 1, 0},
				b: [4]int{r + 2, c, 1, 0},
			},
			want: -1,
		},
		{
			name:   "a down, b down, b<a",
			fields: fields{},
			args: args{
				a: [4]int{r + 2, c, 1, 0},
				b: [4]int{r + 1, c, 1, 0},
			},
			want: 1,
		},
		{
			name:   "a up, b down",
			fields: fields{},
			args: args{
				a: [4]int{r - 1, c, -1, 0},
				b: [4]int{r + 1, c, 1, 0},
			},
			want: -1,
		},
		{
			name:   "a down, b up",
			fields: fields{},
			args: args{
				a: [4]int{r + 1, c, 1, 0},
				b: [4]int{r - 1, c, -1, 0},
			},
			want: 1,
		},
		{
			name:   "a down, b right",
			fields: fields{},
			args: args{
				a: [4]int{r + 1, c, 1, 0},
				b: [4]int{r, c + 1, 0, 1},
			},
			want: 1,
		},
		{
			name:   "a down, b left",
			fields: fields{},
			args: args{
				a: [4]int{r + 1, c, 1, 0},
				b: [4]int{r, c - 1, 0, -1},
			},
			want: -1,
		},
		{
			name:   "a right, b down",
			fields: fields{},
			args: args{
				a: [4]int{r, c + 1, 0, 1},
				b: [4]int{r - 1, c, 1, 0},
			},
			want: -1,
		},
		{
			name:   "a left, b down",
			fields: fields{},
			args: args{
				a: [4]int{r, c - 1, 0, -1},
				b: [4]int{r + 1, c, 1, 0},
			},
			want: 1,
		},
		{
			name:   "a left, b right",
			fields: fields{},
			args: args{
				a: [4]int{r, c - 1, 0, -1},
				b: [4]int{r, c + 1, 0, 1},
			},
			want: 1,
		},
		{
			name:   "a right, b left",
			fields: fields{},
			args: args{
				a: [4]int{r, c + 1, 0, 1},
				b: [4]int{r, c - 1, 0, -1},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Day10{
				eg1: tt.fields.eg1,
				eg2: tt.fields.eg2,
				eg3: tt.fields.eg3,
				eg4: tt.fields.eg4,
				eg5: tt.fields.eg5,
				eg6: tt.fields.eg6,
			}
			if got := d.sort(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
