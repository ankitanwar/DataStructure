package dynamicprogramming

import "fmt"

//cost :- (blankSpaceInLine)**2

//WordBreak : Given a length of the line we need to give the minimize the cost to place the characters in the line
func WordBreak(given []int, lengthOfLine int) {
	helperWordBreak(given, lengthOfLine)
}

var current int = 0
var cost int

func helperWordBreak(given []int, length int) {
	matrix := [][]int{}
	for j := 0; j < len(given); j++ {
		temp := make([]int, len(given))
		matrix = append(matrix, temp)
	}
	for i := 0; i < len(given); i++ {
		temp := length
		var next bool
		extraSpace := 0
		for j := i; j < (len(given)); j++ {
			rem := temp - given[j] - extraSpace
			if rem < 0 {
				next = true
			}
			if next == true {
				matrix[i][j] = 9999999
			} else {
				matrix[i][j] = rem * rem
				temp = rem
			}
			extraSpace++
		}
	}
	fmt.Println(matrix)
	cost := make([]int, len(given))
	arrangement := make([]int, len(given))

	for row := len(given) - 1; row >= 0; row-- {
		col := len(given) - 1
		if matrix[row][col] < 100 {
			cost[row] = matrix[row][col]
		} else {
			arrangement = append(arrangement, 0)
		}
	}
}
