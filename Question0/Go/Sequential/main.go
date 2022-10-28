package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	var leftProd []int
	var rightProd []int
	var result []int

	// [1, 2, 3, 4]
	// right = [24, 12, 4, 1]
	// left = [1, 1, 2, 6]

	// Iterate through the array and calculate both the
	leftProd = append(leftProd, 1)   // Initialize the left side first element
	rightProd = append(rightProd, 1) // Initialize the right side first element

	for i := 1; i < len(nums); i++ {
		product := leftProd[i-1] * nums[i-1]
		leftProd = append(leftProd, product)
		rightProd = append(rightProd, 1)
	}

	for i := len(nums) - 2; i >= 0; i-- {
		product := rightProd[i+1] * nums[i+1]
		rightProd[i] = product
	}

	// Last loop, multiply everything
	for i := 0; i < len(nums); i++ {
		product := leftProd[i] * rightProd[i]
		result = append(result, product)
	}

	return result
}

func main() {
	nums := []int{-1, 1, 0, -3, 3}
	result := productExceptSelf(nums)
	fmt.Println(result)
}
