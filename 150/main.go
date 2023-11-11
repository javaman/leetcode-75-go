package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}

	fmt.Println(search(nums, 9))

	fmt.Println(search(nums, 2))
}
