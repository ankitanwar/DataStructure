package main

import (
	"fmt"

	dynamicprogramming "github.com/ankitanwar/Golang-DataStructure/DynamicProgramming"
)

func main() {
	arr := []int{2, 2, 2}
	ans := dynamicprogramming.Subset(arr, 6, 3)
	fmt.Println(ans)

}
