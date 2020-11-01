package heaps

import "babbar/linkedlist"

//LinkHeap : linkedlist heap
type LinkHeap struct {
	Key   int
	value *linkedlist.Node
}

//MinLinkedHeap : to implement linked list heap
type MinLinkedHeap []*LinkHeap

func (h MinLinkedHeap) Len() int { return len(h) }

func (h MinLinkedHeap) Less(i, j int) bool { return h[i].Key < h[j].Key }
func (h MinLinkedHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// MergeKsortedLinkedList : It will merge the linked list
func MergeKsortedLinkedList(l1, l2, l3 *linkedlist.Linkedlist) {
	current1 := l1.GetHead()
	current2 := l2.GetHead()
	current3 := l3.GetHead()

	for {
		if current1.value == nil {

		}
	}

}
