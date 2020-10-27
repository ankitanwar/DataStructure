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
