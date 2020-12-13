package dynamicprogramming

import "fmt"

//NumberOfBSTPossible : Given an interger n we have to find out the total number of bst possible with the given array
func NumberOfBSTPossible(n int) {
	//approach is kinda similar to the catalan number
	dp := []int{}
	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			dp = append(dp, 1)
		} else {
			x := 0
			y := i - 1
			var ans int
			for {
				if y < 0 {
					break
				}
				temp := dp[x] * dp[y]
				y--
				x++
				ans += temp
			}
			dp = append(dp, ans)
		}
	}
	fmt.Println(dp)
}
