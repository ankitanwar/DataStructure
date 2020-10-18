package linkedlist

//Reverse : This will reverse the linked list
func (l *Linkedlist) Reverse() {
	var prev *Node
	current := l.head

	for {
		if current.next == nil {
			current.next = prev
			l.head = current
			break
		}
		nxt := current.next
		current.next = prev
		prev = current
		current = nxt
	}

}
