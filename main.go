package main

import (
	dynamicprogramming "babbar/DynamicProgramming"
	"fmt"
)

func main() {

	arr := []int{40, 20, 30, 10, 30}
	fmt.Println(dynamicprogramming.MatrixChainMultiplicationDp(arr))

}
