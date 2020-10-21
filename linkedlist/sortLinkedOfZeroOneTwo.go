package linkedlist

import "strconv"

// Sort : This will sort the linked list of 0,1 and 2
func (l *Linkedlist) Sort() {
	count := [3]int{}

	current := l.head

	for {

		if current == nil {
			break
		}
		temp1, _ := strconv.Atoi(current.value)
		count[temp1]++
		current = current.next
	}

	current = l.head
	i := 0
	for {
		if i > 2 || current == nil {
			break
		}
		if count[i] == 0 {
			i++
		}
		temp := strconv.Itoa(i)
		current.value = temp
		current = current.next
		count[i]--
	}
}
