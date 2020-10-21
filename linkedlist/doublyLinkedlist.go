package linkedlist

//Dnode : node of a doubly linked list
type Dnode struct {
	prev  *Dnode
	next  *Dnode
	value string
}

// Dlinkedlist : doubly linked list
type Dlinkedlist struct {
	head *Dnode
}

// Insert : To insert the value in the doubly linked list
func (d *Dlinkedlist) Insert(val string) *Dnode {
	current := d.head
	n := &Dnode{}
	n.value = val
	if d.head == nil {
		d.head = n
		return d.head
	}
	for {
		if current.next == nil {
			current.next = n
			n.prev = current
			return n
		}
		current = current.next
	}
}

// Print : To The doubly linked list
func (d *Dlinkedlist) Print() {
	current := d.head
	ans := ""

	for {
		if current == nil {
			break
		}
		ans += current.value + "-->"
		current = current.next

	}
	println(ans)
}

//Ddlete : To delete the node in the doubly linked list
func (d *Dlinkedlist) Ddlete(s string) {
	current := d.head
	if current.value == s {
		temp := current.next
		d.head = temp
		current.next = nil
		return
	}

	for {
		if current.value == s {
			if current.next == nil {
				prev := current.prev
				prev.next = nil
				break

			} else {
				prev := current.prev
				nxt := current.next
				nxt.prev = current.prev.prev
				prev.next = current.next.next
				break
			}
		}

		current = current.next

	}

}

//Tsum : It will return the triplet sum
// func (d *Dlinkedlist) Tsum(target int) [][]int {
// 	ans := [][]int{}
// 	current := d.head
// 	last := d.head

// 	for {
// 		if last.next == nil {
// 			break
// 		}
// 		last = last.next
// 	}
// 	for {
// 		if current.next.next == nil {
// 			break
// 		}
// 		first := current.next
// 		second := last

// 		println("The value of current is ", current.value)
// 		for {

// 			if first == second || second.next == first {
// 				break
// 			}
// 			sum := first.value + second.value + current.value
// 			if sum == target {
// 				ans = append(ans, []int{first.value, second.value, current.value})
// 				first = first.next
// 				second = second.prev
// 			} else if sum > target {
// 				second = second.prev
// 			} else {
// 				first = first.next
// 			}
// 		}

// 		current = current.next
// 	}
// 	return ans
// }
