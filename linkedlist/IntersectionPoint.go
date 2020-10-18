package linkedlist

//IntersectionP : it will return the intersection point in the linked list
func (l *Linkedlist) IntersectionP(l2 *Linkedlist) string {

	println("The intersection point is function is working")
	count1 := 0
	count2 := 0

	first := l.head
	second := l2.head

	for {
		if first == nil {
			break
		}
		first = first.next
		count1++
	}
	for {
		if second == nil {
			break
		}
		second = second.next
		count2++
	}

	first = l.head
	second = l2.head

	if count1 > count2 {
		diff := count1 - count2
		for i := 0; i < diff; i++ {
			first = first.next
		}
	} else if count2 > count1 {
		diff := count2 - count1
		for i := 0; i < diff; i++ {
			second = second.next
		}
	}

	for {
		if first == second {
			return first.next.value
		}
		if first.next == nil || second.next == nil {
			return "There is no Intersection"
		}

		first = first.next
		second = second.next
	}

}

// MakeIntersection : This will make the intersection between the two linked list
func (l *Linkedlist) MakeIntersection(l2 *Linkedlist) string {
	first := l.head.next.next.next

	second := l2.head
	if second == nil {
		return "Cannot make intersection with  an empty linked list"
	}

	for {
		if second.next == nil {
			break
		}

		second = second.next
	}

	second.next = first
	return "Intersection has been made successfully"

}
