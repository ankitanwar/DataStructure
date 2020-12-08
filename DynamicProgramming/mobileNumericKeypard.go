package dynamicprogramming

import "fmt"

//NumericKeypad : We have to count the total possible numbers we can make we can only move left,right,up and down
func NumericKeypad(n int) {
	canBePressed := [][]int{}
	canBePressed = append(canBePressed, []int{0, 8})       //Options available for key 0
	canBePressed = append(canBePressed, []int{1, 2, 4})    // options available for key 1
	canBePressed = append(canBePressed, []int{1, 2, 3, 5}) //options avaiable for key 2
	canBePressed = append(canBePressed, []int{2, 3, 6})
	canBePressed = append(canBePressed, []int{1, 4, 5, 7})
	canBePressed = append(canBePressed, []int{2, 4, 5, 6, 8})
	canBePressed = append(canBePressed, []int{3, 5, 6, 9})
	canBePressed = append(canBePressed, []int{4, 7, 8})
	canBePressed = append(canBePressed, []int{5, 7, 8, 9, 0})
	canBePressed = append(canBePressed, []int{6, 8, 9})

	dp := [][]int{}
	for i := 0; i <= n; i++ {
		temp := make([]int, 10)
		dp = append(dp, temp)
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < 10; j++ {
			if i == 0 {
				dp[i][j] = 0
			} else if i == 1 {
				dp[i][j] = 1
			} else {
				var temp int
				for k := 0; k < len(canBePressed[j]); k++ {
					num := canBePressed[j][k]
					temp += dp[i-1][num]
				}
				dp[i][j] = temp
			}
		}
	}
	fmt.Println(dp)
	var ans int
	for i := 0; i < len(dp[0]); i++ {
		ans += dp[len(dp)-1][i]
	}
	fmt.Println(ans)

}
