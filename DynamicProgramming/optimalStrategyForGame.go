package dynamicprogramming

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
