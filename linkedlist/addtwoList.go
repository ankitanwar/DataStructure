package linkedlist

//AddTwo : It will add the two given linked list and retun the new linked list with sum
func AddTwo(l1 *Linkedlist, l2 *Linkedlist) *Linkedlist {
	ans := &Linkedlist{}
	carry := 0
	l1.Reverse()
	l2.Reverse()
	first := l1.Head
	second := l2.Head

	for {
		if first == nil || second == nil {
			break
		}

		sum := carry + first.Value.(int) + second.Value.(int)
		if sum >= 10 {
			data := sum % 10
			ans.Insert(data)
			carry = 1
		} else {
			ans.Insert(sum)
			carry = 0
		}
		first = first.Next
		second = second.Next
	}

	for {

		if second == nil {
			break
		}
		sum := carry + second.Value.(int)
		if sum >= 10 {
			data := sum % 10
			ans.Insert(data)
			carry = 1
		} else {
			ans.Insert(sum)
			carry = 0
		}
		second = second.Next
	}
	for {
		if first == nil {
			break
		}
		sum := carry + first.Value.(int)
		if sum >= 10 {
			data := sum % 10
			ans.Insert(data)
			carry = 1
		} else {
			ans.Insert(sum)
			carry = 0
		}
		first = first.Next
	}
	if carry != 0 {
		ans.Insert("1")
	}

	ans.Reverse()

	return ans

}
