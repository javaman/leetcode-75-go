package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1

	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Printf("%v\n", nums1)

	nums1 = []int{1}
	nums2 = []int{}
	merge(nums1, 1, nums2, 0)
	fmt.Printf("%v\n", nums1)

	nums1 = []int{0}
	nums2 = []int{1}
	merge(nums1, 0, nums2, 1)
	fmt.Printf("%v\n", nums1)
}
