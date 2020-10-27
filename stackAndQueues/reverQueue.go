package stackandqueues

//ReverseQueue : To reverse the queue
func ReverseQueue(q Queue) Queue {
	if len(q.arr) == 1 {
		return q
	}
	temp := q.Pop()
	q = ReverseQueue(q)
	q.Insert(temp)
	return q
}
