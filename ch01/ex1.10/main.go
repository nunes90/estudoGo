// Ex 1.10 - implementing shorthand operators
package main

import "fmt"

func main() {
	count := 5
	count += 5
	fmt.Println(count)

	// Increment the value by 1
	count++
	fmt.Println(count)

	// Decrement the value by 1
	count--
	fmt.Println(count)

	// Subtract and assign the result back to itself
	count -= 5
	fmt.Println(count)

	name := "John"
	name += " Smith"
	fmt.Println("Hello,", name)

}
