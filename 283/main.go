package main

func moveZeroes(nums []int) {
	next := 0
	for _, v := range nums {
		if v != 0 {
			nums[next] = v
			next++
		}
	}
	for ; next < len(nums); next++ {
		nums[next] = 0
	}
}
