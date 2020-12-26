package bit

import "fmt"

//AllRepeatingExcept2 : Given an array we need to find out those 2 number which are not repaeting
func AllRepeatingExcept2(arr []int) {
	t := 0
	for i := 0; i < len(arr); i++ {
		t = t ^ arr[i]
	}
	set1 := []int{}
	set2 := []int{}
	rmb := t & -t
	for i := 0; i < len(arr); i++ {
		if arr[i]&rmb == rmb {
			set1 = append(set1, arr[i])
		} else {
			set2 = append(set2, arr[i])
		}
	}
	ans1 := 0
	ans2 := 0
	for i := 0; i < len(set1); i++ {
		ans1 = ans1 ^ set1[i]
	}
	for j := 0; j < len(set2); j++ {
		ans2 = ans2 ^ set2[j]
	}
	fmt.Println(ans1, "-->", ans2)

}
