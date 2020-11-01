package heaps

import (
	"fmt"
)

//KeyValue : To implemnet key value interface in a heap
type KeyValue struct {
	Key   int
	Value int
}

//KeyValuee : Struct for key as string and value as int
type KeyValuee struct {
	Key   int
	Value int
}

//MaxKeyHeap : Max KeyValue Heap
type MaxKeyHeap []*KeyValuee

func (h MaxKeyHeap) Len() int { return len(h) }

func (h MaxKeyHeap) Less(i, j int) bool { return h[i].Key > h[j].Key }
func (h MaxKeyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

//Peek : It will return the first index top of the heap
func (h *MinKeyHeap) Peek() interface{} {
	return (*h)[0]
}

//Peek : it will return the top value of the peek
func (h *MaxKeyHeap) Peek() interface{} {
	return (*h)[0]
}

//MinKeyHeap : to implement min key heap
type MinKeyHeap []*KeyValue

func (h MinKeyHeap) Len() int { return len(h) }

func (h MinKeyHeap) Less(i, j int) bool { return h[i].Key > h[j].Key }
func (h MinKeyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

//Push : To add value in the heap
func (h *MaxKeyHeap) Push(x interface{}) {
	item := x.(*KeyValuee)
	*h = append(*h, item)
}

//Pop : To remove the value from the heap
func (h *MaxKeyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//Push : To add value in the heap
func (h *MinKeyHeap) Push(x interface{}) {
	item := x.(*KeyValue)
	*h = append(*h, item)
}

//Pop : To remove the value from the heap
func (h *MinKeyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//Print : To print all the values in the heap
func (h *MaxKeyHeap) Print() {
	for i := 0; i < h.Len(); i++ {
		fmt.Printf("%v ", (*h)[i])
	}
	println("")
}
