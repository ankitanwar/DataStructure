package sorting

import "fmt"

//FourSum : returns the value of 4 numbers whose sum is equal to the target
func FourSum(arr []int, target int) []int {
	myDict := make(map[int][]int)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			add := arr[i] + arr[j]
			myDict[add] = []int{arr[i], arr[j]}
		}
	}

	for key, value := range myDict {
		fmt.Println("The key and value pair is ", key, value)
		find := target - key
		_, found := myDict[find]
		if found {
			ans := []int{}
			ans = append(value, myDict[find]...)
			return ans
		}
	}

	return []int{-1, -1, -1, -1}
}
