package backtracking

import "fmt"

//EqualSumPartion : To partion the given array into k parts such that there sum should me equal
func EqualSumPartion(given []int, k int) {
	partion := make([][]int, k)
	helpSumPartion(given, 0, 0, k, &partion)

}

func helpSumPartion(given []int, currentIndex, currentPartion, k int, partion *[][]int) {
	if currentIndex == len(given) {
		fmt.Print("The value of currentPation and k is ", currentPartion, k)
		println()
		if currentPartion == k {
			var equal bool
			var sum int
			println("This condition is true", sum)
			for i := 0; i < len((*partion)[0]); i++ {
				sum += (*partion)[0][i]

			}
			for i := 1; i < len(*(partion)); i++ {
				curr := (*partion)[i]
				var temp int
				for j := 0; j < len(curr); j++ {
					temp += curr[j]
				}
				if temp != sum {
					equal = true
					break
				}
			}
			if equal == false {
				for i := 0; i < len(*(partion)); i++ {
					fmt.Printf("%v ", (*partion)[i])
				}
				println("")
			}
		}
		return
	}
	for i := 0; i < len(*partion); i++ {
		if (*partion)[i] != nil {
			(*partion)[i] = append((*partion)[i], given[currentIndex])
			helpSumPartion(given, currentIndex+1, currentPartion, k, partion)
			(*partion)[i] = (*partion)[i][:len((*partion)[i])-1]
		} else {
			(*partion)[i] = append((*partion)[i], given[currentIndex])
			helpSumPartion(given, currentIndex+1, currentPartion+1, k, partion)
			(*partion)[i] = (*partion)[i][:len((*partion)[i])-1]
			break
		}
	}
}
