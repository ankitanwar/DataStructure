package linkedlist

//Middle : It find the middle element of the linked list
func (l *Linkedlist) Middle() string {
	slow := l.head
	fast := l.head

	// it means only one element is present in the list
	if fast.next == nil {
		return fast.value
	}

	for true {
		fast = fast.next.next
		if fast == nil || fast.next == nil {
			return slow.next.value
		}

		slow = slow.next
	}
	return ""

}
