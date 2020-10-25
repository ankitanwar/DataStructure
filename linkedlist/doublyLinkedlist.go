package linkedlist

//Dnode : node of a doubly linked list
type Dnode struct {
	Prev  *Dnode
	Next  *Dnode
	Value string
}

// Dlinkedlist : doubly linked list
type Dlinkedlist struct {
	Head *Dnode
}

// Insert : To insert the value in the doubly linked list
func (d *Dlinkedlist) Insert(val string) *Dnode {
	current := d.Head
	n := &Dnode{}
	n.Value = val
	if d.Head == nil {
		d.Head = n
		return d.Head
	}
	for {
		if current.Next == nil {
			current.Next = n
			n.Prev = current
			return n
		}
		current = current.Next
	}
}

// Print : To The doubly linked list
func (d *Dlinkedlist) Print() {
	current := d.Head
	ans := ""

	for {
		if current == nil {
			break
		}
		ans += current.Value + "-->"
		current = current.Next

	}
	println(ans)
}

//Ddlete : To delete the node in the doubly linked list
func (d *Dlinkedlist) Ddlete(s string) {
	current := d.Head
	if current.Value == s {
		temp := current.Next
		d.Head = temp
		current.Next = nil
		return
	}

	for {
		if current.Value == s {
			if current.Next == nil {
				prev := current.Prev
				prev.Next = nil
				break

			} else {
				prev := current.Prev
				nxt := current.Next
				nxt.Prev = current.Prev.Prev
				prev.Next = current.Next.Next
				break
			}
		}

		current = current.Next

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
