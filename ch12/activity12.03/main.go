// activity12.03 - Measuring elapsed time
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(2 * time.Second)
	end := time.Now()

	length := end.Sub(start)
	fmt.Printf("The execution took exactly %v seconds!\n", length)

	// elapsed := time.Since(start)
	// fmt.Println(elapsed)

}
