package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    minPrice := prices[0]  // Keep track of the minimum price so far
    maxProfit := 0         // Initialize the maximum profit

    for i := 1; i < len(prices); i++ {
        // Calculate the profit if we sell on day i
        profit := prices[i] - minPrice

        // Update maxProfit if the current profit is greater
        if profit > maxProfit {
            maxProfit = profit
        }

        // Update minPrice if the current price is lower
        if prices[i] < minPrice {
            minPrice = prices[i]
        }
    }

    return maxProfit
}


func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
