package main

import (
	dynamicprogramming "babbar/DynamicProgramming"
	"fmt"
)

func main() {
	arr := [][]int{{1, 2, 3},
		{4, 8, 2},
		{1, 5, 3}}
	res := dynamicprogramming.MinimumPath(arr, 2, 2)
	fmt.Println(res)

}
