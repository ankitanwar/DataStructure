package linkedlist

//LastToFront : it moves the element from last to the front
func (l *Linkedlist) LastToFront() {
	var prev *Node
	second := l.Head
	i := 0
	for {
		println("The value of i is ", i)
		if second.Next == nil {
			break
		}
		prev = second
		second = second.Next
		i++
	}
	second.Next = l.Head
	l.Head = second
	prev.Next = nil

}
