package backtracking

import "fmt"

var num string

//MaxNumberPossible : To find out the maximum number by doing at most k swap
func MaxNumberPossible(s string, k int) {
	maxNumberHelper(s, 0, k)
	fmt.Println(num)
}

func swap(s string, i, j int) string {
	start := string(s[:i])
	first := string(s[i : i+1])
	mid := string(s[i+1 : j])
	second := string(s[j : j+1])
	rem := string(s[j+1:])
	s = start + second + mid + first + rem
	return s
}

func maxNumberHelper(number string, currentSwap, k int) {
	if num < number {
		num = number
	}
	if currentSwap == k {
		return
	}
	for i := 0; i < len(number)-1; i++ {
		for j := i + 1; j < len(number); j++ {
			if number[i] < number[j] {
				s := swap(number, i, j)
				maxNumberHelper(s, currentSwap+1, k)
			}
		}
	}
}
