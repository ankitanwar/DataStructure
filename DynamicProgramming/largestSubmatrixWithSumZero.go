package dynamicprogramming

import "fmt"

//LargestRectangularSumArea : Given a matrix we need to find the largest sub matrix whose sum is equal to zero
func LargestRectangularSumArea(matirx [][]int) {
	var leftSide int
	var rightSide int
	var topSide int
	var bottomSide int
	currentArea := (bottomSide - topSide) * (rightSide - leftSide)

	for left := 0; left < len(matirx[0]); left++ {
		rowSumarr := make([]int, len(matirx))
		for right := left; right < len(matirx[0]); right++ {
			for row := 0; row < len(matirx); row++ {
				rowSumarr[row] += matirx[row][right]
			}
			t := helperLargestRectangle(rowSumarr)
			if t[0] != -1 {
				t1 := (right - left) * (t[1] - t[0])
				if t1 > currentArea {
					fmt.Println("the value of t1 is ", t, right, left)
					currentArea = t1
					leftSide = left
					rightSide = right
					topSide = t[0]
					bottomSide = t[1]
				}
			}
		}
	}
	fmt.Println("The Top and bottom side is ", topSide, bottomSide)
	fmt.Println("The left and right side is ", leftSide, rightSide)
	fmt.Println("The area is ", currentArea)
}

func helperLargestRectangle(array []int) []int {
	var start int = -1
	var end int = -1
	prefix := []int{}
	var sum int
	for i := 0; i < len(array); i++ {
		sum += array[i]
		prefix = append(prefix, sum)
	}
	dict := make(map[int]int)
	for i := 0; i < len(prefix); i++ {
		if prefix[i] == 0 {
			start = 0
			end = i
		} else {
			_, found := dict[prefix[i]]
			if found {
				s := dict[prefix[i]]
				l := i - s
				if l > start+end {
					println("The vakue of s is ", s)
					start = s
					end = i
				}
			} else {
				dict[prefix[i]] = i
			}
		}
	}
	ans := []int{start, end}
	return ans

}
