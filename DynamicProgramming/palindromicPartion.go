package dynamicprogramming

import "fmt"

//PailndromicPartion : Minimum Number of partion require to make it paildrome
func PailndromicPartion(given string, i, j int) int {
	if i == j || i > j {
		return 0
	}
	check := isPalindrome(given[i : j+1])
	if check == true {
		return 0
	}
	var res int = 99999999999
	for k := i; k < j; k++ {
		ans := 1 + PailndromicPartion(given, i, k) + PailndromicPartion(given, k+1, j)
		if ans < res {
			res = ans
		}
	}
	return res
}

func isPalindrome(given string) bool {
	i := 0
	j := len(given) - 1
	for {
		if i >= j {
			break
		}
		if given[i] != given[j] {
			return false
		}
		i++
		j--
	}
	return true

}

//PailndromicPartionDp : To Calc the minimum number of pations required to make the string palindorme
func PailndromicPartionDp(given string, i, j int) int {
	matrix := [][]int{}
	for i := 0; i <= len(given); i++ {
		temp := []int{}
		for j := 0; j <= len(given); j++ {
			temp = append(temp, -1)
		}
		matrix = append(matrix, temp)
	}
	ans := palindromicDP(&matrix, given, 0, len(given)-1)
	return ans
}
func palindromicDP(matrix *[][]int, given string, i, j int) int {
	if i == j || i > j {
		return 0
	}
	if (*matrix)[i][j] != -1 {
		fmt.Println("Thiis matrix is working fine")
		return (*matrix)[i][j]
	}
	check := isPalindrome(given[i : j+1])
	if check == true {
		(*matrix)[i][j] = 0
		return (*matrix)[i][j]
	}
	ans := 99999999
	for k := i; k < j; k++ {
		tempAns := 1 + palindromicDP(matrix, given, i, k) + palindromicDP(matrix, given, k+1, j)
		if tempAns < ans {
			ans = tempAns
		}
	}
	(*matrix)[i][j] = ans
	return (*matrix)[i][j]

}
