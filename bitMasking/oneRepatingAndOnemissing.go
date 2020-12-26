package bit

import "fmt"

//OneRepeatingAndOneMissing : Given an array we need to find out the one character which is repeating twice and one character which is missing
func OneRepeatingAndOneMissing(arr []int) {
	t := []int{}
	for i := 1; i <= len(arr); i++ {
		t = append(t, i)
	}
	t1 := 0
	for i := 0; i < len(arr); i++ {
		t1 = t1 ^ arr[i] ^ t[i]
	}
	rmsb := t1 & -t1
	set1 := []int{}
	set2 := []int{}

	for i := 0; i < len(arr); i++ {
		if arr[i]&rmsb == rmsb {
			set1 = append(set1, arr[i])
		} else {
			set2 = append(set2, arr[i])
		}
	}
	for i := 0; i < len(t); i++ {
		if t[i]&rmsb == rmsb {
			set1 = append(set1, t[i])
		} else {
			set2 = append(set2, t[i])
		}
	}
	value1 := 0
	value2 := 0
	for i := 0; i < len(set1); i++ {
		value1 = value1 ^ set1[i]
	}
	for i := 0; i < len(set2); i++ {
		value2 = value2 ^ set2[i]
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == value1 {
			fmt.Println("The missing value is ", value2)
			fmt.Println("The duplicate value is ", value1)
			break
		}
		if arr[i] == value2 {
			fmt.Println("The missing value is ", value1)
			fmt.Println("The duplicate value is ", value2)
			break
		}
	}

}
