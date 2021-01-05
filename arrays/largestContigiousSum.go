package array

import "fmt"

//LargestContigousSum : given an array we need to find out the largest contigous sum
func LargestContigousSum(arr []int) {
	var ans int
	var currentSum int = arr[0]
	for i := 1; i < len(arr); i++ {
		currentSum = max(currentSum+arr[i], arr[i])
		ans = max(ans, currentSum)
	}
	fmt.Println(ans)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
