package stackandqueues

//MinimumSum : to count the square of the minimum sum in the given characters
func MinimumSum(s string) int {
	dict := make(map[string]int)

	for i := 0; i < len(s); i++ {
		current := string(s[i])
		_, found := dict[current]
		if found == false {
			dict[current] = 1
		} else {
			dict[current]++
		}
	}
	max := 0
	ans := 0
	for _, value := range dict {
		current := value
		if current > max {
			ans += max * max
			max = current
		} else {
			ans += value * value
		}
	}
	max--
	ans += max * max

	return ans
}
