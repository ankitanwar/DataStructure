package linkedlist

import "fmt"

// Pairwise : It will swap the linked list pair wise
func (l *Linkedlist) Pairwise(node *Node) {
	current := node
	if current == nil {
		return
	}
	if current.next == nil {
		return
	}
	l.head = current.next
	for {
		fmt.Println("the value of current is ", current.value)
		q := current.next
		temp := q.next
		q.next = current

		if temp == nil || temp.next == nil {
			current.next = temp
			return
		}
		current.next = temp.next
		current = temp
	}
}
