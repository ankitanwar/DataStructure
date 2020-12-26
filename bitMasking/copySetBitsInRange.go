package bit

import "fmt"

//CopySetBits : Given a number we need to tell the set the bit in number 2 from range left to right which are set in number1
func CopySetBits(number1, number2, left, right int) {
	mask1 := 0
	mask1 = mask1 | (1<<right - left + 1)
	mask1--
	mask1 <<= (left - 1)
	mask2 := number1 & mask1
	number2 = number2 | mask2
	fmt.Println(number2)

}
