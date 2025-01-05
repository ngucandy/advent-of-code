package queue

import "container/heap"

type Item [10]int

type PriorityQueue []Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i][0] < pq[j][0]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(Item))
}

func Push(pq *PriorityQueue, item Item) {
	heap.Push(pq, item)
}

func (pq *PriorityQueue) Pop() (x any) {
	x, *pq = (*pq)[len(*pq)-1], (*pq)[:len(*pq)-1]
	return x
}

func Pop(pq *PriorityQueue) Item {
	return heap.Pop(pq).(Item)
}

func NewPriorityQueue() *PriorityQueue {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	return &pq
}
