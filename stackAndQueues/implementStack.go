package stackandqueues

import "fmt"

//Stack : To implement the stack
type Stack struct {
	arr []int
}

// Push : To add the item into the stack
func (s *Stack) Push(val int) {
	s.arr = append(s.arr, val)
}

// Pop : To remove the value from the stack
func (s *Stack) Pop() int {
	removed := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return removed
}

//Print : To print all the values in the stack
func (s *Stack) Print() {
	for i := 0; i < len(s.arr); i++ {
		fmt.Print(s.arr[i])
		fmt.Printf(" ")
	}
	fmt.Println("")
}
