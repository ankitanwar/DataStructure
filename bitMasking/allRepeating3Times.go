package bit

import "fmt"

//AllRepeating3Times : Given an array in which every element occurs three times except 1 we need to find out this element
func AllRepeating3Times(arr []int) {
	t2 := arr
	t3 := arr
	ans := 0
	for i := 0; i < len(arr); i++ {
		ans = ans ^ t2[i] ^ t3[i] ^ arr[i]
	}
	fmt.Println(ans)
}
