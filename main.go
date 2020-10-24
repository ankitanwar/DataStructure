package main

import (
	stackandqueues "babbar/stackAndQueues"
	"fmt"
)

func main() {
	s := stackandqueues.StackUsingDQueues{}
	s.Insert("1")
	s.Insert("2")
	s.Insert("3")
	s.Insert("4")
	s.Insert("5")
	ans := s.Delete()
	fmt.Println(ans)

}
