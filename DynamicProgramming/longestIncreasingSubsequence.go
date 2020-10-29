package dynamicprogramming

//IncreasingSubsequence : To Calc the increasing Subsequence in the given array
func IncreasingSubsequence(arr []int) int {
	ans := []int{}
	for i := 0; i < len(arr); i++ {
		ans = append(ans, 1)
	}
	for i := 1; i < len(arr); i++ { // current we are at i
		for j := 0; j < i; j++ { // checking value for all the previous element
			if arr[j] < arr[i] && ans[j] >= ans[i] {
				ans[i] = ans[j] + 1
			}
		}
	}

	max := ans[0]
	for i := 1; i < len(ans); i++ {
		if ans[i] > max {
			max = ans[i]
		}
	}
	return max
}
