package array

import "fmt"

//CyclicRotateArrayByk : Given an array we need to rotate it by k
func CyclicRotateArrayByk(arr []int, k int) {
	low := 0
	high := len(arr) - 1
	for low < high {
		arr[low], arr[high] = arr[high], arr[low]
		low++
		high--
	}
	s := 0
	e := k - 1
	for s < e {
		arr[s], arr[e] = arr[e], arr[s]
		s++
		e--
	}
	s = k
	e = len(arr) - 1
	for s < e {
		arr[s], arr[e] = arr[e], arr[s]
		s++
		e--
	}
	fmt.Println(arr)

}
