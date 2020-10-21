package linkedlist

import "strconv"

// Multiply : It will multiply the two given linked list
func (l *Linkedlist) Multiply(l2 *Linkedlist) int {
	var num1 int
	var num2 int

	num1 = 0
	num2 = 0

	current := l.head
	for {
		if current == nil {
			break
		}
		temp, _ := strconv.Atoi(current.value)
		num1 = num1*10 + temp
		current = current.next
	}

	current = l2.head
	for {
		if current == nil {
			break
		}
		temp, _ := strconv.Atoi(current.value)
		num2 = num2*10 + temp
		current = current.next
	}

	return num1 * num2

}
