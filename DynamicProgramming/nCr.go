package dynamicprogramming

//NcR : To calulate the value of NcR
func NcR(n, r int) int {
	if r > n {
		return 0
	}
	if r == 0 || r == n {
		return 1
	}
	return NcR(n-1, r-1) + NcR(n-1, r)
}

//NcR2 : To calulate the value of ncr according to DP
func NcR2(n, r int) int {
	matrix := [][]int{}
	for i := 0; i <= n; i++ {
		temp := []int{}
		for j := 0; j <= r; j++ {
			temp = append(temp, -1)
		}
		matrix = append(matrix, temp)
	}
	ans := calculate(matrix, n, r)
	return ans

}

func calculate(mat [][]int, n, r int) int {
	if mat[n][r] != -1 {
		return mat[n][r]
	}
	if r > n {
		mat[n][r] = 0
		return 0
	}
	if r == 0 || r == n {
		mat[n][r] = 1
		return 1
	}
	mat[n][r] = NcR(n-1, r-1) + NcR(n-1, r)
	return mat[n][r]

}
