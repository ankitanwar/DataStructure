package heaps

import "fmt"

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
