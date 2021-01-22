package array

import "fmt"

//MinimumOperations : Given an array we need to tell the minimum number of opeartions require to make it a plaindrome we can merge two adjacent elements means replacing them with their sum
func MinimumOperations(arr []int) {
	i := 0
	var count int
	j := len(arr) - 1
	for i < j {
		if arr[i] != arr[j] {
			if arr[i] < arr[j] {
				arr[i+1] += arr[i]
				i++
			} else {
				arr[j-1] += arr[j]
				j--
			}
			count++
		} else {
			i++
			j--
		}
	}
	fmt.Println(count)

}
