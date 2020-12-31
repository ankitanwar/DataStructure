package dynamicprogramming

import (
	"fmt"
)

//MinimizeTheSubsetSum : Given an array we need to divide it into the two parts such that sum of part1 - sum of aprt2 should me minimum
func MinimizeTheSubsetSum(array []int) {
	/*
		The idea is to find the sum of all all the array lets call it range
		now we have to minimize s2-s1
		now s2=range-s1
		so, range-s1-s1--> so we have to minime range -2(s1)

	*/
	dp := [][]bool{}
	var sum int
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	for i := 0; i <= len(array); i++ {
		t := make([]bool, sum+1)
		dp = append(dp, t)
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			if j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = false
			} else {
				if array[i-1] <= j {
					dp[i][j] = dp[i-1][j] || dp[i-1][j-array[i-1]]
				} else {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	r := (sum / 2) + 1
	var ans int = 99999
	for i := 0; i <= r; i++ {
		if dp[len(array)][i] == true {
			t := sum - 2*i
			if t < ans {
				ans = t
			}
		}
	}
	fmt.Println(ans)
}
