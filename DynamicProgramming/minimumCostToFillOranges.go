package dynamicprogramming

//MinimumCostToFillOranges : it will return the minimum cost to fill the oranges
func MinimumCostToFillOranges(arr []int, capacity, n int) int {
	if capacity == 0 {
		return 0
	}
	if n == 0 && capacity != 0 {
		return 9999
	}
	if capacity >= n {
		return minimum(arr[n-1]+MinimumCostToFillOranges(arr, capacity-n, n), MinimumCostToFillOranges(arr, capacity, n-1))
	}
	return MinimumCostToFillOranges(arr, capacity, n-1)

}

// MinimumCostToFillOranges2 : It will give the minimum cost require to fill the oranges in the given bag
func MinimumCostToFillOranges2(arr []int, capacity int) int {
	matrix := [100][100]int{}

	return memoSolve(&matrix, arr, capacity, len(arr))
}

func memoSolve(mat *[100][100]int, arr []int, capacity, n int) int {
	if mat[n][capacity] != 0 {
		println("This is working")
		return mat[n][capacity]
	}
	if capacity == 0 {
		return 0
	}
	if n == 0 && capacity != 0 {
		mat[n][capacity] = 9999
		return mat[n][capacity]
	}
	if capacity >= n {
		mat[n][capacity] = minimum(arr[n-1]+memoSolve(mat, arr, capacity-n, n), memoSolve(mat, arr, capacity, n-1))
		return mat[n][capacity]
	}
	mat[n][capacity] = memoSolve(mat, arr, capacity, n-1)
	return mat[n][capacity]

}
