package main

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	max := 0
	for i < j {
		var newArea int
		if height[i] < +height[j] {
			newArea = (j - i) * height[i]
			i++
		} else {
			newArea = (j - i) * height[j]
			j--
		}
		if newArea > max {
			max = newArea
		}
	}
	return max
}

func main() {
	println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	println(maxArea([]int{1, 1}))
}
