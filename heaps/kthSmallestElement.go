package heaps

import (
	"container/heap"
)

//KthSmallestElement : To find the kth smallest element in the array
func KthSmallestElement(array []int, k int) interface{} {
	h := &MaxHeap{}
	heap.Init(h)
	for i := 0; i < len(array); i++ {
		heap.Push(h, array[i])
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	ans := heap.Pop(h)
	return ans
}
