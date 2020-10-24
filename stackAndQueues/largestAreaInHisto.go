package stackandqueues

import "fmt"

//LargestAreaInHistogram : To find the largest area
func LargestAreaInHistogram(array []int) {

	// We have to find the nearest smallest to the right and nearest smallest to the left in this

	NSL := []int{}
	NSR := []int{}

	stack1 := []int{}
	stack2 := []int{}

	for i := len(array) - 1; i >= 0; i-- {
		for {
			if len(stack1) == 0 || array[stack1[len(stack1)-1]] < array[i] {
				break

			}
			println("The values are ", array[len(stack1)-1], array[i])
			stack1 = stack1[:len(stack1)-1]
		}

		if len(stack1) == 0 {
			NSL = append(NSL, -1)
			stack1 = append(stack1, i)
		} else {
			NSL = append(NSL, stack1[len(stack1)-1])
			stack1 = append(stack1, i)
		}
	}

	ans := reverse(NSL)

	for i := 0; i < len(array); i++ {
		for {
			if len(stack2) == 0 || array[stack2[len(stack2)-1]] < array[i] {
				break
			}
			stack2 = stack2[:len(stack2)-1]
		}
		if len(stack2) == 0 {
			NSR = append(NSR, len(array)+1)
			stack2 = append(stack2, i)
		} else {
			NSR = append(NSR, stack2[len(stack2)-1])
			stack2 = append(stack2, i)
		}
	}

	for i := 0; i < len(ans); i++ {
		fmt.Print(ans[i])
		fmt.Print(" ")
	}

	fmt.Println("")

	for i := 0; i < len(array); i++ {
		fmt.Print(NSR[i])
		fmt.Print(" ")
	}

}

func reverse(arr []int) []int {
	low := 0
	high := len(arr) - 1
	for {
		if low >= high {
			break
		}
		arr[low], arr[high] = arr[high], arr[low]
		low++
		high--
	}

	return arr
}
