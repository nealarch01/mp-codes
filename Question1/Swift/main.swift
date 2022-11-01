func maxProfit(_ prices: [Int]) -> Int {
    let n: Int = prices.count
    if n == 1 {
        return 0
    }
    // Below is an array contains the greatest selling price right of the current index
    var sellPrices = Array<Int>(repeating: 0, count: n) 
    var highest: Int = prices[n - 1]
    sellPrices[n - 1] = highest
    
    for i in stride(from: n - 2, to: -1, by: -1) {
        if prices[i] > highest {
            highest = prices[i]
        }
        sellPrices[i] = highest
    }
    
    var maxProfit: Int = 0
    for i in 0 ..< n {
        let profit: Int = sellPrices[i] - prices[i]
        if profit > maxProfit {
            maxProfit = profit
        }
    }
    return maxProfit
}

func main() {
    let prices: [Int] = [7, 1, 5, 3, 6, 12, 2, 11, 9, 4, 10, 2, 8]
    let maxProfit: Int = maxProfit(prices)
    print(maxProfit)
}


main()