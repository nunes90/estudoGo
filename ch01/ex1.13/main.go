// Ex1.13 - getting a pointer
package main

import (
	"fmt"
	"time"
)

func main() {
	var count1 *int    // nil
	count2 := new(int) // non-nil
	countTemp := 5
	count3 := &countTemp
	t := &time.Time{}

	fmt.Printf("count1: %#v\n", count1)
	fmt.Printf("count2: %#v\n", count2)
	fmt.Printf("count3: %#v\n", count3)
	fmt.Printf("time : %#v\n", t)
}
