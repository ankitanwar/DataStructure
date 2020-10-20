package main

import (
	"babbar/linkedlist"
)

func main() {
	list := &linkedlist.Linkedlist{}
	// list2 := &linkedlist.Linkedlist{}
	list.Insert("1")
	head := list.GetHead()
	list.Insert("2")
	list.Insert("3")
	list.Insert("4")
	list.Insert("5")
	list.Pairwise(head)
	list.Print()

}
