package sorting

// Count : it returns the number of perfect square withing the range n
func Count(n int) int {
	count := 0

	for i := 1; i < n; i++ {
		if i*i < n {
			count++
		} else {
			break
		}
	}

	return count
}
