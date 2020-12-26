package bit

import (
	"fmt"
)

//AllRepeatingExcept1 : we need to find out the number in an array which is not repeating
func AllRepeatingExcept1(arr []int) {
	t := 0
	for i := 0; i < len(arr); i++ {
		t = arr[i] ^ t

	}
	fmt.Println(t)
}
