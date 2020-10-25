package dynamicprogramming

//EditDistance : To count the number of insertion and deletion require
func EditDistance(first, second string, m, n int) int {
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	if first[m-1] == second[n-1] {
		return EditDistance(first, second, m-1, n-1)
	}
	a := EditDistance(first, second, m-1, n-1) //replace
	b := EditDistance(first, second, m, n-1)   //insert
	c := EditDistance(first, second, m-1, n)   //delete
	ans := 1 + min(a, b, c)
	return ans
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < a && b < c {
		return b
	}
	return c
}

//EditDistance2 : To calulate the the problem using dynamic programming
func EditDistance2(first, second string, m, n int) int {
	matrix := [][]int{}
	for i := 0; i <= m; i++ {
		temp := []int{}
		for j := 0; j <= n; j++ {
			temp = append(temp, -1)
		}
		matrix = append(matrix, temp)
	}

	for i := 0; i <= m; i++ {
		matrix[i][0] = i
	}
	for i := 1; i <= n; i++ {
		matrix[0][i] = i
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if first[i-1] == second[j-1] {
				matrix[i][j] = matrix[i-1][j-1]
			} else {
				matrix[i][j] = 1 + min(matrix[i-1][j-1], matrix[i][j-1], matrix[i-1][j])
			}
		}
	}

	return matrix[m][n]
}
