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

//ReverseRecursive : This function will recursively Reverse the linked list
func (l *Linkedlist) ReverseRecursive(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.next == nil {
		l.head = node
		return node
	}
	var p *Node
	p = node.next
	l.ReverseRecursive(p)

	p.next = node
	node.next = nil
	return node

}
