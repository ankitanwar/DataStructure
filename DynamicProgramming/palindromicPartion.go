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
func PailndromicPartionDp(given string) {
	dp := [][]bool{}
	for i := 0; i < len(given); i++ {
		temp := make([]bool, len(given))
		dp = append(dp, temp)
	}
	for g := 0; g < len(given); g++ {
		column := g
		for row := 0; row < len(given); row++ {
			if column == len(dp) {
				break
			}
			if g == 0 {
				dp[row][column] = true
			} else if g == 1 {
				if given[row] == given[column] {
					dp[row][column] = true
				} else {
					dp[row][column] = false
				}
			} else {
				if given[row] == given[column] {
					dp[row][column] = dp[row+1][column-1]
				} else {
					dp[row][column] = false
				}
			}
			column++
		}
	}
	cuts := [][]int{}
	for i := 0; i < len(given); i++ {
		t := make([]int, len(given))
		cuts = append(cuts, t)
	}
	for g := 0; g < len(given); g++ {
		col := g
		for row := 0; row < len(given); row++ {
			if col == len(cuts) {
				break
			}
			if g == 0 {
				cuts[row][col] = 0
			} else if g == 1 {
				if given[row] == given[col] {
					cuts[row][col] = 0
				} else {
					cuts[row][col] = 1
				}
			} else {
				if dp[row][col] == true {
					cuts[row][col] = 0
				} else {
					var min int = 999999
					for k := row; k < col; k++ {
						left := cuts[row][k]
						right := cuts[k+1][col]
						if left+right+1 < min {
							min = left + right + 1
						}
					}
					cuts[row][col] = min
				}
			}
			col++
		}
	}
	fmt.Println(cuts)

}
