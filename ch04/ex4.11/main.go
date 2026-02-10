// ex4.11 - Using make to control the capacity of a slice
package main

import "fmt"

func genSlices() ([]int, []int, []int) {
	var s1 []int // slice of length 0

	s2 := make([]int, 10) // slice of length 10

	s3 := make([]int, 10, 50) // slice of length 10, capacity 50

	return s1, s2, s3
}

func main() {
	s1, s2, s3 := genSlices()
	fmt.Printf("s1: len = %v cap = %v\n", len(s1), cap(s1))
	fmt.Printf("s2: len = %v cap = %v\n", len(s2), cap(s2))
	fmt.Printf("s3: len = %v cap = %v\n", len(s3), cap(s3))
}
