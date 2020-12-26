package bit

import (
	"fmt"
	"sort"
)

//MinXorPair : To calc the all the pairs whose xor pair is the minimum
func MinXorPair(arr []int) {
	sort.Ints(arr)
	min := 1e6
	ans := [][]int{}
	for i := 0; i < len(arr)-1; i++ {
		t := arr[i] ^ arr[i+1]
		if float64(t) < min {
			min = float64(t)
			ans = ans[len(ans):]
			t1 := []int{arr[i], arr[i+1]}
			ans = append(ans, t1)
		} else if float64(t) == min {
			t2 := []int{arr[i], arr[i+1]}
			ans = append(ans, t2)
		}
	}
	fmt.Println(ans)
}
