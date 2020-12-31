package dynamicprogramming

import (
	"fmt"
	"sort"
)

//PairWithSpecificDifference : Given an array We need to find those pairs whose difference is less than K and we need to maximize the differnece of that pair
func PairWithSpecificDifference(array []int, k int) {
	sort.Ints(array)
	pairs := [][]int{}
	i := len(array) - 2
	for i >= 0 {
		if array[i+1]-array[i] < k {
			t := []int{array[i+1], array[i]}
			pairs = append(pairs, t)
			i -= 2
		} else {
			i--
		}
	}
	fmt.Println(pairs)
}
