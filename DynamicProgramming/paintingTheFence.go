package dynamicprogramming

//PaintingTheFence : To calc the number of ways in which the fens can be painted
func PaintingTheFence(n, k int) int {
	same := k * 1
	different := k * (k - 1)

	total := same + different

	for i := 3; i < n; i++ {
		same = different * 1
		different = total * (k - 1)
		total = same + total
	}

	return total

}
