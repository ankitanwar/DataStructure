package linkedlist

import "strconv"

// Sort : This will sort the linked list of 0,1 and 2
func (l *Linkedlist) Sort() {
	count := [3]int{}

	current := l.Head

	for {

		if current == nil {
			break
		}
		count[current.Value.(int)]++
		current = current.Next
	}

	current = l.Head
	i := 0
	for {
		if i > 2 || current == nil {
			break
		}
		if count[i] == 0 {
			i++
		}
		temp := strconv.Itoa(i)
		current.Value = temp
		current = current.Next
		count[i]--
	}
}
