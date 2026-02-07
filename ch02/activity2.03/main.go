// activity2.03 - bubble sort
package main

import "fmt"

func main() {

	nums := []int{3, 5, 6, 2, 9, 10, 7, 1, 8, 4}
	fmt.Println("Before: ", nums)

	// Sort the values using swapping.
	for swapped := true; swapped; {
		swapped = false
		for i := 1; i < len(nums); i++ {
			if nums[i-1] > nums[i] {
				nums[i-1], nums[i] = nums[i], nums[i-1]
				swapped = true
			}
		}
	}

	fmt.Println("After: ", nums)
}
