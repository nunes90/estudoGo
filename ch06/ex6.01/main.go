// ex6.01  - runtime errors while adding numbers
package main

import "fmt"

func main() {
	nums := []int{2, 4, 6, 8}

	total := 0
	// for i := 0; i <= 10; i++ {
	for i := range nums {
		total += nums[i]
	}
	fmt.Println("Total: ", total)
}
