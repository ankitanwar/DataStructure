package heaps

//MinimumSumFormed : It will return the minimum sum formed from digits of array
func MinimumSumFormed(arr []int, n int) int {
	if n == 0 {
		return 0
	}
	return arr[n-1] + MinimumSumFormed(arr, n-2) + arr[n-2] + MinimumSumFormed(arr, n-4)
}
