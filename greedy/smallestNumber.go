package greedy

import (
	"fmt"
	"strconv"
)

//SmallestNumber : We have given a sum and len and we have to a make a number whose length is equal to given len and sum of digits is equal to sum
func SmallestNumber(length int, sum int) {
	if length == 1 {
		if sum == 1 {
			fmt.Println(1)
			return
		} else {
			fmt.Println("Not possible")
			return
		}
	}
	till := 0
	curr := ""
	for {
		if 9+till <= sum {
			till += 9
			curr += "9"

		} else {
			new := sum - till
			till += new
			break
		}
	}
	if len(curr) != length {
		for {
			if len(curr) == length {
				break
			}
			curr = "0" + curr
		}
	}
	ans := ""
	if string(curr[0]) == "0" {
		ans += "1"
		ans += curr[1 : len(curr)-1]
		last, _ := strconv.Atoi(string(curr[len(curr)-1]))
		last--
		ans += strconv.Itoa(last)
		fmt.Println(ans)

	} else {
		fmt.Println(till)
	}
}
