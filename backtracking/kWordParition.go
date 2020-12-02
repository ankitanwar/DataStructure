package backtracking

import "fmt"

//EqualSumPartion : To partion the given array into k parts such that there sum should me equal
func EqualSumPartion(given []int, k int) {
	partion := make([][]int, k)
	sbsetSum := make([]int, k)
	helpSumPartion(given, 0, len(given), k, &sbsetSum, &partion)

}

func helpSumPartion(given []int, currentIndex, n, k int, subSetSum *[]int, ans *[][]int) {
	if currentIndex == len(given) {
		for i := 0; i < len(*ans); i++ {
			if (*ans)[i] == nil {
				return
			}
		}
		var equal bool = true
		for i := 0; i < len(*subSetSum)-1; i++ {
			if (*subSetSum)[i] != (*subSetSum)[i+1] {
				equal = false
				break
			}
		}
		if equal == true {
			for i := 0; i < len(*ans); i++ {
				fmt.Printf("%v ", (*ans)[i])
			}
			println()

		}
		return
	}
	for i := 0; i < len(*ans); i++ {
		if (*ans)[i] != nil {
			(*ans)[i] = append((*ans)[i], given[currentIndex])
			(*subSetSum)[i] += given[currentIndex]
			helpSumPartion(given, currentIndex+1, n, k, subSetSum, ans)
			(*ans)[i] = (*ans)[i][:len((*ans)[i])-1]
			(*subSetSum)[i] -= given[currentIndex]
		} else if (*ans)[i] == nil {
			(*ans)[i] = append((*ans)[i], given[currentIndex])
			(*subSetSum)[i] += given[currentIndex]
			helpSumPartion(given, currentIndex+1, n, k, subSetSum, ans)
			(*ans)[i] = (*ans)[i][:len((*ans)[i])-1]
			(*subSetSum)[i] -= given[currentIndex]
			break
		}
	}
}
