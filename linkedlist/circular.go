package linkedlist

import "fmt"

// MakeCircular : To make the linked list circular
func (l *Linkedlist) MakeCircular() {
	current := l.head
	for {
		if current.next == nil {
			current.next = l.head
			break
		}

		current = current.next
	}
}

// CircularDelete : To delete the linked list from circular
func (l *Linkedlist) CircularDelete(s string) {
	current := l.head
	if current.value == s {
		println("First condition is owkring")
		temp := l.head
		for {
			if temp.next == l.head {
				temp.next = temp.next.next
				l.head = temp.next.next
				current.next = nil
				return
			}

			temp = temp.next

		}
	}

	fmt.Println("At the starting the value is ", current.value)
	for {

		if current.next == l.head || current.next.value == s {
			break
		}
		current = current.next
	}
	fmt.Println("at present the current is ", current.value)
	if current.next.value == s {
		temp := current.next
		current.next = temp.next
	} else {
		println("Value not found in the database")
	}
}
