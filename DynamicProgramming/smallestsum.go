package dynamicprogramming

//SmallestSum : to find the smallest sum of contiguous sub array
func SmallestSum(arr []int) int {
	var ans int

	ans = arr[0]
	temp := arr[0]
	for i := 1; i < len(arr); i++ {
		temp = minimum(temp+arr[i], arr[i])
		ans = minimum(ans, temp)

	}

	return ans
}
