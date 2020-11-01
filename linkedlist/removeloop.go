package linkedlist

// Removeloop : It its use to remove loop from the linked list
func (l *Linkedlist) Removeloop() string {
	slow := l.Head
	fast := l.Head

	for {

		if fast == nil || fast.Next == nil {
			return "There is no loop in the linked list"
		}
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			slow = l.Head
			for slow.Next != fast.Next {
				slow = slow.Next
				fast = fast.Next

			}
			fast.Next = nil

			return "Loop has been removed from the linked list successfully"
		}
	}

}
