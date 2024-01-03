package main

import "fmt"

func productExceptSelf(nums []int) []int {
	var accum = 1
	var result = make([]int, len(nums))
	for i := range nums {
		result[i] = accum
		accum *= nums[i]
	}
	accum = 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= accum
		accum *= nums[i]
	}
	return result
}

func main() {
	fmt.Printf("%v\n", productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Printf("%v\n", productExceptSelf([]int{-1, 1, 0, -3, 3}))
}
