package stackandqueues

import "fmt"

//FirstNegativeInWindow : To calc the firstnegative integer in every window
func FirstNegativeInWindow(arr []int, k int) {
	stack := []int{}
	ans := []int{}
	for i := 0; i < k; i++ {
		current := arr[i]
		if current < 0 {
			stack = append(stack, i)
			if len(ans) == 0 {
				ans = append(ans, current)
			}
		}
	}
	left := 1
	for right := k; right < len(arr); right++ {
		current := arr[right]
		if current < 0 {
			stack = append(stack, right)
		}
		if len(stack) != 0 && left > stack[0] {
			stack = stack[1:]
		}
		if len(stack) != 0 {
			ans = append(ans, arr[stack[0]])
		} else {
			ans = append(ans, 0)
		}
		left++

	}

	for i := 0; i < len(ans); i++ {
		fmt.Printf("%v ", ans[i])
	}

	println("")
}
