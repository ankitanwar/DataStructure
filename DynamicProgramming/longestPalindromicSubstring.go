package dynamicprogramming

var sublen int
var anss string

//LongesPalindromicSubstring : To calc the length of the longest palindromic substring
func LongesPalindromicSubstring(s string) string {
	helpersubstring(0, s)
	return anss
}

func helpersubstring(currentindex int, s string) {
	if currentindex == len(s) {
		return
	}
	for i := currentindex; i < len(s); i++ {
		current := string(s[currentindex : i+1])
		check := isPalindrome(current)
		if check == true {
			if len(current) > sublen {
				sublen = len(current)
				anss = current
			}
		}
	}
	helpersubstring(currentindex+1, s)
}

//LongesPalindromicSubstringDP : To find out the length of the longest palindromic substring
func LongesPalindromicSubstringDP(s string) string {
	var ans string = string(s[0])
	dp := [][]bool{}
	for i := 0; i < len(s); i++ {
		t := make([]bool, len(s))
		dp = append(dp, t)
	}

	for g := 0; g < len(s); g++ {
		row := 0
		col := g
		for {
			if row == len(s) || col == len(s) {
				break
			}
			if g == 0 {
				dp[row][col] = true
			} else if g == 1 {
				if string(s[row]) == string(s[col]) {
					ans = string(s[row]) + string(s[col])
					dp[row][col] = true
				} else {
					dp[row][col] = false
				}
			} else {
				if string(s[row]) == string(s[col]) {
					if dp[row+1][col-1] == true {
						dp[row][col] = true
					} else {
						dp[row][col] = false
					}
				} else {
					dp[row][col] = false
				}
			}
			row++
			col++
		}
	}
	for col := len(s) - 1; col >= 0; col-- {
		currentCol := col
		row := 0
		for {
			if row == len(s) || currentCol == len(s) {
				break
			}
			if dp[row][currentCol] == true {
				ans = ""
				for i := row; i <= currentCol; i++ {
					ans += string(s[i])
				}
				return ans
			}

			row++
			currentCol++
		}
	}
	return ans
}
