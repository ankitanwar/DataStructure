package main

import (
	dynamicprogramming "babbar/DynamicProgramming"
	"fmt"
)

func main() {
	arr := []int{334, 500, 169, 724, 478, 358, 962, 464, 705, 145, 281, 827, 961, 491, 995, 942, 827}
	ans := dynamicprogramming.MaximumProfit(arr)
	fmt.Println(ans)

}
