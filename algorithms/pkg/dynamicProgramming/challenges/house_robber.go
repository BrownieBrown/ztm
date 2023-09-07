package challenges

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	// Create a DP array to store the maximum money that can be robbed up to house i
	dp := make([]int, n)

	// Initialization: Robbing the first house means you get whatever is in it.
	// Robbing up to the second house means you take the max of the first and second houses.
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	// Loop through the array and populate dynamicProgramming
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}

	// The maximum money that can be robbed from all houses is stored in dynamicProgramming[n-1]
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
