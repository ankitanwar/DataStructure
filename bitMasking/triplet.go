package bit

import "fmt"

//FindTriplet : Given an array we need to find triplet such that 0<i<j<=k
func FindTriplet(arr []int) {
	/*
			And they should full fil this condition also
			 X = arr[i] ^ arr[i+1] ^ ... ^ arr[j - 1]
		    Y = arr[j] ^ arr[j+1] ^ ... ^ arr[k]
	*/

	var count int

	for i := 0; i < len(arr)-1; i++ {
		fmt.Println("The value of i is ", i)
		current := arr[i]
		for j := i + 1; j < len(arr); j++ {
			current = current ^ arr[j]
			if current == 0 {
				count += j - i
			}
		}

	}
	fmt.Println(count)

}
