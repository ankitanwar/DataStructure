package backtracking

import "fmt"

//Idea is first count the frequenct of each character and then divide it into the half and then calc all the permutations of first half and then add the second half in it to make it palindrome

//PalindromicPermutation : To print all the palindromic permutation of a string
func PalindromicPermutation(given string) {
	freq := make(map[string]int)
	for i := 0; i < len(given); i++ {
		_, found := freq[string(given[i])]
		if found {
			freq[string(given[i])]++
		} else {
			freq[string(given[i])] = 1
		}
	}
	var oddCount int
	var oddChar string
	var newStringLen int
	hashmap := make(map[string]int)
	for key := range freq {
		if freq[key]%2 != 0 {
			oddCount++
			oddChar += key
			hashmap[key] = freq[key] / 2
			newStringLen += freq[key] / 2

		} else {
			hashmap[key] = freq[key] / 2
			newStringLen += freq[key] / 2

		}
	}
	calcPermutation(oddChar, "", oddCount, newStringLen, 0, &hashmap)
}

func calcPermutation(oddChar, ans string, oddCount, newStringLen, currentIndex int, hashmap *map[string]int) {
	if currentIndex == newStringLen {
		reverse := ""
		for i := 0; i < len(ans); i++ {
			reverse += string(ans[i])
		}
		if oddCount != 0 {
			ans += oddChar
		}
		ans += reverse
		fmt.Println(ans)
		return

	}
	for key := range *hashmap {
		if (*hashmap)[key] != 0 {
			(*hashmap)[key]--
			calcPermutation(oddChar, ans+key, oddCount, newStringLen, currentIndex+1, hashmap)
			(*hashmap)[key]++
		}
	}
}
