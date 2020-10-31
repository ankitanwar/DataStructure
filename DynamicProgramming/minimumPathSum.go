package dynamicprogramming

//MinimumPath : It will give you the minimum cost to reach the end
func MinimumPath(cost [][]int, m, n int) int {
	if m < 0 || n < 0 {
		return 99999
	} else if m == 0 && n == 0 {
		return cost[m][n]
	}
	return cost[m][n] + min(MinimumPath(cost, m-1, n-1), MinimumPath(cost, m-1, n), MinimumPath(cost, m, n-1))
}
