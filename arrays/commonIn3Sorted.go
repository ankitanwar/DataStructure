package array

import "fmt"

//CommonIn3Sorted : Given 3 sorted array we need to find the common element in between
func CommonIn3Sorted(arr1, arr2, arr3 []int) {
	i := 0
	j := 0
	k := 0
	var ans []int
	for {
		if i == len(arr1) || j == len(arr2) || k == len(arr3) {
			break
		}
		if arr1[i] == arr2[j] && arr2[j] == arr3[k] {
			ans = append(ans, arr1[i])
			i++
			j++
			k++
		} else if arr1[i] < arr2[j] {
			i++
		} else if arr2[j] < arr3[k] {
			j++
		} else {
			k++
		}
	}
	fmt.Println(ans)
}
