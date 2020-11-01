package linkedlist

//Add1 : It will add 1 to the linked list
func (l *Linkedlist) Add1(num int) {
	l.Reverse()
	carry := 0
	current := l.Head
	sum := current.Value.(int) + num

	if sum >= 10 {
		data := sum % 10
		current.Value = data
		carry = 1
		current = current.Next
	} else {
		current.Value = sum
		current = current.Next
	}

	for {

		if current == nil {
			break
		}
		sum := carry + current.Value.(int)

		if sum >= 10 {
			data := sum % 10
			current.Value = data
			carry = 1
		} else {
			current.Value = sum
			carry = 0
		}

		current = current.Next

	}
	if carry != 0 {
		l.Insert("1")
	}
	l.Reverse()

}
