// ex2.11 - using a break and continue to control loops
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for {
		// from rand package to pick a random number between 0 and 8
		r := rand.Intn(8)

		if r%3 == 0 {
			fmt.Println("Skip")
			continue
		} else if r%2 == 0 {
			fmt.Println("Stop")
			break
		}
		fmt.Println(r)
	}
}
