package dynamicprogramming

//BuyAndSell2Transaction : It will return the maximum profit that can be made in case 2 transaction is allowed for the given price array
func BuyAndSell2Transaction(prices []int) int {
	var ans int

	soldAtPoint := []int{}
	currentMin := prices[0]
	soldAtPoint = append(soldAtPoint, 0) // cant made profit it sold today because there were no previous values
	currentMaxProfit := 0
	for i := 1; i < len(prices); i++ {
		if currentMin > prices[i] {
			currentMin = prices[i]
		}
		tempProfit := prices[i] - currentMin
		if tempProfit > currentMaxProfit {
			currentMaxProfit = tempProfit
		}
		soldAtPoint = append(soldAtPoint, currentMaxProfit)
	}

	boughtAtPoint := []int{}
	boughtAtPoint = append(boughtAtPoint, 0) //No future information available after this
	currentSellPoint := prices[len(prices)-1]
	currentProfitIfBoughToday := 0

	for i := len(prices) - 2; i >= 0; i-- {
		if currentSellPoint < prices[i] {
			currentSellPoint = prices[i]
		}
		tempProfit := currentSellPoint - prices[i]
		if tempProfit > currentProfitIfBoughToday {
			currentProfitIfBoughToday = tempProfit
		}
		boughtAtPoint = append(boughtAtPoint, currentProfitIfBoughToday)
	}
	//we are travesing in opposite direction thats why we have reversed the array
	reverse(&boughtAtPoint)

	for i := 0; i < len(boughtAtPoint); i++ {
		if soldAtPoint[i]+boughtAtPoint[i] > ans {
			ans = soldAtPoint[i] + boughtAtPoint[i]
		}
	}
	return ans
}

func reverse(array *[]int) {
	low := 0
	high := len(*array) - 1
	for {
		if low >= high {
			break
		}
		(*array)[low], (*array)[high] = (*array)[high], (*array)[low]
		low++
		high--
	}
}
