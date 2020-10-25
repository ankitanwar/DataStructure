package dynamicprogramming

//FriendsPairing : To cal the number of ways in which a friends can pair
func FriendsPairing(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return FriendsPairing(n-1) + (n-1)*FriendsPairing(n-2)
}

//FriendsPairing2 : To solve the friends pairing problem dynamically
func FriendsPairing2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + (i-1)*dp[i-2]
	}

	return dp[n]
}
