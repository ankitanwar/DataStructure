package heaps

import (
	"container/heap"
)

//MaxOfAllSubArrays : It will return the maximum of all sub arrays of size k
func MaxOfAllSubArrays(arr []int, k int) []int {
	ans := []int{}
	currentMax := 0
	for i := 0; i < k; i++ {
		if arr[currentMax] < arr[i] {
			currentMax = i
		}
	}
	h := &MaxKeyHeap{}
	heap.Init(h)
	temp := &KeyValuee{
		Key:   arr[currentMax],
		Value: currentMax,
	}
	heap.Push(h, temp)
	ans = append(ans, arr[currentMax])
	for i := k; i < len(arr); i++ {
		for {
			if h.Len() == 0 || h.Peek().(*KeyValuee).Value >= i-k+1 && h.Peek().(*KeyValuee).Key > arr[i] {
				temp := &KeyValuee{
					Key:   arr[i],
					Value: i,
				}
				heap.Push(h, temp)
				break
			}

			heap.Pop(h)
		}
		ans = append(ans, h.Peek().(*KeyValuee).Key)
	}

	return ans

}
