package main

import (
    "fmt"
)

func maxProfit(prices []int) int {
    if len(prices) == 1 { // If there is only one value in the array, no profit is made
        return 0
    }
    maxPrices := make([]int, len(prices)) // This array holds the greatest numbers right of the current index
    highest := prices[len(prices) - 1]  // Start at the last index
    for i := len(maxPrices) - 2; i >= 0; i-- { // Iterate backwards filling the max price
        if prices[i] > highest { 
            highest = prices[i]
        }
        maxPrices[i] = highest
    }
    greatestProfit := 0
    for i := 0; i < len(prices); i++ {
        // Calculate the result
        profit := maxPrices[i] - prices[i] 
        if profit > greatestProfit {
            greatestProfit = profit
        }
    }
    return greatestProfit
}

func main() {
    prices := []int{7, 1, 5, 3, 6, 4}
    result := maxProfit(prices)
    fmt.Println(result)
}
