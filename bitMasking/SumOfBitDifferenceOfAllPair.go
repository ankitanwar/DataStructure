package bit

import "fmt"

//SumOfBitDifference : Given an array we need to calc the sum of bit difference in all number
func SumOfBitDifference(array []int) {
	var ans int
	for i := 31; i >= 0; i-- {
		var onCount int
		for j := 0; j < len(array); j++ {
			t := array[j] & (1 << i)
			if t != 0 {
				onCount++
			}
		}
		offCount := len(array) - onCount
		ans += 2 * onCount * offCount //a -->b and b-->a both are considered as different therefore we multiply the whole by 2
	}
	fmt.Println(ans)
}
