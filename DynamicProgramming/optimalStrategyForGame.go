package dynamicprogramming

import "fmt"

//OptimalStrategyForgame : It will return the strategy in order to maximize the amount in order to win the game
func OptimalStrategyForgame(coins []int, start, end int) int {
	if start == end {
		choosen := coins[start]
		println("First choosen", choosen)
		return choosen
	}
	if start+1 == end {
		choosen := maximum(coins[start], coins[end])
		println("second choosen", choosen)
		return choosen
	}
	choosen := maximum(coins[start]+minimum(OptimalStrategyForgame(coins, start+2, end), OptimalStrategyForgame(coins, start+1, end-1)), coins[end]+minimum(OptimalStrategyForgame(coins, start+1, end-1), OptimalStrategyForgame(coins, start, end-2)))
	println("final choosen", choosen)
	return choosen
}

//OptimalStrategyForgameDP : strategy to maximize the profit
func OptimalStrategyForgameDP(coins []int, start, end int) {
	dp := [][]int{}
	for i := 0; i <= len(coins); i++ {
		temp := []int{}
		for j := 0; j <= len(coins); j++ {
			temp = append(temp, -1)
		}
		dp = append(dp, temp)
	}
	ans := strategyHelpeR(&dp, start, end, coins)
	fmt.Println(ans)
}

func strategyHelpeR(dp *[][]int, start, end int, coins []int) int {
	if (*dp)[start][end] != -1 {
		return (*dp)[start][end]
	}
	if start == end {
		(*dp)[start][end] = coins[start]
		return (*dp)[start][end]
	}
	if start+1 == end {
		(*dp)[start][end] = maximum(coins[start], coins[end])
		return (*dp)[start][end]
	}
	selected := maximum(coins[start]+minimum(strategyHelpeR(dp, start+2, end, coins), strategyHelpeR(dp, start+1, end-1, coins)), coins[end]+minimum(strategyHelpeR(dp, start+1, end-1, coins), strategyHelpeR(dp, start, end-2, coins)))
	(*dp)[start][end] = selected
	return (*dp)[start][end]

}
