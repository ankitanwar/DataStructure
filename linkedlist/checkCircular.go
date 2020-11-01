package linkedlist

// CheckCircular : To check whether the given list is circular or not
func (l *Linkedlist) CheckCircular() bool {
	current := l.Head

	if current == nil {
		return true
	}

	node := current.Next

	for {
		if node == current {
			return true
		}
		if node == nil {
			return false
		}
		node = node.Next
	}
}

// Mcircular : to make the linked list circular
func (l *Linkedlist) Mcircular() {
	current := l.Head

	for {
		if current.Next == nil {
			break
		}
		current = current.Next

	}
	current.Next = l.Head
	println("We have made the linked list circular successfully")
}
