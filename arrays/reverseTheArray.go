package array

//ReverseTheArray : Given an array we need to reverse it
func ReverseTheArray(array []int) {
	start := 0
	end := len(array) - 1
	for {
		if start >= end {
			break
		}
		array[start], array[end] = array[end], array[start]
		start++
		end--
	}

}
