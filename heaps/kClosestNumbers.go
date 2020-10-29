package heaps

import (
	"container/heap"
	"math"
)

//KClosestNumbers : to print the values k closest numbers
func KClosestNumbers(arr []int, k, num int) []int {
	ans := []int{}

	h := &KeyHeap{}
	heap.Init(h)

	for i := 0; i < len(arr); i++ {
		temp := &KeyValue{
			Key:   int(math.Abs(float64(num - arr[i]))),
			Value: arr[i],
		}
		heap.Push(h, temp)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	for {
		if h.Len() == 0 {
			break
		}
		current := heap.Pop(h).(*KeyValue)
		ans = append(ans, current.Value)

	}

	return ans
}
