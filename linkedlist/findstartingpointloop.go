package linkedlist

//StartingPoint : it gives the starting point of the loop in the linked list
func (l *Linkedlist) StartingPoint() string {
	slow := l.head
	fast := l.head

	for {
		if fast == nil || fast.next == nil {
			return "There is no loop detected in the linked list"
		}
		fast = fast.next.next
		slow = slow.next

		if fast == slow {
			slow := l.head
			for slow.next != fast.next {
				slow = slow.next
				fast = fast.next
			}

			ans := fast.next.value

			return ans
		}
	}
}
