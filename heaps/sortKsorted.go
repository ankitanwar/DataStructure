package heaps

import "container/heap"

//SortKsorted : To sort the k sorted array
func SortKsorted(arr []int, k int) []int {
	h := &MinHeap{}
	heap.Init(h)
	current := 0
	for i := 0; i < len(arr); i++ {
		heap.Push(h, arr[i])
		if h.Len() > k {
			arr[current] = heap.Pop(h).(int)
			current++
		}
	}
	for {
		if h.Len() == 0 {
			break
		}
		arr[current] = heap.Pop(h).(int)
		current++
	}

	return arr
}
