package dynamicprogramming

//Lcs3String : To find the lcs of 3 given strings
func Lcs3String(str1, str2, str3 string, i, j, k int) int {
	if i == 0 || j == 0 || k == 0 {
		return 0
	}

	if str1[i-1] == str2[j-1] && str1[i-1] == str3[k-1] {
		return 1 + Lcs3String(str1, str2, str3, i-1, j-1, k-1)
	}
	return max(Lcs3String(str1, str2, str3, i-1, j, k), Lcs3String(str1, str2, str3, i, j-1, k), Lcs3String(str1, str2, str3, i, j, k-1))
}

//Lcs3Dp : To computer the lcs of given 3 string using dp
func Lcs3Dp(str1, str2, str3 string) {
	matrix := [][][]int{}

	for i := 0; i <= len(str1); i++ {
		temp := [][]int{}
		for j := 0; j <= len(str2); j++ {
			temp2 := []int{}
			for k := 0; k <= len(str3); k++ {
				temp2 = append(temp2, -1)
			}
			temp = append(temp, temp2)
		}
		matrix = append(matrix, temp)
	}

	for i := 0; i <= len(str1); i++ {
		for j := 0; j <= len(str2); j++ {
			for k := 0; k <= len(str3); k++ {
				if i == 0 || j == 0 || k == 0 {
					matrix[i][j][k] = 0
				} else if str1[i-1] == str2[j-1] && str1[i-1] == str3[k-1] {
					matrix[i][j][k] = 1 + matrix[i-1][j-1][k-1]
				} else {
					matrix[i][j][k] = max(matrix[i-1][j][k], matrix[i][j-1][k], matrix[i][j][k-1])
				}
			}
		}
	}

	println(matrix[len(str1)][len(str2)][len(str3)])
}
