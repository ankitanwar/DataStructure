package dynamicprogramming

import "fmt"

//EvulateExpressionToTrue :
func EvulateExpressionToTrue(s string) int {
	ans := evulateHelper(s, 0, len(s)-1, true)
	return ans
}

func evulateHelper(s string, i, j int, value bool) int {
	if i > j {
		return 0
	}
	if i == j {
		if string(s[i]) == "t" {
			if value == true {
				return 1
			}
			return 0
		} else if string(s[i]) == "f" {
			if value == true {
				return 0
			}
			return 1
		}
	}
	var ans int
	for k := i + 1; k <= j-1; k += 2 {
		leftT := evulateHelper(s, i, k-1, true)
		rightT := evulateHelper(s, k+1, j, true)
		leftF := evulateHelper(s, i, k-1, false)
		rightF := evulateHelper(s, k+1, j, false)

		if string(s[k]) == "&" {
			if value == true {
				ans += leftT * rightT
			} else {
				ans += rightF * leftF
			}
		} else if string(s[k]) == "|" {
			if value == true {
				ans += rightT * leftF
				ans += rightT * leftT
				ans += leftF * rightT
			} else {
				ans += leftF * rightF
			}
		} else if string(s[k]) == "^" {
			if value == false {
				ans += rightT * leftT
				ans += rightF * leftF
			} else {
				ans += rightF * leftT
				ans += leftT * rightT
			}
		}

	}
	fmt.Println("the value of ans is ", ans)

	return ans
}
