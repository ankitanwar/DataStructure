package array

import "fmt"

//KthSmallest : Given an array we need to find out the kth smallest in the array
func KthSmallest(array []int, k int) {
	//The idea is to use quick select algorithm
	low := 0
	high := len(array) - 1

	for {
		get := partion(&array, low, high)
		println("The value of get is ", get)
		if get+1 == k {
			fmt.Println(array[get])
			break
		} else if get+1 > k {
			high = get - 1
		} else {
			low = get + 1
		}
	}

}

func partion(arr *[]int, low, high int) int {
	fmt.Println("The value of array is ", arr)
	i := -1
	for {
		if low >= high {
			break
		}
		if (*arr)[low] <= (*arr)[high] {
			i++
			(*arr)[i], (*arr)[low] = (*arr)[low], (*arr)[i]

		}
		low++
	}

	(*arr)[i+1], (*arr)[high] = (*arr)[high], (*arr)[i+1]
	return i + 1

}
