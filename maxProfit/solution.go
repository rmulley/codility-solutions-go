package solution

import (
	"math"
) //import

func Solution(A []int) int {
	var (
		maxProfit, profitHere int
		minPrice              int
	) //var

	// If there are one or zero prices, then there is nothing to sell
	if len(A) < 2 {
		return maxProfit
	} //if

	// Initialize minimum price to the first day's price
	minPrice = A[0]

	for _, price := range A {
		// Calculate the profit if stock was sold today
		profitHere = int(math.Max(
			0.0,
			float64(price-minPrice),
		))

		// Keep track of lowest stock price so far
		minPrice = int(math.Min(
			float64(minPrice),
			float64(price),
		))

		// Check if selling today results in a new maximum
		maxProfit = int(math.Max(
			float64(maxProfit),
			float64(profitHere),
		))
	} //for

	return maxProfit
} //Solution
