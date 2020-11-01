package linkedlist

// Delnodes : It will deleted those nodes whose Value is greater
func (l *Linkedlist) Delnodes() {
	l.Reverse()
	current := l.Head
	max := current.Value

	for {
		if current == nil || current.Next == nil {
			break
		}

		if current.Next.Value.(int) <= max.(int) {
			temp := current.Next
			current.Next = temp.Next
			temp.Next = nil
		} else {
			current = current.Next
			max = current.Value
		}

	}
	println("Original list")
	l.Reverse()
	l.Print()
}
