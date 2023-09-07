package challenges

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		// Update minPrice if the current price is lower
		if price < minPrice {
			minPrice = price
		}

		// Update maxProfit if selling at the current price yields higher profit
		currentProfit := price - minPrice
		if currentProfit > maxProfit {
			maxProfit = currentProfit
		}
	}

	return maxProfit
}
