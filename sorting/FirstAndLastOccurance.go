package sorting

//FirstAndLastOccurance : return the first and last occurance of the given element index is 1 based
func FirstAndLastOccurance(arr []int, tofind int) (i, j int) {
	a := -1
	b := -1
	for index, val := range arr {
		if val == tofind && a == -1 {
			a = index
			b = index
		} else if val == tofind && a != -1 {
			b = index
		}
	}

	return a, b
}
