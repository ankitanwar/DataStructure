package stackandqueues

var min int

//Insert : To insert the element into the stack
func (s *Stack) Insert(val int) {
	if len(s.arr) == 0 {
		min = val
		s.Push(val)
		return
	}
	if val > min {
		s.Push(val)
	} else {
		new := 2*val - min
		min = val
		s.Push(new)
	}
}

//GetMin :To return the minimum element in O(1)
func (s *Stack) GetMin() int {
	return min
}

//Remove : To remove the element from the stack
func (s *Stack) Remove() int {
	current := s.Pop()
	if current < min {
		min = 2*min - current
	}
	return current
}
