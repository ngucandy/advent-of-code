package queue

import "container/heap"

// Item is the interface we require for items being stored in the queue.
type Item interface {
	Cost() int
}

type store []Item

func (s store) Len() int {
	return len(s)
}

func (s store) Less(i, j int) bool {
	return s[i].Cost() < s[j].Cost()
}

func (s store) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *store) Push(x any) {
	*s = append(*s, x.(Item))
}

func (s *store) Pop() (x any) {
	// highest priority item is at the end
	x = (*s)[len(*s)-1]
	// chop off the last item
	*s = (*s)[:len(*s)-1]
	return x
}

// PQ is a priority queue implementation.
type PQ[E Item] struct {
	q store
}

func (pq *PQ[E]) Push(item E) {
	heap.Push(&pq.q, item)
}

func (pq *PQ[E]) Pop() E {
	return heap.Pop(&pq.q).(E)
}

func (pq *PQ[E]) Len() int {
	return len(pq.q)
}
