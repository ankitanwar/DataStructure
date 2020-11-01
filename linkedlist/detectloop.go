package linkedlist

//DetectLoop : To detect wherther there is an loop in the linked list or not
func (l *Linkedlist) DetectLoop() string {
	slow := l.Head
	fast := l.Head

	for {
		if fast == nil || fast.Next == nil || fast.Next.Next == nil {
			return "There is no loop detcted"
		}
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return "There is an loop in the linked list"
		}
	}
}

// Makeloop : it is used to make the loop in the linked list
func (l *Linkedlist) Makeloop() string {
	first := l.Head.Next.Next
	second := l.Head

	for {
		if second.Next == nil {
			break
		}

		second = second.Next
	}

	second.Next = first

	return "loop has been created successfully"
}
