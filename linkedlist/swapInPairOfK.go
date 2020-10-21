package linkedlist

import "fmt"

//ReverseN : this will reverse the node from the given pointer
func (l *Linkedlist) ReverseN(node *Node) (n *Node) {

	var prev *Node
	current := node

	if current == nil {
		return nil
	}

	for {

		if current.next == nil {
			current.next = prev
			l.head = prev
			return current
		}

		nxt := current.next
		current.next = prev
		prev = current
		current = nxt
	}
}

// ReverseInK : It will the given linked list in the size of k
func (l *Linkedlist) ReverseInK(value int) *Node {
	var newStart *Node
	var q *Node
	head := l.head
	p := first(head, value)

	newStart = p
	q = p
	for {
		p = l.head
		fmt.Println(q)
		break

	}

	return newStart

}

func first(node *Node, k int) *Node {
	current := node
	count := 0

	for {
		if count == k {
			return current
		}
		current = current.next
		count++
	}
}
