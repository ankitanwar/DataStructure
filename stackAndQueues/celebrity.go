package stackandqueues

import "fmt"

// Celebrity : To know whether the given there is any celbrity or not
func Celebrity(arr [][]int) int {
	s := Stack{}
	for i := 0; i < len(arr); i++ {
		s.Push(i)
	}
	for {

		if len(s.arr) < 2 {
			break
		}

		first := s.Pop()
		second := s.Pop()

		if arr[first][second] == 1 {
			s.Push(second)
		} else {
			s.Push(first)
		}
	}

	potential := s.Pop()
	fmt.Println("The potential is ", potential)
	for i := 0; i < len(arr); i++ {

		if i == potential {
			continue

		}

		if arr[potential][i] == 1 || arr[i][potential] == 0 {
			fmt.Println("The value is ", potential, i, arr[i][potential])
			return -1
		}
	}

	return potential
}
