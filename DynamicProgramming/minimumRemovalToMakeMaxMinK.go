package dynamicprogramming

//MinimumRemoval :It will return the minimum removal require to make make max-min < k
func MinimumRemoval(arr []int, min, max, k, n int) int {
	if n < 2 {
		return 999999 // to make sure two has to be distinguish
	}
	if arr[max-1]-arr[min] <= k {
		return 0
	}
	return 1 + minimum(MinimumRemoval(arr, min+1, max, k, n-1), MinimumRemoval(arr, min, max-1, k, n-1))

}

//MinimumRemoval2 : It will return the minimum removal require to make max - min less than 2
func MinimumRemoval2(arr []int, k int) int {
	matrix := [100][100]int{}

	return solveMinimumRemoval(&matrix, arr, 0, len(arr), len(arr), k)
}

func solveMinimumRemoval(matrix *[100][100]int, arr []int, left, right, n, k int) int {
	if matrix[left][right] != 0 {
		return matrix[left][right]
	}
	if n < 2 {
		matrix[left][right] = 999999
		return matrix[left][right]
	}
	if arr[right-1]-arr[left] <= k {
		matrix[left][right] = 0
		return matrix[left][right]
	}
	matrix[left][right] = 1 + minimum(solveMinimumRemoval(matrix, arr, left+1, right, n-1, k), solveMinimumRemoval(matrix, arr, left, right-1, n, k))
	return matrix[left][right]
}
