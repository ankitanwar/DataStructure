package heaps

import (
	"fmt"
	"math"
)

//Heaps : heaps struct
type Heaps struct {
	Arr      []int
	Size     int
	Capacity int
}

//Newheaps : constructor to the heap
func Newheaps(cap int) *Heaps {
	temp := &Heaps{
		Arr:      []int{},
		Size:     0,
		Capacity: cap,
	}
	return temp
}

//Parent : This will return the parent of the node
func (h *Heaps) Parent(index int) int {
	return (index - 1) / 2
}

//Leftchild : This will return the left child of that node
func (h *Heaps) Leftchild(index int) int {
	return 2*index + 1
}

//Rightchild : this will return the right child of the node
func (h *Heaps) Rightchild(index int) int {
	return 2*index + 2
}

//Insert : To insert the values into the heap(Min heap)
func (h *Heaps) Insert(value int) {
	if h.Size == h.Capacity {
		return
	}
	h.Arr = append(h.Arr, value)
	i := h.Size
	for {
		if i == 0 || h.Arr[h.Parent(i)] < h.Arr[i] {
			break
		}
		h.Arr[i], h.Arr[h.Parent(i)] = h.Arr[h.Parent(i)], h.Arr[i]
		i = h.Parent(i)
	}
	h.Size++
}

//Print : To print all the values in the heap
func (h *Heaps) Print() {
	for i := 0; i < len(h.Arr); i++ {
		fmt.Printf("%v ", h.Arr[i])
	}
	println("")
}

//MinHeapify : To re arrange the heap in (min heap)
func (h *Heaps) MinHeapify(i int) {
	smallest := i
	right := h.Rightchild(i)
	left := h.Leftchild(i)

	if right < h.Size && h.Arr[right] < h.Arr[smallest] {
		smallest = right
	}
	if left < h.Size && h.Arr[left] < h.Arr[smallest] {
		smallest = left
	}

	if smallest != i {
		h.Arr[i], h.Arr[smallest] = h.Arr[smallest], h.Arr[i]
		h.MinHeapify(smallest)
	}
}

//MaxHeapify : to re arrange the heap in Max Heap
func (h *Heaps) MaxHeapify(index int) {
	largest := index
	right := h.Rightchild(index)
	left := h.Leftchild(index)

	if right < h.Size && h.Arr[right] > h.Arr[largest] {
		largest = right
	}
	if left < h.Size && h.Arr[left] > h.Arr[largest] {
		largest = left
	}
	if largest != index {
		h.Arr[largest], h.Arr[index] = h.Arr[index], h.Arr[largest]
		h.MaxHeapify(largest)
	}

}

//GetMin : It will return the minimum value in the heap works like a pop functio
func (h *Heaps) GetMin() int {
	if h.Size == 0 {
		return int(math.Inf(0))
	}
	if h.Size == 1 {
		return h.Arr[0]
	}

	h.Arr[0], h.Arr[h.Size-1] = h.Arr[h.Size-1], h.Arr[0]
	ans := h.Arr[h.Size-1]
	h.Size--
	h.MinHeapify(0)
	return ans
}

//Update : To update the value in the heap
func (h *Heaps) Update(value, index int) {
	if h.Size < index {
		return
	}
	h.Arr[index] = value
	h.UpHeapify(index)

}

//UpHeapify : To rearrange all the elements
func (h *Heaps) UpHeapify(index int) {
	for {
		if index == 0 || h.Arr[h.Parent(index)] < h.Arr[index] {
			break
		}
		h.Arr[h.Parent(index)], h.Arr[index] = h.Arr[index], h.Arr[h.Parent(index)]
		index = h.Parent(index)
	}
}

//Delete : To delete the value in the heap
func (h *Heaps) Delete(index int) {
	if h.Size < index {
		return
	}
	h.Arr[index] = int(math.Inf(-1))

	//After this we are actually doing extract minimum
	//Pushing the require element to the last and the heapify the rest of the tree
	h.UpHeapify(index)
	h.Arr[0], h.Arr[h.Size-1] = h.Arr[h.Size-1], h.Arr[0]
	h.Size--
	h.MinHeapify(0)

	h.Print()

}

//BuildMinHeap : we have given an array and we need to build the min heap out of it
func (h *Heaps) BuildMinHeap() {
	//wt we do is we start from the bottom node and call Minheapify till we reach 0
	//why bottom because in heapify we assume that the below nodes are already heapfiy
	//complexity O(n)
	for i := (h.Size - 2) / 2; i >= 0; i-- {
		h.MinHeapify(i)
	}
}

//BuildMaxHeap : It will build the max heap
func (h *Heaps) BuildMaxHeap() {
	for i := (h.Size - 2) / 2; i >= 0; i-- {
		h.MaxHeapify(i)
	}
}
