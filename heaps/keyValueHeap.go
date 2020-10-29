package heaps

import "fmt"

//KeyValue : To implemnet key value interface in a heap
type KeyValue struct {
	Key   int
	Value int
}

//KeyHeap : KeyValue Heap
type KeyHeap []*KeyValue

func (h KeyHeap) Len() int { return len(h) }

func (h KeyHeap) Less(i, j int) bool { return h[i].Key < h[j].Key }
func (h KeyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

//Push : To add value in the heap
func (h *KeyHeap) Push(x interface{}) {
	item := x.(*KeyValue)
	*h = append(*h, item)
}

//Pop : To remove the value from the heap
func (h *KeyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//Print : To print all the values in the heap
func (h *KeyHeap) Print() {
	for i := 0; i < h.Len(); i++ {
		fmt.Printf("%v ", (*h)[i])
	}
	println("")
}
