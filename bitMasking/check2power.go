package bit

import "fmt"

//Check2Power : given a number we need to check wether the given number is 2 ki power or not
func Check2Power(n int) {
	if n&(n-1) == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
