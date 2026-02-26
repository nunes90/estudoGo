package main

import "fmt"

func findMaxInt(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

/*
 If we wanted to now find the maximum value for a different type of input, such as floating-point  values, then we’d have to add a new function containing duplicate logic:
*/

func findMaxFloat(nums []float64) float64 {
	if len(nums) == 0 {
		return -1
	}

	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func findMaxGeneric[Num int | float64](nums []Num) Num {
	if len(nums) == 0 {
		return -1
	}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max

}

func main() {

	maxInt := findMaxInt([]int{1, 32, 5, 8, 10, 11})
	fmt.Printf("max integer value: %v\n", maxInt)

	maxFloat := findMaxFloat([]float64{1.0, 32.0, 5.0, 8.0, 10.0, 11.0})
	fmt.Printf("max float value: %v\n", maxFloat)

	maxGenericInt := findMaxGeneric([]int{1, 32, 5, 8, 10, 10})
	fmt.Printf("max genericInt value: %v\n", maxGenericInt)

}
