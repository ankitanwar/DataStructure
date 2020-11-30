package backtracking

import "fmt"

//WordBreak -> given a string and map of words we need to find out the numbers of ways in which we can divide the given string into the the map of given characters
func WordBreak(given, currentAns string, words map[string]int) {
	if len(given) == 0 {
		fmt.Println(currentAns)
		return
	}
	var till string
	for i := 0; i < len(given); i++ {
		till += string(given[i])
		_, found := words[till]
		if found {
			currentAns += till + " "
			right := given[i+1:]
			WordBreak(right, currentAns, words)
		}
	}

}
