package stackandqueues

import "fmt"

//Nstack : To intiliaze the Nstack in an single array
type Nstack struct {
	array []int
	free  int
	top   []int
	next  []int
}

//Nstacks : to initialize the N stacks in an arary
func (n *Nstack) Nstacks(length, k int) {
	n.array = make([]int, length)
	n.free = 0
	n.top = make([]int, k+1)
	n.next = make([]int, length)
	for i := 0; i < length-1; i++ {
		n.next[i] = i + 1
	}
	n.next[len(n.next)-1] = -1
}

//Insert : To insert the value in the stack
func (n *Nstack) Insert(value, k int) {
	if n.free == -1 {
		return
	}
	n.array[n.free] = value
	temp := n.free
	n.free = n.next[n.free]
	n.next[temp] = n.top[k]
	n.top[k] = temp

}

//Delete : To delete the value from the stack
func (n *Nstack) Delete(k int) string {
	if n.top[k] == -1 {
		return " stack is already empty"
	}

	index := n.top[k]
	n.array[index] = -1
	n.top[k] = n.next[index]
	n.next[index] = n.free
	n.free = index
	return "item has been deleted successfully"

}

//Print : to print the element in the stack
func (n *Nstack) Print() {
	for i := 0; i < len(n.array); i++ {
		fmt.Printf("%v ", n.array[i])
	}
	fmt.Println("")
}
