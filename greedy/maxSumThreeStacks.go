package greedy

import "fmt"

//MaxSumOf3Stack : We have given a 3 stack we have to find out whether it is possible to make sum of three stack equal if it is possible then the sum would be maximum otherwise the sum would be 0
func MaxSumOf3Stack(stack1, stack2, stack3 []int) {
	var sum1 int
	var sum2 int
	var sum3 int

	for i := 0; i < len(stack1); i++ {
		sum1 += stack1[i]
	}
	for i := 0; i < len(stack2); i++ {
		sum2 += stack2[i]
	}
	for i := 0; i < len(stack3); i++ {
		sum3 += stack3[i]
	}
	for {
		if len(stack1) == 0 || len(stack2) == 0 || len(stack3) == 0 {
			fmt.Println(0)
			break
		}
		if sum1 == sum2 && sum2 == sum3 {
			fmt.Println(sum1)
			break
		}
		if sum1 > sum2 && sum1 > 3 {
			sum1 -= stack1[0]
			stack1 = stack1[1:]

		} else if sum2 > sum1 && sum2 > sum3 {
			sum2 -= stack2[0]
			stack2 = stack2[1:]
		} else {
			sum3 -= stack3[0]
			stack3 = stack3[1:]
		}
	}
}
