package linkedlist

//IntersectionP : it will return the intersection point in the linked list
func (l *Linkedlist) IntersectionP(l2 *Linkedlist) string {

	println("The intersection point is function is working")
	count1 := 0
	count2 := 0

	first := l.Head
	second := l2.Head

	for {
		if first == nil {
			break
		}
		first = first.Next
		count1++
	}
	for {
		if second == nil {
			break
		}
		second = second.Next
		count2++
	}

	first = l.Head
	second = l2.Head

	if count1 > count2 {
		diff := count1 - count2
		for i := 0; i < diff; i++ {
			first = first.Next
		}
	} else if count2 > count1 {
		diff := count2 - count1
		for i := 0; i < diff; i++ {
			second = second.Next
		}
	}

	for {
		if first == second {
			return first.Next.Value.(string)
		}
		if first.Next == nil || second.Next == nil {
			return "There is no Intersection"
		}

		first = first.Next
		second = second.Next
	}

}

// MakeIntersection : This will make the intersection between the two linked list
func (l *Linkedlist) MakeIntersection(l2 *Linkedlist) string {
	first := l.Head.Next.Next.Next

	second := l2.Head
	if second == nil {
		return "Cannot make intersection with  an empty linked list"
	}

	for {
		if second.Next == nil {
			break
		}

		second = second.Next
	}

	second.Next = first
	return "Intersection has been made successfully"

}
