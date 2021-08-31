package main

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]     // last
	*h = old[0 : n-1] // pop off last
	return x
}

func findKthLargest(nums []int, k int) int {

	hp := &IntHeap{}
	heap.Init(hp)

	for _, num := range nums {
		heap.Push(hp, num)
		// only keep the amount of k
		// in the min heap
		// that way, when you pop off, the top of the heap is the kth last number
		if hp.Len() > k {
			heap.Pop(hp)
		}
	}

	return heap.Pop(hp).(int)
}
