package dynamicprogramming

import (
	"fmt"
	"strconv"
)

//MinimumJumpsToReachEnd : it will return the minimum numbers of steps required to reach the end
func MinimumJumpsToReachEnd(array []int) int {
	dp := make([]int, len(array))
	path := make([]int, len(array))
	for i := 0; i < len(array); i++ {
		dp[i] = 9999
	}
	path[0] = 0

	dp[0] = 0
	for i := 1; i < len(array); i++ {
		for j := 0; j < i; j++ {
			if j+array[j] >= i {
				if dp[i] > 1+dp[j] {
					dp[i] = 1 + dp[j]
					path[i] = array[j]
				}
			}
		}
	}

	fmt.Printf("%v ", path)

	return dp[len(array)-1]
}

type jumpPath struct {
	path  string
	jump  int
	index int
}

//MinimumJumpsToReachTheEnd2 : To calc the minimum jumps required to reacht the end
func MinimumJumpsToReachTheEnd2(arr []int) {
	dp := make([]int, len(arr))
	dp[len(dp)-1] = 0
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i]+i >= len(arr)-1 {
			dp[i] = 1
		} else if arr[i] == 0 {
			dp[i] = 99999999
		} else {
			min := 99999999
			for j := 1; j <= arr[i]; j++ {
				min = minimum(min, dp[i+j])
			}
			dp[i] = 1 + min
		}
	}
	path := []jumpPath{}
	t := jumpPath{
		path:  "0",
		jump:  dp[0],
		index: 0,
	}
	path = append(path, t)
	for len(path) > 0 {
		fmt.Println(path)
		p := path[0]
		path = path[1:]
		if p.jump == 0 {
			println("This is working")
			fmt.Print(p.path)
		} else {
			for i := 1; i <= p.jump; i++ {
				if i+p.index < len(arr) && dp[i+p.index] == p.jump-1 {
					t := jumpPath{
						path:  p.path + "->" + strconv.Itoa(i+p.index),
						jump:  dp[i+p.index],
						index: i + p.index,
					}
					path = append(path, t)
				}
			}
		}
	}
}
