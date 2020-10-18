package sorting

import "fmt"

// SearchWhereAdjacentDiffer : it return the index of the element to be find where the difference between the adjacent element differ by k
func SearchWhereAdjacentDiffer(arr []int, find, k int) int {
	i := 0
	for i < len(arr) {
		fmt.Println("the value of i is ", i)
		if arr[i] == find {
			return i
		}
		i = i + max(abs(arr[i]-find)/k)
	}

	return -1
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a int) int {
	if a > 1 {
		return a
	}
	return 1
}
