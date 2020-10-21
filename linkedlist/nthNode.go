package linkedlist

//Nthnode : It will return the nth node from the end
func (l *Linkedlist) Nthnode(n int) string {
	current := l.head
	fast := l.head

	for i := 0; i < n; i++ {
		fast = fast.next
	}

	for {
		if fast == nil {
			return current.value
		}
		current = current.next
		fast = fast.next
	}
}
