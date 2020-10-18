package sorting

// PairWithGivenDifference : return pair with given difference
func PairWithGivenDifference(arr []int, target int) []int {
	myDict := make(map[int]int)

	for index, val := range arr {
		find := target - val
		indexx, found := myDict[find]
		if found {
			return ([]int{indexx, index})
		}
		myDict[val] = index
	}

	return []int{-1, -1}
}
