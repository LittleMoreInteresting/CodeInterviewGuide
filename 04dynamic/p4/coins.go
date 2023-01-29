package p4

func coins5(arr []int, aim int) int {

	dp := make([]int, aim+1)
	dp[0] = 1
	for i := 0; i < len(arr); i++ {
		for j := 1; j <= aim; j++ {
			if j-arr[i] >= 0 {
				dp[j] = dp[j] + dp[j-arr[i]]
			}
		}
	}

	return dp[aim]
}
