package bit

import "fmt"

//calc the number of bits which are 1 in the given integer

//Kernighan : To calc the number of bits which are 1 in he given integer
func Kernighan(n int) {
	var count int

	for {
		if n == 0 {
			break
		}
		rmb := n & -n
		n -= rmb
		count++
	}

	fmt.Println("The number of 1 in the given number is ", count)
}
