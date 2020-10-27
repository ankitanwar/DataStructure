package stackandqueues

import "fmt"

//NextSmaller : To give the next smaller element in the give array
func NextSmaller(arr []int) {
	ans := []int{}
	stack := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		current := arr[i]

		for {
			if len(stack) == 0 || stack[len(stack)-1] < current {
				break
			}
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			ans = append(ans, -1)
		} else {
			ans = append(ans, stack[len(stack)-1])
		}
		stack = append(stack, current)
	}

	left := 0
	right := len(ans) - 1
	for left < right {
		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}
	println("The bvalue of ans is ")
	for i := 0; i < len(ans); i++ {
		fmt.Printf("%v ", ans[i])
	}
}
