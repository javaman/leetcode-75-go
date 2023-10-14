package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, 0, len(candies))
	var max = -1
	for _, value := range candies {
		if value > max {
			max = value
		}
	}
	for _, value := range candies {
		result = append(result, value+extraCandies >= max)
	}
	return result
}

func main() {
	fmt.Printf("%v", kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
}
