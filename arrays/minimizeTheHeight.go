package array

import (
	"fmt"
	"sort"
)

//MinimizeTheHeight : Given an array we can add or subtract k value from evey element(make sure it doesn't become negative after subtracting) and we need to minimize the diff between the maximum and minimum of the array
func MinimizeTheHeight(arr []int, k int) {
	sort.Ints(arr)
	last := arr[len(arr)-1]
	if last-k >= 0 {
		last = last - k
	} else {
		last = last + k
	}
	fmt.Println(arr)
	fmt.Println("The value of last is ", last)
	s := arr[0]
	var s1 int = 999999
	var s2 int
	if s-k > 0 {
		s1 = s - k
	}
	s2 = s + k
	fmt.Println(min(abs(last-s1), abs(last-s2)))

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
