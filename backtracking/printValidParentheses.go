package backtracking

var ans map[string]bool

// RemoveInvalidParentheses : To remove the invalid parentheses from the given string
func RemoveInvalidParentheses(s string) []string {
	ans = make(map[string]bool)
	minr := minRemovals(s)
	helpers(s, minr)
	res := []string{}
	for key := range ans {
		res = append(res, key)
	}
	return res
}

func helpers(s string, minRemovalsAllowed int) {
	if minRemovalsAllowed == 0 {
		if minRemovals(s) == 0 && contains(s) == false {
			ans[s] = true
		}
		return
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

//to avoid duplicate printing
func contains(target string) bool {
	_, found := ans[target]
	if found {
		return true
	}
	return false
}
