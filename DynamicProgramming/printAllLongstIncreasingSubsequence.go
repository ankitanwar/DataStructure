package dynamicprogramming

import (
	"fmt"
	"strconv"
)

type subsequencePath struct {
	path  string
	index int
	value int
}

//LongestIncreasingSubsequence : To print the all longest increasing subsequence
func LongestIncreasingSubsequence(arr []int) {
	dp := []int{}
	dp = append(dp, 1)
	var maxValue int = 1
	var maxValueIndex = 0
	for i := 1; i < len(arr); i++ {
		var temp int = 1
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				temp = maximum(temp, dp[j]+1)
			}
		}
		if temp > maxValue {
			maxValue = temp
			maxValueIndex = i
		}
		dp = append(dp, temp)
	}
	t := subsequencePath{
		path:  strconv.Itoa(arr[maxValueIndex]),
		index: maxValueIndex,
		value: maxValue,
	}
	path := []subsequencePath{}
	path = append(path, t)
	for len(path) > 0 {
		p := path[0]
		path = path[1:]
		if p.value == 1 {
			fmt.Println(p.path)
		} else {
			for i := 0; i < p.index; i++ {
				if dp[i] == p.value-1 && arr[i] < arr[p.index] {
					t := subsequencePath{
						path:  p.path + "->" + strconv.Itoa(arr[i]),
						index: i,
						value: dp[i],
					}
					path = append(path, t)
				}
			}
		}

	}

}
