package main

import ( 
    "fmt"
)


func maxProfit(prices []int) int {
    greatest := int(0)
    for i := 0; i < len(prices) - 1; i++ {
        for j := i + 1; j < len(prices); j++ {
            profit := prices[j] - prices[i]
            if profit <= greatest {
                continue
            }
            greatest = profit
        }
    }
    return greatest
}


func main() {
    prices := []int{7, 1, 5, 3, 6, 4}
    result := maxProfit(prices)
    fmt.Println(result)
}
