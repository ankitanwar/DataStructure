package dynamicprogramming

//PermutationCoff : To Calculate the value of permuation cofficient
func PermutationCoff(n, k int) int {
	if k > n {
		return 0
	}
	if k == 0 || n == k {
		return 1
	}
	if k == 1 {
		return n
	}
	temp := n
	ans := PermutationCoff(n-1, k-1)
	ans = ans * temp
	return ans

}
