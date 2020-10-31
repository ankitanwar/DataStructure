package dynamicprogramming

//MaxSquareWithSize1 : to find out the the largest square possible with given 1
func MaxSquareWithSize1(matrix [][]int) int {
	var _max int
	ans := [][]int{}
	for i := 0; i <= len(matrix); i++ {
		temp := []int{}
		for j := 0; j <= len(matrix[0]); j++ {
			temp = append(temp, -1)
		}
		ans = append(ans, temp)
	}

	for i := 0; i <= len(matrix); i++ {
		for j := 0; j <= len(matrix[0]); j++ {
			if i == 0 || j == 0 {
				ans[i][j] = 0
			} else if matrix[i-1][j-1] == 0 {
				ans[i][j] = 0
			} else {
				ans[i][j] = 1 + min(ans[i-1][j-1], ans[i-1][j], ans[i][j-1])
			}

		}
	}

	for i := 0; i < len(ans); i++ {
		for j := 0; j < len(ans[0]); j++ {
			if ans[i][j] > _max {
				_max = ans[i][j]
			}
		}
	}

	return _max
}
