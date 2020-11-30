package backtracking

var ans []string

// RemoveInvalidParentheses : To remove the invalid parentheses from the given string
func RemoveInvalidParentheses(s string) []string {
	ans = ans[:0]
	minr := minRemovals(s)
	helpers(s, minr)
	return ans
}

func helpers(s string, minRemovalsAllowed int) {
	if minRemovalsAllowed == 0 {
		if minRemovals(s) == 0 && contains(s) == false {
			ans = append(ans, s)
		}
	}
	for i := 0; i < len(s); i++ {
		left := s[0:i]
		right := s[i+1:]
		helpers(left+right, minRemovalsAllowed-1)
	}

}

func minRemovals(s string) int {
	var stack []string
	for i := 0; i < len(s); i++ {
		current := string(s[i])
		if current == ")" {
			if len(stack) == 0 {
				stack = append(stack, ")")
			} else if stack[len(stack)-1] == ")" {
				stack = append(stack, ")")
			} else {
				stack = stack[:len(stack)-1]
			}
		} else if current == "(" {
			stack = append(stack, "(")
		}

	}

	return len(stack)

}

func contains(target string) bool {
	for i := 0; i < len(ans); i++ {
		if ans[i] == target {
			return true
		}
	}
	return false
}
