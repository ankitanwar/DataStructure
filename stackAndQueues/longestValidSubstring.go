package stackandqueues

import "math"

//LongestValidSubstring : to calc the length of longest valid sunstring
func LongestValidSubstring(s string) int {
	stack := []int{}
	ans := -9999
	stack = append(stack, -1)

	for i := 0; i < len(s); i++ {
		current := string(s[i])
		if current == ")" {
			stack = stack[:len(stack)-1]
			if len(stack) != 0 {
				ans = int(math.Max(float64(ans), float64(i-stack[len(stack)-1])))
			} else {
				stack = append(stack, i)
			}
		} else {
			stack = append(stack, i)
		}
	}
	return ans
}
