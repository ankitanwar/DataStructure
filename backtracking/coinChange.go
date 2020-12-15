package backtracking

import (
	"fmt"
	"strconv"
)

//Coinchange : We are given a set of coins and we need to determine whether we cant can make k out of it or not
func Coinchange(arr []int, k int) {
	coinChangeHelper(arr, k, 0, 0, "")
}

func coinChangeHelper(coins []int, target, currentSum, currentIndex int, asf string) {
	if currentSum == target {
		fmt.Println(asf)
		return
	}
	if currentIndex == len(coins) || currentSum > target {
		return
	}

	coinChangeHelper(coins, target, currentSum+coins[currentIndex], currentIndex, asf+strconv.Itoa(coins[currentIndex]))
	coinChangeHelper(coins, target, currentSum, currentIndex+1, asf)
}
