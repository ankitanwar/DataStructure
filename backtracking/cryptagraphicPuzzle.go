package backtracking

import (
	"fmt"
	"strconv"
)

//Cryptagraphy : it will return the cryptagraphy order of the puzzle
func Cryptagraphy(s1, s2, s3 string) {
	unique := ""
	hashmap := make(map[string]int)
	for i := 0; i < len(s1); i++ {
		_, found := hashmap[string(s1[i])]
		if !found {
			hashmap[string(s1[i])] = -1
			unique += string(s1[i])
		}
	}
	for i := 0; i < len(s2); i++ {
		_, found := hashmap[string(s2[i])]
		if !found {
			hashmap[string(s2[i])] = -1
			unique += string(s2[i])
		}
	}
	for i := 0; i < len(s3); i++ {
		_, found := hashmap[string(s3[i])]
		if !found {
			hashmap[string(s3[i])] = -1
			unique += string(s3[i])
		}
	}

	usedNumbers := make([]bool, 10)
	solve(&hashmap, s1, s2, s3, unique, &usedNumbers, 0)

}

func getNumber(s string, hashmap *map[string]int) int {
	number := ""
	for i := 0; i < len(s); i++ {
		number += strconv.Itoa((*hashmap)[string(s[i])])
	}
	new, _ := strconv.Atoi(number)
	return new
}
func toCharStr(i int) string {
	return string('a' - 1 + i)
}

func solve(hashmap *map[string]int, s1, s2, s3, unique string, usedNumbers *[]bool, currentIndex int) {
	if currentIndex == len(unique) {
		num1 := getNumber(s1, hashmap)
		num2 := getNumber(s2, hashmap)
		num3 := getNumber(s3, hashmap)

		if num1+num2 == num3 {
			for i := 1; i <= 26; i++ {
				_, found := (*hashmap)[toCharStr(i)]
				if found {
					fmt.Printf("%v-%v ", toCharStr(i), (*hashmap)[toCharStr(i)])
				}
			}
			fmt.Println("")
		}
		return
	}
	for i := 0; i <= 9; i++ {
		if (*usedNumbers)[i] == false {
			(*hashmap)[string(unique[currentIndex])] = i
			(*usedNumbers)[i] = true
			solve(hashmap, s1, s2, s3, unique, usedNumbers, currentIndex+1)
			(*hashmap)[string(unique[currentIndex])] = -1
			(*usedNumbers)[i] = false
		}
	}

}
