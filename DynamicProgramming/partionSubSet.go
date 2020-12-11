package dynamicprogramming

//PartionSubset : We have to calc whether we can divide the subset into two equal parts or not
func PartionSubset(arr []int) bool {
	var temp int
	for i := 0; i < len(arr); i++ {
		temp += arr[i]
	}
	if temp%2 != 0 {
		return false
	}
	found := temp / 2
	matrix := [][]int{}
	for i := 0; i <= found; i++ {
		temp := []int{}
		for j := 0; j <= len(arr); j++ {
			temp = append(temp, -1)
		}
		matrix = append(matrix, temp)
	}

	ans := partionSubsetDP(&matrix, arr, 0, 0, found)
	return ans
}
func helperPartion(arr []int, sumTillnow, currentIndex, found int) bool {
	if found == sumTillnow {
		return true
	}
	if currentIndex == len(arr) {
		return false
	}
	selected := helperPartion(arr, sumTillnow+arr[currentIndex], currentIndex+1, found)
	notselected := helperPartion(arr, sumTillnow, currentIndex+1, found)
	return selected || notselected
}
func partionSubsetDP(matrix *[][]int, arr []int, sumTillnow, currentIndex, found int) bool {
	if currentIndex == len(arr) || sumTillnow > found {
		return false
	}
	if sumTillnow == found {
		(*matrix)[sumTillnow][currentIndex] = 1
		return true
	}
	if (*matrix)[sumTillnow][currentIndex] != -1 {
		println("This is working")
		if (*matrix)[sumTillnow][currentIndex] == 0 {
			return false
		}
		return true
	}
	slct := partionSubsetDP(matrix, arr, sumTillnow+arr[currentIndex], currentIndex+1, found)
	notSlct := partionSubsetDP(matrix, arr, sumTillnow, currentIndex+1, found)
	temp := slct || notSlct
	if temp == true {
		(*matrix)[sumTillnow][currentIndex] = 1
	} else {
		(*matrix)[sumTillnow][currentIndex] = 0
	}
	return temp
}
