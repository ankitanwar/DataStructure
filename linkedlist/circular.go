package linkedlist

import "fmt"

// MakeCircular : To make the linked list circular
func (l *Linkedlist) MakeCircular() {
	current := l.Head
	for {
		if current.Next == nil {
			current.Next = l.Head
			break
		}

		current = current.Next
	}
}

// CircularDelete : To delete the linked list from circular
func (l *Linkedlist) CircularDelete(s string) {
	current := l.Head
	if current.Value == s {
		temp := l.Head
		for {
			if temp.Next == l.Head {
				temp.Next = temp.Next.Next
				l.Head = temp.Next.Next
				current.Next = nil
				return
			}

			temp = temp.Next

		}
	}

	fmt.Println("At the starting the Value is ", current.Value)
	for {

		if current.Next == l.Head || current.Next.Value == s {
			break
		}
		current = current.Next
	}
	fmt.Println("at present the current is ", current.Value)
	if current.Next.Value == s {
		temp := current.Next
		current.Next = temp.Next
	} else {
		println("Value not found in the database")
	}
}
