package dynamicprogramming

import "fmt"

//LongestBitonicSubsequence : To find the longst bitonic subsequence bitnonic subsequence is a subsequence which is increasing at first and then starts to decreasing
func LongestBitonicSubsequence(arr []int) {
	lis := []int{}
	lis = append(lis, 1) //its the first element and nothing is at the left
	for i := 1; i < len(arr); i++ {
		var temp int = 1
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				temp = maximum(temp, lis[j]+1)
			}
		}
		lis = append(lis, temp)
	}

	//now finding the backward decreasing subsequence
	dis := make([]int, len(arr))
	dis[len(arr)-1] = 1
	for i := len(arr) - 2; i >= 0; i-- {
		var temp int = 1
		for j := len(arr) - 1; j > i; j-- {
			if arr[i] > arr[j] {
				temp = maximum(temp, dis[j]+1)
			}
		}
		dis[i] = temp
	}
	var temp int = -1
	for i := 0; i < len(arr); i++ {
		temp = maximum(temp, lis[i]+dis[i]-1)
	}
	fmt.Println(temp)

}
