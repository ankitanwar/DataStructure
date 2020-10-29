package heaps

import (
	"container/heap"
)

//KthLargest : It will return the kth largest element
func KthLargest(arr []int, k int) interface{} {
	h := &MinHeap{}
	heap.Init(h)
	for i := 0; i < len(arr); i++ {
		heap.Push(h, arr[i])
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return heap.Pop(h)
}
