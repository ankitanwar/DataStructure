package linkedlist

import "strconv"

// Delnodes : It will deleted those nodes whose value is greater
func (l *Linkedlist) Delnodes() {
	l.Reverse()
	current := l.head
	max := current.value

	for {
		if current == nil || current.next == nil {
			break
		}
		temp1, _ := strconv.Atoi(max)
		temp2, _ := strconv.Atoi(current.next.value)

		if temp2 <= temp1 {
			temp := current.next
			current.next = temp.next
			temp.next = nil
		} else {
			current = current.next
			max = current.value
		}

	}
	println("Original list")
	l.Reverse()
	l.Print()
}
