package dynamicprogramming

import "fmt"

//DiffBetweenAdjIs1 : To count all the longestSubsequence Such that difference between adj is 1
func DiffBetweenAdjIs1(arr []int) {
	ans := []int{}
	for i := 0; i < len(arr); i++ {
		ans = append(ans, 1)
	}
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j]-arr[i] == 1 || arr[j]-arr[i] == -1 {
				ans[i] = maximum(ans[i], ans[j]+1)
			}
		}
	}

	for i := 0; i < len(ans); i++ {
		fmt.Printf("%v ", ans[i])
	}
	fmt.Println("")
}
