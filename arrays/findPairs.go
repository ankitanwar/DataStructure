package array

import "fmt"

//FindPairs : Given an array we need to find the pairs whose sum is equal to the given target
func FindPairs(arr []int, target int) {
	dict := make(map[int]int)
	var ans int
	for i := 0; i < len(arr); i++ {
		_, found := dict[target-arr[i]]
		if found {
			ans += dict[target-arr[i]]
			_, f := dict[arr[i]]
			if f {
				dict[arr[i]]++
			} else {
				dict[arr[i]] = 1
			}
		} else {
			_, f := dict[arr[i]]
			if f {
				dict[arr[i]]++
			} else {
				dict[arr[i]] = 1
			}
		}
	}
	fmt.Println(ans)
}
