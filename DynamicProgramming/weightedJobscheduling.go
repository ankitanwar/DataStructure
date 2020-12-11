package dynamicprogramming

import (
	"fmt"
	"sort"
)

type job struct {
	startTime  int
	finishTime int
	profit     int
}

//WeightJobScheduling : to schedule the jobs in order to obtain the max profit
func WeightJobScheduling(finishTime, startTime, profit []int, currentIndex int) {
	arr := []job{}
	for i := 0; i < len(finishTime); i++ {
		temp := &job{
			finishTime: finishTime[i],
			startTime:  startTime[i],
			profit:     profit[i],
		}
		arr = append(arr, *temp)
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i].finishTime < arr[j].finishTime })
	fmt.Println(arr)
	ans := helperJobScheduling(arr, 0)
	fmt.Println(ans)
}

func nonOverlapping(arr []job, currentIndex int) int {
	if currentIndex+1 == len(arr) {
		return -1
	}
	for i := currentIndex + 1; i < len(arr); i++ {
		if arr[i].startTime >= arr[currentIndex].finishTime {
			return i
		}
	}
	return -1
}

func helperJobScheduling(arr []job, currentIndex int) int {
	if currentIndex == len(arr) {
		return 0
	}
	included := arr[currentIndex].profit
	nxt := nonOverlapping(arr, currentIndex)
	if nxt != -1 {
		included += helperJobScheduling(arr, nxt)
	}
	excluded := helperJobScheduling(arr, currentIndex+1)
	return maximum(included, excluded)

}
