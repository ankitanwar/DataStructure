package backtracking

import "fmt"

var used []bool

//PermutationOfAString2 : To find the permutation of a string
func PermutationOfAString2(s string) {
	used = make([]bool, len(s))
	permutataionHelper(s, "")
}

func permutataionHelper(s, asf string) {
	if len(asf) == len(s) {
		fmt.Println(asf)
		return
	}
	for i := 0; i < len(s); i++ {
		if used[i] == false {
			used[i] = true
			permutataionHelper(s, asf+string(s[i]))
			used[i] = false
		}
	}
}
