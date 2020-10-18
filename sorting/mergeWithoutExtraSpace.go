package sorting

import "sort"

//Merge : this will merge the two arrays firsst array will contain the smaller elements than second will contains the greater
func Merge(arr, arr2 []int) {
	i := 0
	j := 0

	for i < len(arr) && j < len(arr2) {
		if arr2[j] < arr[i] {
			arr2[j], arr[i] = arr[i], arr2[j]
			sort.Ints(arr2)
		} else {
			i++
		}
	}
}
