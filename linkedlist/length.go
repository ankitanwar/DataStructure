package linkedlist

// Length : It calculate the length of the linked list
func (l *Linkedlist) Length() {
	count := 0
	current := l.head
	for {
		if current == nil {
			break
		}
		count++
		current = current.next
	}

	println("The length of the linked list is ", count)
}
