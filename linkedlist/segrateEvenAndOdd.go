package linkedlist

import (
	"fmt"
)

// Segrate : It will segrate the even and the odd linked list
func (l *Linkedlist) Segrate() *Linkedlist {
	var eEnd *Node
	var oEnd *Node
	evenlist := &Linkedlist{}
	oddlist := &Linkedlist{}

	current := l.Head

	for {
		if current == nil {
			break
		}
		if current.Value.(int)%2 == 0 {
			if evenlist.Head == nil {
				evenlist.Head = current
				eEnd = evenlist.Head
			} else {
				eEnd.Next = current
				eEnd = eEnd.Next
			}
		} else {
			if oddlist.Head == nil {
				oddlist.Head = current
				oEnd = oddlist.Head
			} else {
				oEnd.Next = current
				oEnd = oEnd.Next
			}
		}

		current = current.Next

	}
	fmt.Println(eEnd.Value, oEnd.Value)
	if eEnd == nil || oEnd == nil {
		return l
	}

	eEnd.Next = oddlist.Head
	oEnd.Next = nil
	evenlist.Print()
	return evenlist

}
