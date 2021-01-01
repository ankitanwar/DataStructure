package array

import "fmt"

//DutchFlagAlgo : Given an array consisting of 0,1,2 and we need to sort it
func DutchFlagAlgo(arr []int) {
	low := 0
	mid := 0
	high := len(arr) - 1
	for mid <= high {
		if arr[mid] == 0 {
			arr[low], arr[mid] = arr[mid], arr[low]
			mid++
			low++
		} else if arr[mid] == 1 {
			mid++
		} else if arr[mid] == 2 {
			arr[high], arr[mid] = arr[mid], arr[high]
			high--
		}
	}
	fmt.Println(arr)
}
