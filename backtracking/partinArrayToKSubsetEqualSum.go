package backtracking

import "fmt"

// DivideIntoKSubset : To divide the array into k subset such that sum of all the k subset all equal
func DivideIntoKSubset(arr []int, k int) {
	mat := [][]int{}
	for i := 0; i < k; i++ {
		t := []int{}
		mat = append(mat, t)
	}
	ksetHelper(arr, &mat, 0)

}

func ksetHelper(arr []int, k *[][]int, currentIndex int) {
	if currentIndex == len(arr) {
		checkPrint((*k))
		return
	}
	for i := 0; i < len((*k)); i++ {
		if len((*k)[i]) != 0 {
			(*k)[i] = append((*k)[i], arr[currentIndex])
			ksetHelper(arr, k, currentIndex+1)
			(*k)[i] = (*k)[i][:len((*k)[i])-1]
		} else {
			(*k)[i] = append((*k)[i], arr[currentIndex])
			ksetHelper(arr, k, currentIndex+1)
			(*k)[i] = (*k)[i][:len((*k)[i])-1]
			break
		}
	}
}

func checkPrint(k [][]int) {
	var sum int
	var same bool = true
	for i := 0; i < len(k[0]); i++ {
		sum += k[0][i]
	}
	for j := 1; j < len(k); j++ {
		var temp int
		for i := 0; i < len(k[j]); i++ {
			temp += k[j][i]
		}
		if temp != sum {
			same = false
			break
		}
	}
	if same == true {
		fmt.Println(k)
	}
}
