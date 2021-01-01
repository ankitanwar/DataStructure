package array

import "fmt"

//MaximumAndMinimumInArray : Given an array we need to cal the maximum and minimum in the given array
func MaximumAndMinimumInArray(array []int) {
	//check wether the given array if of odd length or even length

	var max int
	var min int
	var i int

	//for even length
	if len(array)%2 == 0 {
		if array[0] > array[1] {
			max = array[0]
			min = array[1]
		} else {
			max = array[1]
			min = array[0]
		}
		i = 2
	} else { // for odd array
		min = array[0]
		max = array[0]
		i = 1
	}
	for {
		if i >= len(array) {
			break
		}
		if array[i] < array[i+1] {
			if array[i+1] > max {
				max = array[i+1]
			}
			if array[i] < min {
				min = array[i]
			}
		} else {
			if array[i+1] < min {
				min = array[i+1]
			}
			if array[i] > max {
				max = array[i]
			}
		}

		i += 2
	}
	fmt.Println("Max and min value in the given array is ", max, min)

}
