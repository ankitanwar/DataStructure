package dynamicprogramming

import "fmt"

// MaximumSumRectangle : Given a matrix we need to find its maximum sum rectangle
func MaximumSumRectangle(matrix [][]int) {
	var rectangleSumAns int = -99999
	var rowSum []int
	var leftAns int
	var Top int
	var Bottom int
	var rightAns int
	for left := 0; left < len(matrix[0]); left++ {
		rowSum = make([]int, len(matrix))
		for right := left; right < len(matrix[0]); right++ {
			for row := 0; row < len(matrix); row++ {
				rowSum[row] += matrix[row][right]
			}
			t := kadnesAlgo(rowSum)
			if t[0] > rectangleSumAns {
				rectangleSumAns = t[0]
				leftAns = left
				rightAns = right
				Top = t[1]
				Bottom = t[2]
			}

		}
	}
	fmt.Println("The Left and Right co-ordiantes", leftAns, rightAns)
	fmt.Println("The Top and Bottom co-ordiantes", Top, Bottom)
	fmt.Println(rectangleSumAns)
}

//ouput array contains arr[0]=maxSubArrayPossible arr[1]=TopIndex arr[2]=BottomIndex
func kadnesAlgo(array []int) []int {
	var maxSum int = -1e9
	var top int
	var bottom int
	for i := 0; i < len(array); i++ {
		currentSum := maxSum + array[i]
		if array[i] > currentSum {
			top = i
			maxSum = array[i]
		} else if currentSum > maxSum {
			maxSum = currentSum
			bottom = i
		}
	}
	arr := []int{maxSum, top, bottom}
	return arr
}
