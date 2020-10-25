package stackandqueues

// PermutationPossible : To check whether the permutation is possible or not
func PermutationPossible(input, output []int) bool {
	temp := []int{}
	j := 0
	i := 0
	for i < len(output) {
		if len(temp) == 0 {
			temp = append(temp, output[i])
			i++
		}
		if temp[len(temp)-1] == input[j] {
			temp = temp[:len(temp)-1]
			j++
		} else {
			temp = append(temp, output[i])
			i++
		}

	}
	if len(temp) == 0 && j == len(input) {
		return true
	}
	return false
}
