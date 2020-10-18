package sorting

// MajoirtyElement : it returns the element whose frequency is greater than n/2
func MajoirtyElement(arr []int) (majority int) {
	majority = -1
	count := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] == majority {
			count++
		} else {
			count--
		}

		if count == 0 {
			majority = arr[i]
			count = 1
		}
	}

	newCount := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == majority {
			newCount++
		}
	}

	if newCount > len(arr)/2 {
		return majority
	}
	return -1

}
