package dynamicprogramming

//MaxDiffOf0and1 : it will return the maximum difference of  0 and 1
func MaxDiffOf0and1(s string) int {
	var ans int
	zero := 0
	one := 0

	for i := 0; i < len(s); i++ {
		current := string(s[i])
		if current == "0" {
			zero++
		} else {
			one++
		}
		if one >= zero {
			one = 0
			zero = 0
		} else {
			ans = maximum(ans, zero-one)
		}
	}

	return ans
}
