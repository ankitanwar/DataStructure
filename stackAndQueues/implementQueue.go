package stackandqueues

import "fmt"

// Queue : To declare the queue
type Queue struct {
	arr []string
}

//Insert : TO insert the item into the queue
func (q *Queue) Insert(val string) {
	q.arr = append(q.arr, val)
}

// Pop : To delete the item from the queue
func (q *Queue) Pop() string {
	remove := q.arr[0]
	q.arr = q.arr[1:]
	return remove
}

// Print : To print all the values in the queue
func (q *Queue) Print() {
	for i := 0; i < len(q.arr); i++ {
		fmt.Print(q.arr[i])
		fmt.Print(" ")
	}
	fmt.Println("")
}

//GetLen : To return the length of the queue
func (q *Queue) GetLen() int {
	return len(q.arr)
}

func (q *Queue) Peek() string {
	return q.arr[0]
}
