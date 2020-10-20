package linkedlist

// Divide : It will divide the circular linked list into two equal half
func (l *Linkedlist) Divide() *Linkedlist {
	slow := l.head
	fast := l.head
	second := &Linkedlist{}
	var newHead *Node

	for true {
		fast = fast.next.next
		if fast == l.head {
			newHead = slow.next
			break
		} else if fast.next == l.head {
			newHead = slow.next.next
			break
		}

		slow = slow.next
	}
	slow = l.head
	for {
		if slow.next == newHead {
			slow.next = l.head
			break
		}
		slow = slow.next
	}
	second.head = newHead
	for {
		if newHead.next == l.head {
			newHead.next = second.head
			break
		}
		newHead = newHead.next
	}

	return second

}
