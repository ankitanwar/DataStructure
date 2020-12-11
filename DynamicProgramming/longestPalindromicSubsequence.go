package dynamicprogramming

import "fmt"

var length int

//LongestPalindromicSubsequence : To calc the longesPalindromicSubsequence
func LongestPalindromicSubsequence(s string) {
	helperPalindromic(s, "", 0)
	fmt.Println(length)
}

func helperPalindromic(s, current string, currentIndex int) {
	if currentIndex == len(s) {
		check := isPalindrome(current)
		if check == true {
			if len(current) > length {
				length = len(current)
			}
		}
		return
	}
	helperPalindromic(s, current+string(s[currentIndex]), currentIndex+1)
	helperPalindromic(s, current, currentIndex+1)

}
