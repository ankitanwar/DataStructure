package stackandqueues

//NonRepeating : To find the non repeating character in the queue
func NonRepeating(s string) []string {
	ans := []string{}
	queue := Queue{}
	dict := make(map[string]int)
	for i := 0; i < len(s); i++ {
		current := string(s[i])
		_, found := dict[current]
		if found == false {
			dict[current] = 1
			queue.Insert(current)
		} else {
			dict[current]++
		}
		for {
			if queue.GetLen() == 0 || dict[queue.Peek()] == 1 {
				break
			}
			queue.Pop()
		}
		if queue.GetLen() == 0 {
			ans = append(ans, "0")
		} else {
			ans = append(ans, queue.Peek())
		}
	}

	return ans
}
