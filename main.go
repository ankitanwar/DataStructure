package main

import (
	"babbar/heaps"
	"fmt"
)

func main() {
	arr := []int{8, 5, 10, 7, 9, 4, 15, 12, 90, 13}
	first := heaps.MaxOfAllSubArrays(arr, 4)
	fmt.Printf("%v ", first)
}
