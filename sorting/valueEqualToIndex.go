package sorting

//ValueEqualToIndex : return the number of element whose index is equal to the element
func ValueEqualToIndex(arr []int) []int {
	ans := []int{}
	for index, val := range arr {
		if index+1 == val {
			ans = append(ans, val)
		}
	}
	return ans
}
