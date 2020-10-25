package dynamicprogramming

import "math"

//GoldMine : To calculate the maximum amount of gold that can be obtained
func GoldMine(array [][]int) int {

	//Creating the answer matrix
	ans := [][]int{}
	for i := 0; i < len(array); i++ {
		temp := []int{}
		for j := 0; j < len(array[0]); j++ {
			temp = append(temp, -1)
		}
		ans = append(ans, temp)
	}
	for j := len(ans[0]) - 1; j >= 0; j-- {
		for i := 0; i < len(ans); i++ {
			if j == len(ans[0])-1 {
				ans[i][j] = array[i][j]
				continue
			}
			if i == 0 {
				ans[i][j] = array[i][j] + int(math.Max(float64(ans[i][j+1]), float64(ans[i+1][j+1])))
				continue
			}
			if i == len(ans)-1 {
				ans[i][j] = array[i][j] + int(math.Max(float64(ans[i][j+1]), float64(ans[i-1][j+1])))
				continue
			} else {
				ans[i][j] = array[i][j] + max(ans[i][j+1], ans[i-1][j+1], ans[i+1][j+1])
			}
		}
	}
	gg := -99
	for i := 0; i < len(array); i++ {
		current := ans[i][0]
		gg = newmax(current, gg)
	}
	return gg
}
func newmax(a, b int) int {
	if a > b {
		return a
	}
	return b

}
func max(a, b, c int) int {
	if a > b && a > c {
		return a
	}
	if b > a && b > c {
		return b
	}
	return c

}
