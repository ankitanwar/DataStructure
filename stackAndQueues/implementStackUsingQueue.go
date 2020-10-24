package stackandqueues

//StackUsingQueues : Implement stack using queue
type StackUsingQueues struct {
	queue1 []int
	queue2 []int
}

//Insert : To insert the value into the queue
func (s *StackUsingQueues) Insert(value int) {
	s.queue1 = append(s.queue1, value)
}

//Delete : to delete the item
func (s *StackUsingQueues) Delete() int {
	var removed int
	for {
		if len(s.queue1) == 1 {
			removed = s.queue1[0]
			break
		}
		s.queue2 = s.queue1[:1]
		s.queue1 = s.queue1[1:]
	}
	for {
		if len(s.queue2) == 0 {
			break
		}
		s.queue1 = s.queue2[:1]
		s.queue2 = s.queue2[1:]

	}

	return removed
}
