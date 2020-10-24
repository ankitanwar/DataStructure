package stackandqueues

import "babbar/linkedlist"

//StackUsingDQueues : To implement stack using deque
type StackUsingDQueues struct {
	head *linkedlist.Dnode
	tail *linkedlist.Dnode
}

//Insert : To insert the Values
func (s *StackUsingDQueues) Insert(val string) {
	n := &linkedlist.Dnode{}
	n.Value = val
	current := s.head
	if current == nil {
		s.head = n
		s.tail = n
		return
	}
	s.tail.Next = n
	n.Prev = s.tail
	s.tail = s.tail.Next
}

//Delete : To delete Item
func (s *StackUsingDQueues) Delete() string {
	if s.head == nil {
		return "Nothing to be deleted"
	}
	if s.head.Next == nil {
		removed := s.head
		s.head = nil
		return removed.Value
	}

	removed := s.tail.Value
	prev := s.tail.Prev
	s.tail.Prev = nil
	prev.Next = nil
	s.tail = prev

	return removed

}
