package stackandqueues

//RemovedMid :to remove the middle element from the stack
func (s *Stack) RemovedMid(k int) {
	if len(s.arr) == 0 {
		return
	}
	if k == 1 {
		s.Pop()
		return
	}
	temp := s.Pop()
	s.RemovedMid(k - 1)
	s.Push(temp)
}
