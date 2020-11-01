package linkedlist

//Reverse : This will reverse the linked list
func (l *Linkedlist) Reverse() {
	var prev *Node
	current := l.Head

	for {
		if current.Next == nil {
			current.Next = prev
			l.Head = current
			break
		}
		nxt := current.Next
		current.Next = prev
		prev = current
		current = nxt
	}

}

//ReverseRecursive : This function will recursively Reverse the linked list
func (l *Linkedlist) ReverseRecursive(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.Next == nil {
		l.Head = node
		return node
	}
	var p *Node
	p = node.Next
	l.ReverseRecursive(p)

	p.Next = node
	node.Next = nil
	return node

}
