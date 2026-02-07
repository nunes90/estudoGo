// Ex2.01 - a simple if statement
package main

import "fmt"

func main() {
	input := 5

	if input%2 == 0 {
		fmt.Println(input, "is even")
	}
	if input%2 == 1 {
		fmt.Println(input, "is odd")
	}
}
