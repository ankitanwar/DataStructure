package backtracking

import "fmt"

var miin int = 9999999
var ans1 []int
var ans2 []int

//TugOfWars : To divide the gievn array in two arrays first and second such that the difference between the first array and the second array should me minimum
func TugOfWars(arr []int) (ans1, ans2 []int) {
	first := []int{}
	second := []int{}
	tugHelper(arr, first, second, 0, 0, 0)
	return ans1, ans2
}

func tugHelper(given, first, second []int, currentIndex, sumOfFirst, sumOfsecond int) {
	if currentIndex == len(given) {
		del := abs(sumOfFirst - sumOfsecond)
		if del < miin {
			miin = del
			ans1 = first
			ans2 = second
			fmt.Println(ans1, ans2)
		}
		return
	}
	if len(first) <= (len(given)+1)/2 {
		first = append(first, given[currentIndex])
		tugHelper(given, first, second, currentIndex+1, sumOfFirst+given[currentIndex], sumOfsecond)
		first = first[:len(first)-1]
	}
	if len(second) <= (len(given)+1)/2 {
		second = append(second, given[currentIndex])
		tugHelper(given, second, second, currentIndex+1, sumOfsecond, sumOfsecond+given[currentIndex])
		second = second[:len(second)-1]
	}
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
