package stackandqueues

import (
	"fmt"
)

// TwoStackInArray : to implement the two stack in an single array
type TwoStackInArray struct {
	arr   []int
	size  int
	first int
	last  int
}

//Init : constructor
func (t *TwoStackInArray) Init(size int) {
	t.arr = make([]int, size)
	t.first = 0
	t.last = size
}

//Add1 : To add the value in the stack
func (t *TwoStackInArray) Add1(value int) {
	if t.first+1 < t.last {
		t.first++
		t.arr[t.first] = value
		return
	}
}

//Add2 : To add the value in the stack
func (t *TwoStackInArray) Add2(value int) {
	if t.first < t.last-1 {
		t.last--
		t.arr[t.last] = value
		return
	}
	fmt.Println(" Stack overflow")
}

//Pop1 : to remove the element from the first stack
func (t *TwoStackInArray) Pop1() int {
	if t.first < 0 {
		return -1
	}
	removed := t.arr[t.first]
	t.first--
	return removed
}

// Pop2 : To remove the element from the second stack
func (t *TwoStackInArray) Pop2() int {
	if t.last < len(t.arr) {
		removed := t.arr[t.last]
		t.last++
		return removed
	}
	return -1

}

// Print : To print all the elements in the array
func (t *TwoStackInArray) Print() {
	for i := 0; i < len(t.arr); i++ {
		fmt.Print(t.arr[i])
		fmt.Print(" ")
	}
	fmt.Println("")
}
