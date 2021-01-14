package array

import "fmt"

var ans int = 0

//CountInversion : Given an array we need to count the number of inversion in the given array
func CountInversion(arr []int) {
	inversionHelper(arr)
	fmt.Println(ans)
}

func inversionHelper(arr []int) []int {
	if len(arr) == 1 {
		t := []int{arr[0]}
		return t
	}
	mid := len(arr) / 2
	l := arr[:mid]
	r := arr[mid:]
	l1 := inversionHelper(l)
	r1 := inversionHelper(r)
	a := inversionCountHelper(l1, r1)
	ans += a
	m := merge(l, r)
	return m
}

func inversionCountHelper(left, right []int) int {
	l := 0
	r := 0
	count := 0
	for {
		if l >= len(left) || r >= len(right) {
			break
		}
		if left[l] > right[r] {
			count += len(left) - l
			l++
		} else {
			r++
		}
	}
	return count
}

func merge(left, right []int) []int {
	ans := []int{}
	l := 0
	r := 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			ans = append(ans, left[l])
			l++
		} else {
			ans = append(ans, right[r])
			r++
		}
	}
	for l < len(left) {
		ans = append(ans, left[l])
		l++
	}
	for r < len(right) {
		ans = append(ans, right[r])
		r++
	}
	return ans

}
