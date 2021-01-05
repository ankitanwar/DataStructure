package array

//FindDuplicate : Given an array of nums containing n+1 integers where each integer is the range 1,n inclusive
func FindDuplicate(nums []int) int {
	var ans int

	a := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > a {
			t := nums[i] % a
			nums[t] += a
		} else {
			t := nums[i]
			nums[t] += a
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 2*a {
			ans = i
			nums[i] %= a
		} else {
			nums[i] %= a
		}
	}

	return ans
}
