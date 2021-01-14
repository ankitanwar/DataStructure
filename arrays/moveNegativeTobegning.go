package array

import "fmt"

//MoveNegativeToLeft : Given an array we need to move all the negative element to the lest side
func MoveNegativeToLeft(arr []int) {
	//The idea is keep all the neagative at the left of the array and positive at the right by pation along the 0 and then at negative swap it with positive element ans swap the array

	currentIndex := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			currentIndex++
			arr[currentIndex], arr[i] = arr[i], arr[currentIndex]
		}
	}
	negative := 0
	positive := currentIndex + 1
	for {
		if negative >= len(arr) || positive >= len(arr) {
			break
		}
		arr[negative], arr[positive] = arr[positive], arr[negative]
		negative += 2
		positive++
	}
	fmt.Println(arr)

}
