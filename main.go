package main

import (
	"babbar/heaps"
	"container/heap"
	"fmt"
)

func main() {
	h := &heaps.MaxHeap{}
	heap.Init(h)
	h2 := &heaps.MaxHeap{}
	heap.Init(h2)

	arr := []int{10, 5, 6, 2}
	arr2 := []int{12, 7, 9}

	for i := 0; i < len(arr); i++ {
		h.Push(arr[i])
	}
	for i := 0; i < len(arr2); i++ {
		h.Push(arr2[i])
	}

	ans := heaps.MergeTwoMaxHeaps(h, h2)
	fmt.Printf("%v ", ans)

}
