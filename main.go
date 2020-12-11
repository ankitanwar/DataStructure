package main

import (
	dynamicprogramming "github.com/ankitanwar/Golang-DataStructure/DynamicProgramming"
)

func main() {
	s := []int{1, 3, 6, 2}
	e := []int{2, 5, 19, 100}
	p := []int{50, 20, 100, 200}
	dynamicprogramming.WeightJobScheduling(e, s, p, 0)
}
