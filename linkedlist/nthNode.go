package linkedlist

//Nthnode : It will return the nth node from the end
func (l *Linkedlist) Nthnode(n int) string {
	current := l.Head
	fast := l.Head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for {
		if fast == nil {
			return current.Value.(string)
		}
		current = current.Next
		fast = fast.Next
	}
}
