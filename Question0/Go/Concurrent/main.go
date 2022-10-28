package main

import (
	"fmt"
	"sync"
)

func productOfArrayExceptSelf(nums []int) []int {
	var leftProd []int
	var rightProd []int = make([]int, len(nums))
	var result []int

	wg := sync.WaitGroup{}

	// [1, 2, 3, 4]
	// right = [24, 12, 4, 1]
	// left = [1, 1, 2, 6]

	// Iterate through the array and calculate both the
	leftProd = append(leftProd, 1) // Initialize the left side first element
	rightProd[len(nums)-1] = 1     // Initialize the right side last element

	wg.Add(2)
	go func() {
		for i := 1; i < len(nums); i++ {
			product := leftProd[i-1] * nums[i-1]
			leftProd = append(leftProd, product)
		}
		wg.Done()
	}()

	go func() {
		for i := len(nums) - 2; i >= 0; i-- {
			product := rightProd[i+1] * nums[i+1]
			rightProd[i] = product
		}
		wg.Done() // Tell the waitgroup this task is finished
	}()
	wg.Wait()

	// Last loop, multiply everything
	for i := 0; i < len(nums); i++ {
		product := leftProd[i] * rightProd[i]
		result = append(result, product)
	}

	return result
}

func main() {
	numbers := []int{-1, 1, 0, -3, 3}
	fmt.Println("Input: ", numbers)
	numbers2 := productOfArrayExceptSelf(numbers)
	fmt.Println("Output: ", numbers2)
}
