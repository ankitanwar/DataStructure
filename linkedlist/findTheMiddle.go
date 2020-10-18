package linkedlist

//Middle : It find the middle element of the linked list
func (l *Linkedlist) Middle() string {
	slow := l.head
	fast := l.head

	println("The value of the slow and fast is ", slow.value, fast)

	for {
		if fast.next == nil || fast.next.next == nil {
			return slow.value
		}

		slow = slow.next
		fast = fast.next.next
	}
}
