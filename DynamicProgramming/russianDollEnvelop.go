package dynamicprogramming

import (
	"sort"
)

type russianDoll struct {
	width  int
	height int
}

// RussianDollEnvelop : To count the max number of sqaures that can be fit inside the
func RussianDollEnvelop(height, width []int) {
	arr := []russianDoll{}
	for i := 0; i < len(height); i++ {
		t := &russianDoll{
			width:  width[i],
			height: height[i],
		}
		arr = append(arr, *t)
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i].height < arr[j].height })
	dp := []int{}
	dp = append(dp, 1)
	for i := 1; i < len(arr); i++ {
		var temp int = 1
		for j := 0; j < i; j++ {
			if arr[j].width < arr[i].width {
				temp = maximum(temp, dp[j]+1)
			}
		}
		dp = append(dp, temp)
	}
	var ans int = -1
	for i := 0; i < len(dp); i++ {
		if dp[i] > ans {
			ans = dp[i]
		}
	}

}
