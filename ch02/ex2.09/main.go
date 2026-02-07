// ex2.09 - looping over arrays and slices
package main

import "fmt"

func main() {
	names := []string{"Jim", "Jane", "Joe", "June"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
