package bit

import "fmt"

//PositionOfSetBits : Given an integer we need to calc its set bit postion and if there are more than 1 or zero set bits we need to print -1
func PositionOfSetBits(n int) {
	if n&(n-1) != 0 {
		fmt.Println(-1)
	} else {
		pos := 0
		for {
			t := 1 << pos
			if t&n != 0 {
				break
			}
			pos++
		}
		fmt.Println(pos + 1)
	}
}
