package linkedlist

//Pailndrome : It will check whether the given linked list is Pailndrome or not
func (l Linkedlist) Pailndrome() bool {
	var ans bool
	slow := l.head
	fast := l.head
	var newHead *Node
	for true {
		fast = fast.next.next
		if fast == nil {
			newHead = slow.next
			break
		}

		if fast.next == nil {
			newHead = slow.next.next
			break
		}

		slow = slow.next
	}

	println("in starting the linked list is ")
	l.Print()
	var prev *Node
	current := newHead
	for {
		if current.next == nil {
			current.next = prev
			newHead = current
			break
		}

		nxt := current.next
		current.next = prev
		prev = current
		current = nxt
	}

	println("In the mid the linked list is ")
	l.Print()

	slow.next = newHead
	println("The final linked list is ")
	l.Print()
	slow = l.head
	for {
		if newHead == nil {
			ans = true
			return ans
		}

		if slow.value != newHead.value {
			ans = false
			return ans
		}
		slow = slow.next
		newHead = newHead.next
	}

}
