package linkedlist

//Middle : It find the middle element of the linked list
func (l *Linkedlist) Middle() interface{} {
	slow := l.Head
	fast := l.Head

	// it means only one element is present in the list
	if fast.Next == nil {
		return fast.Value
	}

	for true {
		fast = fast.Next.Next
		if fast == nil || fast.Next == nil {
			return slow.Next.Value
		}

		slow = slow.Next
	}
	return ""

}
