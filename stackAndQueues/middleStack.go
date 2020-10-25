package stackandqueues

import (
	"babbar/linkedlist"
)

//Midstack : to implement the stack which perfom mid operation in O(1)
type Midstack struct {
	head   *linkedlist.Dnode
	tail   *linkedlist.Dnode
	middle *linkedlist.Dnode
	count  int
}

//Insert : To insert the value in the stack
func (m *Midstack) Insert(value string) {
	n := &linkedlist.Dnode{}
	n.Value = value
	if m.head == nil {
		m.head = n
		m.tail = m.head
		m.middle = m.head
		m.count++
		return
	}
	m.tail.Next = n
	n.Prev = m.tail
	m.tail = m.tail.Next
	m.count++
	if m.count%2 != 0 {
		m.middle = m.middle.Next
	}
}

//GetMiddle : To get the middle value of the stack
func (m *Midstack) GetMiddle() string {
	return m.middle.Value
}

//Delete : to delete the value in the stack
func (m *Midstack) Delete() {
	prev := m.tail.Prev
	prev.Next = nil
	m.tail.Prev = nil
	m.tail = prev
	m.count--
	if m.count%2 == 0 {
		println("This is working")
		m.middle = m.middle.Prev
	}
}

//Print : to print the value of the stack
func (m *Midstack) Print() {
	current := m.head
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
