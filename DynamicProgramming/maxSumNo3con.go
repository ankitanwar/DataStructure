package dynamicprogramming

//MaxSumNo3Consecutive : Maximum sum such that no 3 element are consecutive
func MaxSumNo3Consecutive(arr []int, n, count int) int {
	if n < 0 {
		return 0
	}
	if count == 2 {
		return MaxSumNo3Consecutive(arr, n-1, 0)
	}
	return maximum(arr[n]+MaxSumNo3Consecutive(arr, n-1, count+1), MaxSumNo3Consecutive(arr, n-1, count))

}
