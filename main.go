package main

import (
	dynamicprogramming "babbar/DynamicProgramming"
	"fmt"
)

func main() {
	rodlength := 17
	arr := []int{10, 11, 3}
	res := dynamicprogramming.MaximizeCut2(rodlength, arr)
	fmt.Println(res)

}
