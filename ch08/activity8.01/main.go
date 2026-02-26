// Activity8.01 - a minimum value
package main

import "fmt"

func findMinGeneric[Num int | float64](mins []Num) Num {
	if len(mins) == 0 {
		return -1
	}

	min := mins[0]
	for _, num := range mins {
		if num < min {
			min = num
		}
	}
	return min
}

func main() {
	minInt := findMinGeneric([]int{1, 32, 5, 8, 10, 11})
	fmt.Printf("max integer value: %v\n", minInt)

	minFloat := findMinGeneric([]float64{1.1, 32.1, 5.1, 8.1, 10.1, 11.1})
	fmt.Printf("max float value: %v\n", minFloat)

}
