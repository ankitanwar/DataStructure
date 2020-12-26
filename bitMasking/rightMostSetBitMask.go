package bit

import (
	"fmt"
)

//RightMostSetBit : Given a number we need to find its right most bit mask
func RightMostSetBit(n int) {
	/*
		Explanation - Take the given number and flip all the 0 and 1 and then add 1 to it (2s compliment basically) and then & with the original number
	*/
	t := n & -n
	fmt.Printf("%b", t)
}

//example 76 --> Binary --> 100100
//RightMostBitMask -> 000100 means except the right most 1 everything will be zero
