package linkedlist

import (
	"strconv"
)

//Add1 : It will add 1 to the linked list
func (l *Linkedlist) Add1(num int) {
	l.Reverse()
	carry := 0
	current := l.head
	temp, _ := strconv.Atoi(current.value)
	sum := temp + num

	if sum >= 10 {
		data := sum % 10
		current.value = strconv.Itoa(data)
		carry = 1
		current = current.next
	} else {
		current.value = strconv.Itoa(sum)
		current = current.next
	}

	for {

		if current == nil {
			break
		}
		temp, _ := strconv.Atoi(current.value)
		sum := temp + carry

		if sum >= 10 {
			data := sum % 10
			current.value = strconv.Itoa(data)
			carry = 1
		} else {
			current.value = strconv.Itoa(sum)
			carry = 0
		}

		current = current.next

	}
	if carry != 0 {
		l.Insert("1")
	}
	l.Reverse()

}
