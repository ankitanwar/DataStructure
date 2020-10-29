package dynamicprogramming

//MatrixChainMultiplication : To perform the matrix chain multiplication
func MatrixChainMultiplication(array []int, i, j int) int {
	if i == j {
		return 0
	}
	_min := 9999999
	for k := i; k < j; k++ {
		count := MatrixChainMultiplication(array, i, k) + MatrixChainMultiplication(array, k+1, j) + (array[i-1] * array[k] * array[j])
		if count < _min {
			_min = count
		}

	}

	return _min
}

//MatrixChainMultiplicationDp : To calcilate mcm with dp
func MatrixChainMultiplicationDp(array []int) int {
	dp := [][]int{}
	for m := 0; m <= len(array); m++ {
		temp := []int{}
		for n := 0; n <= len(array); n++ {
			temp = append(temp, -1)
		}
		dp = append(dp, temp)
	}
	return solve(dp, array, 1, len(array)-1)
}

func solve(mat [][]int, array []int, i, j int) int {
	if mat[i][j] != -1 {
		println("This is working perfectly fine")
		return mat[i][j]
	}
	if i >= j {
		return 0
	}
	_min := 9999999
	for k := i; k < j; k++ {
		count := solve(mat, array, i, k) + solve(mat, array, k+1, j) + (array[i-1] * array[k] * array[j])
		mat[i][j] = count
		if count < _min {
			_min = count
		}

	}

	return _min

}
