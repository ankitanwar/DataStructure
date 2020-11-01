package linkedlist

// Divide : It will divide the circular linked list into two equal half
func (l *Linkedlist) Divide() *Linkedlist {
	slow := l.Head
	fast := l.Head
	second := &Linkedlist{}
	var newHead *Node

	for true {
		fast = fast.Next.Next
		if fast == l.Head {
			newHead = slow.Next
			break
		} else if fast.Next == l.Head {
			newHead = slow.Next.Next
			break
		}

		slow = slow.Next
	}
	slow = l.Head
	for {
		if slow.Next == newHead {
			slow.Next = l.Head
			break
		}
		slow = slow.Next
	}
	second.Head = newHead
	for {
		if newHead.Next == l.Head {
			newHead.Next = second.Head
			break
		}
		newHead = newHead.Next
	}

	return second

}
