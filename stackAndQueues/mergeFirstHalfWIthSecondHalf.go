package stackandqueues

//InterleaveFirstHalfWithSecond : to merge first half with the second half
func InterleaveFirstHalfWithSecond(q Queue) {
	// Its only valid for input number of integers though we can modify it for output also
	queue2 := Queue{}
	ans := Queue{}

	mid := q.GetLen() / 2

	for i := 0; i < mid; i++ {
		current := q.Pop()
		queue2.Insert(current)
	}
	q.Print()
	for {
		if queue2.GetLen() == 0 || q.GetLen() == 0 {
			break
		}
		first := queue2.Pop()
		second := q.Pop()
		ans.Insert(first)
		ans.Insert(second)
	}
	if q.GetLen() != 0 {
		current := q.Pop()
		ans.Insert(current)
	}
	println("This is workign")
	ans.Print()

}
