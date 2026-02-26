// ex8.01 - calculate the maximum value using intefaces
package main

import "fmt"

// represent the types of values we allow as input
type Number interface {
	int | float64
}

func findMaxGeneric[Num Number](nums []Num) Num {
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
	maxGenericInt := findMaxGeneric([]int{1, 32, 5, 8, 10, 11})
	fmt.Printf("max generic int: %v\n", maxGenericInt)

	maxGenericFloat := findMaxGeneric([]float64{1.1, 32.1, 5.1, 8.1, 10.1, 11.1})
	fmt.Printf("max generic float: %v\n", maxGenericFloat)
}
