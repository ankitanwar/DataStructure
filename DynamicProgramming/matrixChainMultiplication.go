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

//MatrixChainMultiplication2 : To calc the minimum cost required to multiply an matrix
func MatrixChainMultiplication2(arr []int) {
	dp := [][]int{}
	for i := 0; i < len(arr); i++ {
		temp := make([]int, len(arr))
		dp = append(dp, temp)
	}
	for g := 0; g < len(arr); g++ {
		col := g
		for row := 0; row < len(arr); row++ {
			if col == len(arr) {
				break
			}
			if g == 0 {
				dp[row][col] = 0
			} else if g == 1 {
				dp[row][col] = arr[row] * arr[col] * arr[col+1]
			} else {
				var min int = 99999999999
				for k := row; k < col; k++ {
					left := dp[row][k]
					right := dp[k+1][col]
					costOfLeftAndRight := arr[row] * arr[k+1] * arr[col+1]
					if min > left+right+costOfLeftAndRight {
						min = left + right + costOfLeftAndRight
					}
				}
				dp[row][col] = min
			}
			col++
		}
	}
}
