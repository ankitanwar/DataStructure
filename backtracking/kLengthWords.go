package backtracking

import (
	"fmt"
)

var count int

//KLengthWords : We have to place
func KLengthWords(given string, k int) {
	count = 0
	ans := ""
	helperLengthWords(given, ans, k, 0, 0)
	fmt.Println(count)
}

func helperLengthWords(given, ans string, k, selected, currentIndex int) {
	if currentIndex == len(given) {
		if selected == k {
			count++
			fmt.Println(ans)
		}
		return
	}
	current := string(given[currentIndex])
	for i := currentIndex; i < len(given); i++ {
		helperLengthWords(given, ans+current, k, selected+1, currentIndex+1)
		helperLengthWords(given, ans, k, selected, currentIndex+1)
	}
}
