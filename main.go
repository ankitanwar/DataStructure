package main

import (
	stackandqueues "babbar/stackandqueues"
	"fmt"
)

func main() {
	s := stackandqueues.Midstack{}
	s.Insert("1")
	s.Insert("2")
	s.Insert("3")
	s.Insert("4")
	s.Insert("5")
	fmt.Println(s.GetMiddle())
	s.Print()
	s.Delete()
	println("x-x-x-x-x-x-x--x")
	fmt.Println(s.GetMiddle())
	s.Print()
	s.Delete()
	s.Delete()
	println("x-x-x-x-x-x-x--x")
	fmt.Println(s.GetMiddle())
	s.Print()

}
