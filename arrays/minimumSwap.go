package array

import "fmt"

//MinimumSwap : Given an array of n positive integers and a number k. Find the minimum number of swaps required to bring all the numbers less than or equal to k together.
func MinimumSwap(arr []int, k int) {
	var smallerElement int
	var ans int = 1e4
	for i := 0; i < len(arr); i++ {
		if arr[i] <= k {
			smallerElement++
		}
	}
	var presentElements int
	for i := 0; i < smallerElement; i++ {
		if arr[i] <= k {
			presentElements++
		}
	}
	ans = min(ans, smallerElement-presentElements)
	i := 0
	for j := smallerElement; j < len(arr); j++ {
		if arr[i] <= k {
			presentElements--
		}
		if arr[j] <= k {
			presentElements++
		}
		ans = min(ans, smallerElement-presentElements)
		i++
	}
	fmt.Println(ans)
}
