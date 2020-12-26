package bit

//GreyCode : To write grey code of a given number
func GreyCode(n int) []string {
	if n == 1 {
		ans := []string{"0", "1"}
		return ans
	}
	t := GreyCode(n - 1)
	var ans []string
	for i := 0; i < len(t); i++ {
		t1 := "0" + string(t[i])
		ans = append(ans, t1)
	}
	for i := len(t) - 1; i >= 0; i-- {
		t2 := "1" + string(t[i])
		ans = append(ans, t2)
	}
	return ans

}
