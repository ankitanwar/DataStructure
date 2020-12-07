package main

import (
	"fmt"

	dynamicprogramming "github.com/ankitanwar/Golang-DataStructure/DynamicProgramming"
)

func main() {
	given := "aab"
	ans := dynamicprogramming.PailndromicPartionDp(given, 0, len(given)-1)
	fmt.Println(ans)
}
