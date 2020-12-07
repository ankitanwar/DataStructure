package backtracking

import "fmt"

//PermutationOfAString : To print all the permutataion of a string
func PermutationOfAString(given string, k int) {
	ans := make([]string, k)
	helperPermutation(given, &ans, 0, 0)
}

func helperPermutation(given string, ans *[]string, currentIndex, currentPartion int) {
	if currentPartion == len(*ans) {
		fmt.Printf("%v ", ans)
		println()
		return
	}
	for i := 0; i < len(*ans); i++ {
		if len((*ans)[i]) == 0 {
			(*ans)[i] = string(given[currentIndex])
			helperPermutation(given, ans, currentIndex+1, currentPartion+1)
			(*ans)[i] = ""
		}

	}
}
