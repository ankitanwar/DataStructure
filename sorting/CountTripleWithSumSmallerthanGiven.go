package sorting

import (
	"sort"
)

//CountTriple : it returns the number of count in which sum is smaller than the given target
func CountTriple(arr []int, target int) int {
	count := 0
	sort.Ints(arr)
	for i := 0; i < len(arr)-2; i++ {
		current := arr[i]
		low := i + 1
		high := len(arr) - 1

		for low < high {
			sum := current + arr[low] + arr[high]

			if sum >= target {
				high--
			} else {
				count += high - low
				low++
			}

		}
	}

	return count
}
