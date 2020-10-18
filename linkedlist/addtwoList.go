package linkedlist

import (
	"strconv"
)

//AddTwo : It will add the two given linked list and retun the new linked list with sum
func AddTwo(l1 *Linkedlist, l2 *Linkedlist) *Linkedlist {
	ans := &Linkedlist{}
	carry := 0
	l1.Reverse()
	l2.Reverse()
	first := l1.head
	second := l2.head

	for {
		if first == nil || second == nil {
			break
		}
		temp1, _ := strconv.Atoi(first.value)
		temp2, _ := strconv.Atoi(second.value)

		sum := carry + temp1 + temp2
		if sum >= 10 {
			data := sum % 10
			ans.Insert(strconv.Itoa(data))
			carry = 1
		} else {
			ans.Insert(strconv.Itoa(sum))
			carry = 0
		}
		first = first.next
		second = second.next
	}

	for {

		if second == nil {
			break
		}
		temp, _ := strconv.Atoi(second.value)
		sum := carry + temp
		if sum >= 10 {
			data := sum % 10
			ans.Insert(strconv.Itoa(data))
			carry = 1
		} else {
			ans.Insert(strconv.Itoa(sum))
			carry = 0
		}
		second = second.next
	}
	for {
		if first == nil {
			break
		}
		temp, _ := strconv.Atoi(first.value)
		sum := carry + temp
		if sum >= 10 {
			data := sum % 10
			ans.Insert(strconv.Itoa(data))
			carry = 1
		} else {
			ans.Insert(strconv.Itoa(sum))
			carry = 0
		}
		first = first.next
	}
	if carry != 0 {
		ans.Insert("1")
	}

	ans.Reverse()

	return ans

}
