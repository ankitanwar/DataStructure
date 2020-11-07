package greedy

import "sort"

type job struct {
	deadline int
	profit   int
}

//JobSequence : tO schedule the job in order to obtain the maximum profit
func JobSequence(deadline, profit []int) int {
	var ans int
	array := []job{}
	for i := 0; i < len(deadline); i++ {
		temp := job{
			deadline: deadline[i],
			profit:   profit[i],
		}
		array = append(array, temp)
	}
	sort.Slice(array[:], func(i, j int) bool { return array[i].profit > array[j].profit })
	available := [100]bool{}

	for i := 0; i < len(array); i++ {
		for j := array[i].deadline - 1; j >= 0; j-- {
			if available[j] == false {
				available[j] = true
				ans += array[i].profit
				break
			}
		}
	}
	return ans

}
