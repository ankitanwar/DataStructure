package dynamicprogramming

//SubsequencesLessThanK : to count all the subsequences whose product is less than k
func SubsequencesLessThanK(array []int, k, length int) int {
	current := 1
	left := 0
	total := 0
	right := 0
	for right < len(array) {
		current *= array[right]
		for current >= k {
			current /= array[left]
			left++
		}
		total += right - left + 1
		println("the value of total is ", total)
		right++
	}
	return total
}
