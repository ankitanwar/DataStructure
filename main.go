package main

import "github.com/ankitanwar/Golang-DataStructure/backtracking"

func main() {
	given := "leetcode"
	dict := make(map[string]int)
	dict["leet"] = 1
	dict["code"] = 1
	backtracking.WordBreak(given, "", dict)
}
