package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	first := math.MaxInt32
	second := math.MaxInt32
	for _, value := range(nums) {
		if (value <= first) {
			first = value
		} else if (value <= second) {
			second = value
		} else {
			return true
		}
	}
	return false	   
}

func main() {
	fmt.Printf("%v\n", increasingTriplet([]int{1, 2, 3, 4, 5}))
	fmt.Printf("%v\n", increasingTriplet([]int{5, 4, 3, 2, 1}))
	fmt.Printf("%v\n", increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
}