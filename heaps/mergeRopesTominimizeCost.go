package heaps

import "container/heap"

//ConnectRopes : We have to connect ropes such that the cost of adding them is minimum
func ConnectRopes(array []int) int {

	h := &MinHeap{}
	heap.Init(h)

	for i := 0; i < len(array); i++ {
		heap.Push(h, array[i])
	}

	var cost int

	for {
		if h.Len() == 1 {
			return cost
		}
		first := heap.Pop(h).(int)
		second := heap.Pop(h).(int)
		sum := first + second
		cost += sum
		heap.Push(h, sum)

	}

}
