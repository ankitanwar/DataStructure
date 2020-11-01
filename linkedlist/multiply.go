package linkedlist

// Multiply : It will multiply the two given linked list
func (l *Linkedlist) Multiply(l2 *Linkedlist) int {
	var num1 int
	var num2 int

	num1 = 0
	num2 = 0

	current := l.Head
	for {
		if current == nil {
			break
		}
		num1 = num1*10 + current.Value.(int)
		current = current.Next
	}

	current = l2.Head
	for {
		if current == nil {
			break
		}
		num2 = num2*10 + current.Value.(int)
		current = current.Next
	}

	return num1 * num2

}
