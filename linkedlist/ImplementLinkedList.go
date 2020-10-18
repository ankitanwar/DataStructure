package linkedlist

//Node : it represent the node of the linked list
type Node struct {
	value string
	next  *Node
}

//Linkedlist : it represnt the linked list
type Linkedlist struct {
	head *Node
}

//Insert : insert new element in the linked list
func (l *Linkedlist) Insert(val string) {
	n := Node{}
	n.value = val

	if l.head == nil {
		l.head = &n
		return
	}

	current := l.head
	for {
		if current.next == nil {
			break
		}
		current = current.next
	}

	current.next = &n

}

// Delete : It is used to remove the node from this linked list
func (l *Linkedlist) Delete(val string) string {
	current := l.head
	if current.value == val {
		current = current.next
		l.head = current
		return "Element has been deleted Successfully"
	}
	prev := current
	current = current.next
	for {
		if current == nil {
			return "Element not found in the LinkedList"
		}
		if current.value == val && current.next != nil {
			prev.next = current.next.next
			return "Element has been deleted Successfully"
		} else if current.value == val && current.next == nil {
			prev.next = nil
			return "Element has been deleted Successfully"
		}
		prev = current
		current = current.next
	}
}

//Print : To print all the elements in the linked list
func (l *Linkedlist) Print() {
	current := l.head
	ans := ""
	for {
		if current == nil {
			break
		}
		ans += (current.value + "-->")
		current = current.next
	}

	println(ans)
}
