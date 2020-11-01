package linkedlist

import "fmt"

// Pairwise : It will swap the linked list pair wise
func (l *Linkedlist) Pairwise(node *Node) {
	current := node
	if current == nil {
		return
	}
	if current.Next == nil {
		return
	}
	l.Head = current.Next
	for {
		fmt.Println("the value of current is ", current.Value.(string))
		q := current.Next
		temp := q.Next
		q.Next = current

		if temp == nil || temp.Next == nil {
			current.Next = temp
			return
		}
		current.Next = temp.Next
		current = temp
	}
}
