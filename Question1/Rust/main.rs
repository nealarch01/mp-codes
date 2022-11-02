fn max_profit(prices: Vec<i32>) -> i32 {
    let n = prices.len();
    let mut sell_prices = vec![0; prices.len()];
    let mut highest_price = prices[n - 1]; // Start at the very last
    sell_prices[n - 1] = highest_price;
    // Iterate backwards starting at n - 2 by using skip
    for (i, price) in prices.iter().enumerate().rev().skip(1) {
        if *price > highest_price {
            highest_price = *price;
        }
        sell_prices[i] = highest_price;
    }
    let mut max_profit = 0;
    for i in 0..n {
        let profit = sell_prices[i] - prices[i];
        if profit > max_profit {
            max_profit = profit;
        }
    }
    return max_profit;
}


fn main() {
    let prices = vec![7, 1, 5, 3, 6, 4];
    println!("{}", max_profit(prices));
}