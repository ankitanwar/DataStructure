package stackandqueues

//InsertBottom : To insert the value at the bottom of the stack
func (s *Stack) InsertBottom(val int) {
	if len(s.arr) == 0 {
		s.Push(val)
		return
	}
	temp := s.Pop()
	println("The length of stack is  ", len(s.arr))
	s.InsertBottom(val)
	s.Push(temp)

}
