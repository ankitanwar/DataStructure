package dynamicprogramming

import "fmt"

//InterleavedString : it will return whether the 3 string is innterleabed string or not
func InterleavedString(str1, str2, str3 string) bool {
	ans := solveInterleave(str1, str2, str3, len(str1), len(str2), len(str3))
	matrix := [10][10]bool{}
	fmt.Println("Ans withoud dp is ", ans)
	solveInterleaveDP(&matrix, str1, str2, str3, len(str1), len(str2), len(str3))
	return matrix[len(str1)][len(str2)]

}

func solveInterleave(str1, str2, str3 string, i, j, k int) bool {
	if i == 0 && j == 0 && k == 0 {
		return true
	}
	if k == 0 && (j != 0 || i != 0) {
		return false
	}
	var first bool
	var second bool

	if i > 0 {
		first = str1[i-1] == str3[k-1] && solveInterleave(str1, str2, str3, i-1, j, k-1)
	}
	if j > 0 {
		second = str2[j-1] == str3[k-1] && solveInterleave(str1, str2, str3, i, j-1, k-1)
	}
	return first || second
}

func solveInterleaveDP(matrix *[10][10]bool, str1, str2, str3 string, i, j, k int) bool {
	if i == 0 && j == 0 && k == 0 {
		return true
	}
	if k == 0 && (j != 0 || i != 0) {
		return false
	}
	if matrix[i][j] != false {
		println("This is working")
		return matrix[i][j]
	}
	var first bool
	var second bool

	if i > 0 {
		first = str1[i-1] == str3[k-1] && solveInterleaveDP(matrix, str1, str2, str3, i-1, j, k-1)
	}
	if j > 0 {
		second = str2[j-1] == str3[k-1] && solveInterleaveDP(matrix, str1, str2, str3, i, j-1, k-1)
	}

	matrix[i][j] = first || second
	return matrix[i][j]
}
