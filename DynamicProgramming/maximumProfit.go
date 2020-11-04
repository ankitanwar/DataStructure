package dynamicprogramming

import (
	"babbar/heaps"
	"container/heap"
)

//MaximumProfit : It will return the maximum profit when shares can be bought at max twice
func MaximumProfit(arr []int) int {
	profit := &heaps.MaxHeap{}
	heap.Init(profit)

	var current int

	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] > current {
			current = arr[i] - arr[i-1]
		} else {
			heap.Push(profit, current)
			current = 0
		}
	}
	if current != 0 {
		heap.Push(profit, current)
	}
	if profit.Len() < 2 {
		return 0
	}

	var ans int
	for {
		if profit.Len() == 0 {
			break
		}
		ans += heap.Pop(profit).(int)
	}

	return ans

}
