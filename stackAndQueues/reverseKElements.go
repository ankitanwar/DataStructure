package stackandqueues

//ReverseKElements : To reverse the k elements in the queue
func ReverseKElements(k int, q Queue) Queue {
	queue2 := Queue{}

	q = ReverseQueue(q)
	length := q.GetLen()

	limit := length - k
	for i := 0; i < limit; i++ {
		current := q.Pop()
		queue2.Insert(current)

	}
	if queue2.GetLen() == 0 {
		return q
	}
	queue2 = ReverseQueue(queue2)
	for {
		if queue2.GetLen() == 0 {
			break
		}
		current := queue2.Pop()
		q.Insert(current)
	}

	return q

}
