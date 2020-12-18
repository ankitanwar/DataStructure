package greedy

import (
	"container/heap"
	"fmt"

	"github.com/ankitanwar/Golang-DataStructure/heaps"
)

//MinimumCostOfRopes : We have given a cost of ropes we need to tell the minimum cost requires to join them
func MinimumCostOfRopes(cost []int) {
	h := &heaps.MinHeap{}
	heap.Init(h)
	for i := 0; i < len(cost); i++ {
		heap.Push(h, cost[i])
	}
	var ans int
	for h.Len() > 1 {
		first := heap.Pop(h).(int)
		second := heap.Pop(h).(int)
		temp := first + second
		ans += temp
		heap.Push(h, temp)

	}
	fmt.Println(ans)
}
