package linkedlist

import (
	"fmt"
	"strconv"
)

// Segrate : It will segrate the even and the odd linked list
func (l *Linkedlist) Segrate() *Linkedlist {
	var eEnd *Node
	var oEnd *Node
	evenlist := &Linkedlist{}
	oddlist := &Linkedlist{}

	current := l.head

	for {
		if current == nil {
			break
		}

		temp, _ := strconv.Atoi(current.value)
		if temp%2 == 0 {
			if evenlist.head == nil {
				evenlist.head = current
				eEnd = evenlist.head
			} else {
				eEnd.next = current
				eEnd = eEnd.next
			}
		} else {
			if oddlist.head == nil {
				oddlist.head = current
				oEnd = oddlist.head
			} else {
				oEnd.next = current
				oEnd = oEnd.next
			}
		}

		current = current.next

	}
	fmt.Println(eEnd.value, oEnd.value)
	if eEnd == nil || oEnd == nil {
		return l
	}

	eEnd.next = oddlist.head
	oEnd.next = nil
	evenlist.Print()
	return evenlist

}
