package dynamicprogramming

import "fmt"

//MaxSumIncreasingSub : It will give the sum of maximum subsequence
func MaxSumIncreasingSub(arr []int) int {
	ans := []int{}
	var res int
	ans = append(ans, arr[0])
	for i := 1; i < len(arr); i++ {
		current := arr[i]
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				current += arr[j]
			}
		}
		ans = append(ans, current)
	}

	println("The valur of the ans is ")
	for i := 0; i < len(ans); i++ {

		fmt.Printf("%v ", ans[i])
	}
	println("")

	for i := 0; i < len(ans); i++ {
		if res < ans[i] {
			res = ans[i]
		}
	}

	return res
}
