package linkedlist

//LastToFront : it moves the element from last to the front
func (l *Linkedlist) LastToFront() {
	var prev *Node
	second := l.head
	i := 0
	for {
		println("The value of i is ", i)
		if second.next == nil {
			break
		}
		prev = second
		second = second.next
		i++
	}
	second.next = l.head
	l.head = second
	prev.next = nil

}
