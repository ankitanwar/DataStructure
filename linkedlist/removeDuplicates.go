package linkedlist

//RemoveDuplicatesSorted : It removes the duplicate in the sorted linked list
func (l *Linkedlist) RemoveDuplicatesSorted() string {
	slow := l.head
	if slow.next == nil {
		return "There is no duplicates element in the linked list"
	}
	fast := l.head.next

	for {
		if fast == nil {
			break
		}
		if slow.value == fast.value {
			fast = fast.next
		} else {
			slow.next = fast
			slow = fast
			fast = fast.next
		}
	}
	slow.next = fast

	return "All the duplicates has been removed successfully"
}
