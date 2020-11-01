package dynamicprogramming

//LargestSum : To calculate the largest sum of the sub array possible
func LargestSum(arr []int) int {
	var ans int

	ans = arr[0]
	temp := arr[0]

	for i := 1; i < len(arr); i++ {
		temp = maximum(temp+arr[i], arr[i])
		ans = maximum(temp, ans)
	}

	return ans
}
