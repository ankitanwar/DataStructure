package linkedlist

// CheckCircular : To check whether the given list is circular or not
func (l *Linkedlist) CheckCircular() bool {
	current := l.head

	if current == nil {
		return true
	}

	node := current.next

	for {
		if node == current {
			return true
		}
		if node == nil {
			return false
		}
		node = node.next
	}
}

// Mcircular : to make the linked list circular
func (l *Linkedlist) Mcircular() {
	current := l.head

	for {
		if current.next == nil {
			break
		}
		current = current.next

	}
	current.next = l.head
	println("We have made the linked list circular successfully")
}
