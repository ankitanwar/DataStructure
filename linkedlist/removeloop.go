package linkedlist

// Removeloop : It its use to remove loop from the linked list
func (l *Linkedlist) Removeloop() string {
	slow := l.head
	fast := l.head

	for {

		if fast == nil || fast.next == nil {
			return "There is no loop in the linked list"
		}
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			slow = l.head
			for slow.next != fast.next {
				slow = slow.next
				fast = fast.next

			}
			fast.next = nil

			return "Loop has been removed from the linked list successfully"
		}
	}

}
