package main

import (
	"babbar/linkedlist"
	"fmt"
)

func main() {
	list := &linkedlist.Linkedlist{}
	//list2 := &linkedlist.Linkedlist{}
	list.Insert("1")
	// list.Insert("2")
	// list.Insert("3")
	// list.Insert("4")
	//list.Mcircular()
	fmt.Println(list.Middle())

}
