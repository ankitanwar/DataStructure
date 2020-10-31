package heaps

//SumBetweenK1K2 : to calc the sum between the k1 and k2 smallest number
func SumBetweenK1K2(arr []int, k1, k2 int) int {
	sum := 0
	first := KthSmallestElement(arr, k1)
	second := KthSmallestElement(arr, k2)

	for i := 0; i < len(arr); i++ {
		if arr[i] > first.(int) && arr[i] < second.(int) {
			sum += arr[i]
		}
	}

	return sum

}
