package backtracking

import (
	"fmt"
)

//PermutationIterative : to calc the permutation of a string
func PermutationIterative(given string) {
	length := len(given)
	fact := fac(length)
	for i := 0; i < fact; i++ {
		tempString := given
		temp := i
		for j := length; j >= 1; j-- {
			quotient := temp / j
			rem := temp % j
			fmt.Printf("%v ", string(tempString[rem]))
			tempString = tempString[:rem] + tempString[rem+1:]
			temp = quotient

		}
		println()
	}

}

func fac(n int) int {
	var ans int = 1
	for i := 2; i <= n; i++ {
		ans *= i
	}

	return ans

}
