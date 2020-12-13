package dynamicprogramming

import "fmt"

//CatalanNumber : To calc the catalan number upto given n
func CatalanNumber(n int) {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			dp[i] = 1
		} else {
			x := 0
			y := i - 1
			var ans int = 0
			for {
				if y < 0 {
					break
				}
				temp := dp[x] * dp[y]
				ans += temp
				x++
				y--
			}
			dp[i] = ans
		}
	}
	fmt.Println(dp)
}

// catalan number:- 0 --> 1   1 -->1   2-> c0.c1 + c1.c0   3-> c0.c2+c1.c1+c2.c0    4-> c0.c3+c1.c2+c2.c1+c3.c0
//countNumberOfBST
//Count number of mountains both are the varitaion of catalan number
//given a circle with 2n points connect the rod between that such that no rod intersect each other
