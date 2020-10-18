package linkedlist

//DetectLoop : To detect wherther there is an loop in the linked list or not
func (l *Linkedlist) DetectLoop() string {
	slow := l.head
	fast := l.head

	for {
		if fast == nil || fast.next == nil || fast.next.next == nil {
			return "There is no loop detcted"
		}
		fast = fast.next.next
		slow = slow.next

		if fast == slow {
			return "There is an loop in the linked list"
		}
	}
}

// Makeloop : it is used to make the loop in the linked list
func (l *Linkedlist) Makeloop() string {
	first := l.head.next.next
	second := l.head

	for {
		if second.next == nil {
			break
		}

		second = second.next
	}

	second.next = first

	return "loop has been created successfully"
}
