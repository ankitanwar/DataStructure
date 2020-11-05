package heaps

import (
	"container/heap"
	"fmt"
)

//MedianInStream : To give the value of the media in a stream
func MedianInStream(array []int) {
	high := &MaxHeap{}
	low := &MinHeap{}
	heap.Init(high)
	heap.Init(low)
	for i := 0; i < len(array); i++ {
		current := array[i]
		addNumber(current, high, low)
		balance(high, low)
		fmt.Println(getMedian(high, low))
	}
	fmt.Printf("%v ", high)
	fmt.Printf("%v ", low)
}

//addNumber : To add the numbers into the heap
func addNumber(number int, h1 *MaxHeap, h2 *MinHeap) {
	if h1.Len() == 0 || h1.Peek().(int) > number {
		heap.Push(h1, number)
	} else {
		heap.Push(h2, number)
	}
}

func balance(h1 *MaxHeap, h2 *MinHeap) {
	if h1.Len() > h2.Len() && h1.Len()-h2.Len() >= 2 {
		temp := heap.Pop(h1)
		heap.Push(h2, temp)
	} else if h1.Len() < h2.Len() {
		temp := heap.Pop(h2)
		heap.Push(h1, temp)
	}
}

//getMedian : It will return the median of an element
func getMedian(h1 *MaxHeap, h2 *MinHeap) int {
	if h1.Len() == h2.Len() {
		ans := (h1.Peek().(int) + h2.Peek().(int)) / 2
		return ans
	}
	return h1.Peek().(int)
}
