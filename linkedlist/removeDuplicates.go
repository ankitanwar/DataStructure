package linkedlist

//RemoveDuplicatesSorted : It removes the duplicate in the sorted linked list
func (l *Linkedlist) RemoveDuplicatesSorted() string {
	slow := l.Head
	if slow.Next == nil {
		return "There is no duplicates element in the linked list"
	}
	fast := l.Head.Next

	for {
		if fast == nil {
			break
		}
		if slow.Value == fast.Value {
			fast = fast.Next
		} else {
			slow.Next = fast
			slow = fast
			fast = fast.Next
		}
	}
	slow.Next = fast

	return "All the duplicates has been removed successfully"
}
