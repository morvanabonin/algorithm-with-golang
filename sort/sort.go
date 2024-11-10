package main

import (	
	"fmt"
)

func main() {

	// Create a slice of integers.
	nums := []int{9, 6, 5, 3, 8, 0, 4, 2, 7, 1}

	// Sort the slice of integers.
	sort(nums)

	// Display the sorted slice of integers.
	fmt.Println(nums)
}

func sort(nums []int) ([]int) {
	// the for start from 0 and continue until the length of the slice
	for i := 0; i < len(nums); i++ {
		// the get the next elemento to be compared and inserted in the sorted slice
		key := nums[i]
		// Insert nums[i] into the sorted sequence nums[1..i-1]
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j + 1] = nums[j]
			j = j - 1
		}
		nums[j + 1] = key
	}

	return nums
}

