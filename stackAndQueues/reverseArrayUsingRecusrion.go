package stackandqueues

//RecusriveSortArray : to sore the array recusrively
func (s *Stack) RecusriveSortArray() []int {
	if len(s.arr) == 1 || len(s.arr) == 1 {
		return s.arr
	}
	current := s.Pop()
	s.RecusriveSortArray()
	s.arr = insert(current, s.arr)
	return s.arr
}

func insert(element int, arr []int) []int {
	if len(arr) == 0 {
		arr = append(arr, element)
		return arr
	}
	current := arr[len(arr)-1]
	println("the value of current is ", current)
	arr = arr[:len(arr)-1]
	if current < element {
		arr = append(arr, element)
		arr = append(arr, current)
		return arr
	}
	arr = insert(element, arr)
	arr = append(arr, current)
	return arr

}
