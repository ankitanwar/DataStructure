package sorting

//MissingAndRepeating : it returns the value of  missing and repeating element
func MissingAndRepeating(arr []int) (miss, repeat int) {
	miss = -1
	repeat = -1

	myDict := make(map[int]int)

	for index, val := range arr {
		_, found := myDict[val]
		if found {
			repeat = val
		} else {
			myDict[val] = index
		}
	}

	for i := 1; i < len(arr)+1; i++ {
		_, found := myDict[i]
		if found == false {
			miss = i
			break
		}
	}

	return miss, repeat
}
