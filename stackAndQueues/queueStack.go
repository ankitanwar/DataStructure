package stackandqueues

//QueueUsingStack : Implement queue using stack
type QueueUsingStack struct {
	array1 []int
	array2 []int
}

//Insert : To insert the value into the stack
func (q *QueueUsingStack) Insert(value int) {
	q.array1 = append(q.array1, value)
}

//Delete : To delete the value from the stack
func (q *QueueUsingStack) Delete() int {
	var removed int
	if len(q.array1) == 0 {
		return -1
	}
	for {
		if len(q.array1) == 1 {
			removed = q.array1[0]
			q.array1 = q.array1[:1]
			break
		}

		current := q.array1[len(q.array1)-1]
		q.array1 = q.array1[:len(q.array1)-1]
		q.array2 = append(q.array2, current)

	}
	for {
		if len(q.array2) == 0 {
			break
		}
		current := q.array2[len(q.array1)-1]
		q.array2 = q.array2[:len(q.array1)-1]
		q.array1 = append(q.array1, current)
	}

	return removed
}
