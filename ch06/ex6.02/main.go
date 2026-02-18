// ex6.02 - a semantic error with walking distance
package main

import "fmt"

func main() {
	km := 2
	// if km > 2 {
	if km >= 2 {
		fmt.Println("Take the car")
	} else {
		fmt.Println("Going to walk today")
	}
}
