package array

import "fmt"

//MoveNegativeToLeft : Given an array we need to move all the negative element to the lest side
func MoveNegativeToLeft(arr []int) {
	negative := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			negative++
			arr[negative], arr[i] = arr[i], arr[negative]
		}
	}
	fmt.Println(arr)
}
