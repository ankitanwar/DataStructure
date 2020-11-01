package linkedlist

// RemoveDuplicatesUnsorted : It will remove the duplicates from the unsorted list
func (l *Linkedlist) RemoveDuplicatesUnsorted() {
	myDict := make(map[string]int)
	slow := l.Head
	fast := l.Head.Next
	i := 1
	myDict[slow.Value.(string)] = i
	for {
		if fast == nil {
			break
		}
		_, found := myDict[slow.Value.(string)]
		if found {
			fast = fast.Next
		} else {
			myDict[slow.Value.(string)] = i
			slow.Next = fast
			slow = fast
			fast = fast.Next
		}
	}
	slow.Next = fast

}
