package linkedlist

// RemoveDuplicatesUnsorted : It will remove the duplicates from the unsorted list
func (l *Linkedlist) RemoveDuplicatesUnsorted() {
	myDict := make(map[string]int)
	slow := l.head
	fast := l.head.next
	i := 1
	myDict[slow.value] = i
	for {
		if fast == nil {
			break
		}
		_, found := myDict[fast.value]
		if found {
			fast = fast.next
		} else {
			myDict[fast.value] = i
			slow.next = fast
			slow = fast
			fast = fast.next
		}
	}
	slow.next = fast

}
