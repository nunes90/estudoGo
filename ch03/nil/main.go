// The nil value

/*
`nil` is not a type but a special value in Go. It represents an empty value of no type. When working with pointers, maps, and intefaces, you need to be sure they are not `nil`. If you try to interact with a `nil` value, your code will crash.
*/
package main

import "fmt"

func main() {
	var message []string
	if message == nil {
		fmt.Println("error, unexpected nil value")
		return
	}
	fmt.Println(message)
}
