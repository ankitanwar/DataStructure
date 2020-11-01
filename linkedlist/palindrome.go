package linkedlist

//Pailndrome : It will check whether the given linked list is Pailndrome or not
func (l Linkedlist) Pailndrome() bool {
	var ans bool
	slow := l.Head
	fast := l.Head
	var newHead *Node
	for true {
		fast = fast.Next.Next
		if fast == nil {
			newHead = slow.Next
			break
		}

		if fast.Next == nil {
			newHead = slow.Next.Next
			break
		}

		slow = slow.Next
	}

	println("in starting the linked list is ")
	l.Print()
	var prev *Node
	current := newHead
	for {
		if current.Next == nil {
			current.Next = prev
			newHead = current
			break
		}

		nxt := current.Next
		current.Next = prev
		prev = current
		current = nxt
	}

	println("In the mid the linked list is ")
	l.Print()

	slow.Next = newHead
	println("The final linked list is ")
	l.Print()
	slow = l.Head
	for {
		if newHead == nil {
			ans = true
			return ans
		}

		if slow.Value.(string) != newHead.Value.(string) {
			ans = false
			return ans
		}
		slow = slow.Next
		newHead = newHead.Next
	}

}
