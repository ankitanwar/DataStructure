package dynamicprogramming

import "fmt"

//MinimumPath : It will give you the minimum cost to reach the end
func MinimumPath(cost [][]int, m, n int) int {
	if m < 0 || n < 0 {
		return 99999
	} else if m == 0 && n == 0 {
		return cost[m][n]
	}
	return cost[m][n] + min(MinimumPath(cost, m-1, n-1), MinimumPath(cost, m-1, n), MinimumPath(cost, m, n-1))
}

type path struct {
	path   string
	row    int
	column int
}

//MinimumPathDP : To find the minimum cost required to reach the end and we have to print the path also
func MinimumPathDP(cost [][]int) {
	dp := [][]int{}
	for i := 0; i < len(cost); i++ {
		t := make([]int, len(cost[0]))
		dp = append(dp, t)
	}
	//the idea is to traverse from the last and keep adding the value and then do the bfs to find the answer

	for i := len(dp) - 1; i >= 0; i-- {
		for j := len(dp[0]) - 1; j >= 0; j-- {
			if i == len(dp)-1 && j == len(dp[0])-1 {
				dp[i][j] = cost[i][j]
			} else if i == len(dp)-1 {
				dp[i][j] = cost[i][j] + dp[i][j+1]
			} else if j == len(dp[0])-1 {
				dp[i][j] = cost[i][j] + dp[i+1][j]
			} else {
				dp[i][j] = cost[i][j] + minimum(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	arr := []path{}
	t := &path{
		path:   "",
		row:    0,
		column: 0,
	}
	arr = append(arr, *t)
	for len(arr) > 0 {
		current := arr[0]
		arr = arr[1:]
		if current.row+1 == len(dp) && current.column+1 == len(dp[0]) {
			fmt.Println(current.path)
		} else if current.row+1 == len(dp) {
			t := path{
				path:   current.path + "H",
				row:    current.row,
				column: current.column + 1,
			}
			arr = append(arr, t)

		} else if current.column+1 == len(dp[0]) {
			t := path{
				path:   current.path + "V",
				row:    current.row + 1,
				column: current.column,
			}
			arr = append(arr, t)
		} else {
			if dp[current.row][current.column+1] == dp[current.row+1][current.column] {
				t1 := path{
					path:   current.path + "H",
					row:    current.row,
					column: current.column + 1,
				}
				arr = append(arr, t1)
				t2 := path{
					path:   current.path + "V",
					row:    current.row + 1,
					column: current.column,
				}
				arr = append(arr, t2)
			} else if dp[current.row][current.column+1] < dp[current.row+1][current.column] {
				t := path{
					path:   current.path + "H",
					row:    current.row,
					column: current.column + 1,
				}
				arr = append(arr, t)

			} else {
				t := path{
					path:   current.path + "V",
					row:    current.row + 1,
					column: current.column,
				}
				arr = append(arr, t)
			}
		}

	}

}
