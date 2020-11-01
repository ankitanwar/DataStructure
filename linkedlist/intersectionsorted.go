package linkedlist

//IntersectionS : it gives the intersection of the two given linked list
func (l *Linkedlist) IntersectionS(l2 *Linkedlist) *Linkedlist {
	first := l.Head
	second := l2.Head
	ans := &Linkedlist{}

	for {
		if second == nil || first == nil {
			break
		} else if second.Value == first.Value {
			ans.Insert(first.Value)
			first = first.Next
			second = second.Next
		} else if first.Value.(int) > second.Value.(int) {
			second = second.Next
		} else {
			first = first.Next
		}
	}

	return ans

}
