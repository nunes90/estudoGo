// Ex 1.02 - declaring a variable using var
package main

import "fmt"

// package level scope
var foo string = "bar"

func main() {
	var baz string = "qux"

	fmt.Println(foo, baz)
}
