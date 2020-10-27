package dynamicprogramming

//MaximizeCut : To maximize the number of cut possible
func MaximizeCut(RodLength, length int, arr []int) int {
	if RodLength == 0 && length > 0 {
		return 0
	}
	if length <= 0 && RodLength != 0 {
		return -99
	}

	if arr[length-1] <= RodLength {
		return maximum(1+MaximizeCut(RodLength-arr[length-1], length, arr), MaximizeCut(RodLength, length-1, arr))
	}
	return MaximizeCut(RodLength, length-1, arr)

}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//MaximizeCut2 : To maximize the number of cut in the rod of given lenght
func MaximizeCut2(RodLength int, arr []int) int {
	dp := [][]int{}
	for i := 0; i <= len(arr); i++ {
		temp := []int{}
		for j := 0; j <= RodLength; j++ {
			temp = append(temp, -1)
		}
		dp = append(dp, temp)
	}

	for i := 0; i <= RodLength; i++ {
		dp[0][i] = -1
	}
	for j := 0; j <= len(arr); j++ {
		dp[0][j] = 0
	}
	for i := 1; i <= len(arr); i++ {
		for j := 1; j <= RodLength; j++ {
			if j >= arr[i-1] {
				dp[i][j] = maximum(1+dp[i][j-arr[i-1]], dp[i-1][j])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(arr)][RodLength]
}
