package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i, hasFlower := range flowerbed {
		if n == 0 {
			return true
		}
		if hasFlower == 0 {
			leftIsEmpty := (i-1 < 0) || (flowerbed[i-1] == 0)
			rightIsEmpty := (i+1 >= len(flowerbed)) || (flowerbed[i+1] == 0)
			if leftIsEmpty && rightIsEmpty {
				flowerbed[i] = 1
				n--
			}
		}
	}
	return n == 0
}

func main() {
	fmt.Printf("%t\n", canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Printf("%t\n", canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
}
