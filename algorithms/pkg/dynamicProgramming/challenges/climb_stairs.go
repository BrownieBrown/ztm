package challenges

func climbStairsMemo(n int, memo map[int]int) int {
	if n <= 2 {
		return n
	}

	if val, exists := memo[n]; exists {
		return val
	}

	memo[n] = climbStairsMemo(n-1, memo) + climbStairsMemo(n-2, memo)

	return memo[n]
}

func climbStairs(n int) int {
	memo := make(map[int]int)
	return climbStairsMemo(n, memo)
}
