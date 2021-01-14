package array

import (
	"fmt"
	"sort"
)

//MergeOverlappingInterval : Given an array where a[i]=a[i][0]--> start and a[i][1]--> end we need to merge all overlapping interval and return the array of non overlapping interval
func MergeOverlappingInterval(inervals [][]int) {
	sort.Slice(inervals, func(i, j int) bool { return inervals[i][0] <= inervals[j][0] })
	var ans [][]int
	i := 1
	current := inervals[0]
	for i < len(inervals) {
		if current[1] >= inervals[i][0] {
			f := min(current[0], inervals[i][0])
			e := max(current[1], inervals[i][1])
			current[0] = f
			current[1] = e
		} else {
			ans = append(ans, current)
			current = inervals[i]
		}
		i++
	}
	ans = append(ans, current)
	fmt.Println(ans)
}
