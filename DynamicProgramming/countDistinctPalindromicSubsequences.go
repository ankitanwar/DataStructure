package dynamicprogramming

import "fmt"

//CountDistinctPalindromicSubsequences : To count all the number of distinct palindrome subsequences in a string
func CountDistinctPalindromicSubsequences(s string) {
	dp := [][]int{}
	for i := 0; i < len(s); i++ {
		t := make([]int, len(s))
		dp = append(dp, t)
	}
	nxt := []int{} // to keep track of the next occurance of the same character
	pev := []int{} //to keep track of the previous occurance of the same character
	ndict := make(map[string]int)
	for i := 0; i < len(s); i++ {
		_, found := ndict[string(s[i])]
		if found {
			nxt = append(nxt, ndict[string(s[i])])
			ndict[string(s[i])] = i
		} else {
			nxt = append(nxt, -1)
			ndict[string(s[i])] = i
		}
	}
	pdict := make(map[string]int)
	for i := len(s) - 1; i >= 0; i-- {
		_, found := pdict[string(s[i])]
		if found {
			pev = append(pev, pdict[string(s[i])])
			pdict[string(s[i])] = i
		} else {
			pev = append(pev, -1)
			pdict[string(s[i])] = i
		}
	}
	for g := 0; g < len(s); g++ {
		col := g
		for row := 0; row < len(s); row++ {
			if col == len(s) {
				break
			}
			if g == 0 {
				dp[row][col] = 1
			} else if g == 1 {
				dp[row][col] = 2
			} else {
				if s[row] != s[col] {
					dp[row][col] = dp[row][col-1] + dp[row+1][col] - dp[row+1][col-1]
				} else {
					next := nxt[row]
					previous := pev[col]
					if next > previous {
						dp[row][col] = (2 * dp[row+1][col-1]) + 2
					} else if next == previous {
						dp[row][col] = (2 * dp[row+1][col-1]) + 1
					} else {
						dp[row][col] = (2 * dp[row+1][col-1]) - dp[next+1][previous-1]
					}
				}
			}
			col++
		}
	}
	fmt.Println(dp)
}
