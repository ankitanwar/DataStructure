package heaps

import (
	"babbar/linkedlist"
	"container/heap"
)

// MergeKsortedLinkedList : It will merge the linked list
func MergeKsortedLinkedList(l1, l2, l3 *linkedlist.Linkedlist) *linkedlist.Linkedlist {
	current1 := l1.Head
	current2 := l2.Head
	current3 := l3.Head

	h := &MinHeap{}
	heap.Init(h)

	for {
		if current1 != nil {
			heap.Push(h, current1.Value.(int))
			current1 = current1.Next
		}
		if current2 != nil {
			heap.Push(h, current2.Value.(int))
			current2 = current2.Next
		}
		if current3 != nil {
			heap.Push(h, current3.Value.(int))
			current3 = current3.Next
		}
		if current1 == nil && current2 == nil && current3 == nil {
			break
		}
	}

	newList := &linkedlist.Linkedlist{}

	for {
		if h.Len() == 0 {
			break
		}
		current := heap.Pop(h)
		newList.Insert(current)
	}

	return newList

}
