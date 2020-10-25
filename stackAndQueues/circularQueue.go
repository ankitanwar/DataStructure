package stackandqueues

import "babbar/linkedlist"

//CircularQueue : To build and ciruclar queue
type CircularQueue struct {
	Head  *linkedlist.Node
	Tail  *linkedlist.Node
	Count int
}

//Enqueue : To insert into the circular queue
func (c *CircularQueue) Enqueue(value string) string {
	if c.Count == 0 {
		return "Queue is Full"
	}
	n := &linkedlist.Node{}
	n.Value = value
	if c.Head == nil {
		c.Head = n
		c.Tail = n
		c.Tail.Next = c.Head
		c.Count--
		return c.Head.Value
	}
	c.Tail.Next = n
	c.Tail = c.Tail.Next
	c.Tail.Next = c.Head
	c.Count--
	return c.Tail.Value

}

//Dequeue : To delete the value from the queue
func (c *CircularQueue) Dequeue() string {
	if c.Head == nil {
		return "cannot delete from the empty list"
	}
	if c.Head.Next == c.Head {
		removed := c.Head.Value
		c.Head = nil
		c.Count++
		return removed
	}
	temp := c.Head
	c.Tail.Next = c.Head.Next
	c.Head = c.Head.Next
	temp.Next = nil
	return temp.Value

}

//Print : To print the circular queue
func (c *CircularQueue) Print() {
	if c.Head == nil {
		println("Queue is empty")
		return
	}
	current := c.Head
	ans := current.Value + "-->"
	current = current.Next
	for {
		if current == nil || current == c.Head {
			break
		}
		ans += current.Value + "-->"
		current = current.Next
	}

	print(ans)

}
