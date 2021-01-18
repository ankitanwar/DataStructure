package array

//SubArrayWithSum0 : Given an array we need to find out if there exist an subarray whose sum is equal to 0
func SubArrayWithSum0(arr []int) {
	dict := make(map[int]int)
	var current int
	var found bool
	for i := 0; i < len(arr); i++ {
		current += arr[i]
		_, contains := dict[current]
		if contains {
			found = true
			println("The range is ", dict[current]+1, i)
			break
		} else {
			dict[current] = i
		}
	}
	if found == false {
		println("No there exist no sub array whose sum is equal to the 0")
	}
}
