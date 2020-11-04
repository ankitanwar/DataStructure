package linkedlist

import "strconv"

//Node : it represent the node of the linked list
type Node struct {
	Value interface{}
	Next  *Node
}

//Linkedlist : it represnt the linked list
type Linkedlist struct {
	Head *Node
}

//Insert : insert new element in the linked list
func (l *Linkedlist) Insert(val interface{}) {
	n := Node{}
	n.Value = val

	if l.Head == nil {
		l.Head = &n
		return
	}

	current := l.Head
	for {
		if current.Next == nil {
			break
		}
		current = current.Next
	}

	current.Next = &n

}

// Delete : It is used to remove the node from this linked list
func (l *Linkedlist) Delete(val string) string {
	current := l.Head
	if current.Value == val {
		current = current.Next
		l.Head = current
		return "Element has been deleted Successfully"
	}
	prev := current
	current = current.Next
	for {
		if current == nil {
			return "Element not found in the LinkedList"
		}
		if current.Value == val && current.Next != nil {
			prev.Next = current.Next.Next
			return "Element has been deleted Successfully"
		} else if current.Value == val && current.Next == nil {
			prev.Next = nil
			return "Element has been deleted Successfully"
		}
		prev = current
		current = current.Next
	}
}

//Print : To print all the elements in the linked list
func (l *Linkedlist) Print() {
	current := l.Head
	ans := ""
	ans += strconv.Itoa(current.Value.(int)) + "-->"
	current = current.Next
	for {
		if current == nil || current == l.Head {
			break
		}
		ans += strconv.Itoa(current.Value.(int)) + "-->"
		current = current.Next
	}

	println(ans)
}
