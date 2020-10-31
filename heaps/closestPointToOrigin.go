package heaps

import (
	"container/heap"
	"math"
)

//ClosestPointToOrigin : IT will return K closest Point To origin
func ClosestPointToOrigin(matrix [][]int, k int) [][]int {
	ans := [][]int{}

	h := &MaxKeyHeap{}
	heap.Init(h)

	for i := 0; i < len(matrix); i++ {
		temp := &KeyValue{
			Key:   int(math.Abs(float64(matrix[i][0]*matrix[i][0])) + math.Abs(float64(matrix[i][1]*matrix[i][1]))),
			Value: i,
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
		current := heap.Pop(h).(*KeyValue).Value
		coordinates := []int{matrix[current][0], matrix[current][1]}
		ans = append(ans, coordinates)
	}

	return ans
}
