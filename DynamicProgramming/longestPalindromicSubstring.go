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
func LongesPalindromicSubstringDP(s string) {
	println("hello world")
}
