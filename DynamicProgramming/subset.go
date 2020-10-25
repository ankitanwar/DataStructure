package dynamicprogramming

import "fmt"

//Subset : whether its possible to break the array into 2 or not
func Subset(arr []int, sum, n int) bool {
	if sum == 0 {
		return true
	}
	if n == 0 && sum != 0 {
		return false
	}
	if sum%2 != 0 {
		return false
	}
	newsum := sum / 2
	return findPartation(arr, newsum, n)

}

func findPartation(arr []int, add, n int) bool {
	if n == 0 && add != 0 {
		return false
	}
	if add == 0 {
		return true
	}
	if add >= arr[n-1] {
		return findPartation(arr, add-arr[n-1], n-1) || findPartation(arr, add, n-1)
	}
	return findPartation(arr, add, n-1)
}

//Subset2 : To solve the subset program dynamically
func Subset2(arr []int, sum int) bool {
	matrix := [][]bool{}
	for i := 0; i <= len(arr); i++ {
		temp := []bool{}
		for j := 0; j <= sum; j++ {
			temp = append(temp, false)
		}
		matrix = append(matrix, temp)
	}

	for i := 0; i <= len(arr); i++ {
		matrix[i][0] = true
	}
	for i := 1; i <= sum; i++ {
		matrix[0][i] = false
	}

	for i := 1; i < len(arr); i++ {
		current := arr[i-1]
		for j := 1; j <= sum; j++ {
			if j >= current {
				matrix[i][j] = matrix[i-1][j-current] || matrix[i-1][j]
			} else {
				matrix[i][j] = matrix[i-1][j]
			}
		}
	}
	println("The values in matrix is ")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j <= sum; j++ {
			fmt.Printf("%v ", matrix[i][j])
		}
		println("")
	}
	return matrix[len(arr)][sum]
}
