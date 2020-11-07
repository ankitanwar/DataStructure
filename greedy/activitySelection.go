package greedy

import (
	"sort"
)

type meeting struct {
	start  int
	finish int
}

//MaxMeetingInRoom : It will return the max number of meeting possible inside the room
func MaxMeetingInRoom(start, finish []int) int {
	count := 1

	array := []meeting{}
	for i := 0; i < len(start); i++ {
		temp := meeting{
			start:  start[i],
			finish: finish[i],
		}
		array = append(array, temp)
	}

	sort.Slice(array[:], func(i, j int) bool {
		return array[i].finish < array[j].finish
	})

	current := array[0].finish
	for i := 1; i < len(array); i++ {
		if array[i].start >= current {
			count++
			current = array[i].finish
		}
	}

	return count
}
