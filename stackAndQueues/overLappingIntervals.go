package stackandqueues

import "math"

//MergeOverlappingIntervals : To merge the overlapping intervals
func MergeOverlappingIntervals(array [][]int) [][]int {
	if len(array) == 0 || len(array) == 1 {
		return array
	}
	slow := 0
	fast := 1
	for fast < len(array) {
		if array[slow][1] >= array[fast][0] {
			array[slow][0] = int(math.Min(float64(array[slow][0]), float64(array[fast][0])))
			array[slow][1] = int(math.Max(float64(array[slow][1]), float64(array[fast][1])))
			array = append(array[:slow+1], array[fast+1:]...)

		} else {
			slow++
			fast++
		}
	}

	return array
}
