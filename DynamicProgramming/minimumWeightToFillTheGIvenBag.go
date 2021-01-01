package dynamicprogramming

import "fmt"

//MinimumWeightToFillTheBag : Given a weight we need to fill it and minimize the cost the cost[i] indicates the cost to fill i kg
func MinimumWeightToFillTheBag(weight int, cost []int) {
	/*
		We may assume that we have infinite supply of every kg
	*/
	ans := helperWeightMinimum(weight, 1, cost)
	if ans == 99999999 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
	helperWeightMinimumDp(weight, cost)
}

func helperWeightMinimum(weight, currentIndex int, cost []int) int {
	if weight == 0 {
		return 0
	}
	if currentIndex > len(cost) {
		return 99999999
	}
	if weight >= currentIndex && cost[currentIndex-1] != -1 {
		cost := minimum(cost[currentIndex-1]+helperWeightMinimum(weight-currentIndex, currentIndex, cost), helperWeightMinimum(weight, currentIndex+1, cost))
		return cost
	}
	return helperWeightMinimum(weight, currentIndex+1, cost)
}

func helperWeightMinimumDp(weight int, cost []int) {
	dp := [][]int{}
	for i := 0; i <= len(cost); i++ {
		t := make([]int, weight+1)
		dp = append(dp, t)
	}
	for i := 0; i <= len(cost); i++ {
		for j := 0; j <= weight; j++ {
			if j == 0 {
				dp[i][j] = 0
			} else if i == 0 {
				dp[i][j] = 99999
			} else {
				if j >= i && cost[i-1] != -1 {
					dp[i][j] = minimum(cost[i-1]+dp[i][j-i], dp[i-1][j])
				} else {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	fmt.Println(dp[len(cost)][weight])

}
