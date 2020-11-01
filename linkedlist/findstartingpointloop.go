package linkedlist

//StartingPoint : it gives the starting point of the loop in the linked list
func (l *Linkedlist) StartingPoint() interface{} {
	slow := l.Head
	fast := l.Head

	for {
		if fast == nil || fast.Next == nil {
			return "There is no loop detected in the linked list"
		}
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			slow := l.Head
			for slow.Next != fast.Next {
				slow = slow.Next
				fast = fast.Next
			}

			ans := fast.Next.Value

			return ans
		}
	}
}
