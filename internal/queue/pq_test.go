package queue

import (
	"testing"
)

type TestItem struct {
	n int
}

func (t TestItem) Cost() int {
	return t.n
}

func TestPQ_PushPop(t *testing.T) {
	type args struct {
		items []TestItem
	}
	tests := []struct {
		name string
		args args
		want TestItem
	}{
		{"push1",
			args{
				[]TestItem{{1}},
			},
			TestItem{1},
		},
		{"push2",
			args{
				[]TestItem{{2}, {1}},
			},
			TestItem{1},
		},
		{"push3",
			args{
				[]TestItem{{100}, {1}, {2}},
			},
			TestItem{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := &PQ{}
			for _, item := range tt.args.items {
				pq.Push(item)
			}
			item := pq.Pop()
			if item != tt.want {
				t.Errorf("item= %v, want %v", item, tt.want)
			}
		})
	}
}
