package stackandqueues

import (
	"fmt"
	"strconv"
)

//Postfix : To evulate the postfix expression
func Postfix(str string) {
	s := Stack{}
	for i := 0; i < len(str); i++ {
		current := string(str[i])
		switch current {
		case "*":
			first := s.Pop()
			second := s.Pop()

			ans := int(first) * int(second)
			s.Push(ans)

		case "+":
			first := s.Pop()
			second := s.Pop()

			ans := int(second) + int(first)

			s.Push(ans)

		case "-":
			first := s.Pop()
			second := s.Pop()

			ans := int(second) - int(first)
			s.Push(ans)

		case "/":
			first := s.Pop()
			second := s.Pop()

			ans := int(second) / int(first)
			s.Push(ans)
		default:
			new, _ := strconv.Atoi(current)
			s.Push(new)
		}
	}

	ans := s.arr[len(s.arr)-1]
	fmt.Println(ans)
}
