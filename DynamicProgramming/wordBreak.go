package dynamicprogramming

import "fmt"

//WordWarp : Given a string and set of words of dictionary we need to find out wether it is possible to segment a string to space separated sequence of dictonary words
func WordWarp(s string, words []string) {
	dict := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		dict[string(words[i])] = true
	}
	ans := wordWarpRecurssive(s, dict)
	fmt.Println(ans)
}

func wordWarpRecurssive(s string, dict map[string]bool) bool {
	if len(s) == 0 {
		return true
	}
	var current string
	for i := 0; i < len(s); i++ {
		current += string(s[i])
		_, found := dict[current]
		if found {
			restString := string(s[i+1:])
			check := wordWarpRecurssive(restString, dict)
			if check == true {
				return true
			}
		}
	}
	return false
}
