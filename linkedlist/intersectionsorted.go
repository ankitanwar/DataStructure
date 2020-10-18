package linkedlist

//IntersectionS : it gives the intersection of the two given linked list
func (l *Linkedlist) IntersectionS(l2 *Linkedlist) *Linkedlist {
	first := l.head
	second := l2.head
	ans := &Linkedlist{}

	for {
		if second == nil || first == nil {
			break
		} else if second.value == first.value {
			ans.Insert(first.value)
			first = first.next
			second = second.next
		} else if first.value > second.value {
			second = second.next
		} else {
			first = first.next
		}
	}

	return ans

}
