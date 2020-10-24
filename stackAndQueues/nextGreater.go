package stackandqueues

//NextGreater :to find the element greater to the right
func NextGreater(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	s := Stack{}
	ans := []int{}

	for i := len(arr) - 1; i >= 0; i-- {
		current := arr[i]
		for {

			if len(s.arr) == 0 || s.arr[len(s.arr)-1] > current {
				break
			}
			s.Pop()
		}

		if len(s.arr) == 0 {
			ans = append(ans, -1)
			s.Push(current)
		} else {
			ans = append(ans, s.arr[len(s.arr)-1])
			s.Push(current)
		}
	}

	return ans

}
