package backtracking

import (
	"fmt"
	"strconv"
)

//FriendsPairing : Number of ways in which friends can be paired
func FriendsPairing(n int) {
	occupied := make([]bool, n+1)
	ans := ""
	solvePairing(n, 0, &occupied, &ans)
}

func solvePairing(total, current int, arr *[]bool, asf *string) {
	if current == total {
		fmt.Printf("%v ", asf)
		println("")
		return
	}
	if (*arr)[current] == true {
		solvePairing(total, current+1, arr, asf)
	} else {
		(*arr)[current] = true
		fmt.Println("(", current, ")")
		for j := current + 1; j <= total; j++ {
			if (*arr)[j] == false {
				(*arr)[j] = true
				*(asf) += "(" + strconv.Itoa(current) + strconv.Itoa(j) + ")"
				solvePairing(total, current+1, arr, asf)
				(*arr)[j] = false
			}
		}
		(*arr)[current] = false
	}
}
