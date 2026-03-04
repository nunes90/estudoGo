// ex13.01 - saying hello using a name passed as an argument
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: go run main.go <name>")
		return
	}

	name := args[1]

	greeting := fmt.Sprintf("Hello, %s! Welcome to the command line.", name)
	fmt.Println(greeting)
}
