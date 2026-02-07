// activity2.02 - implementing FizzBuzz
package main

import (
	"fmt"
	"strconv"
)

func main() {

	for i := 1; i <= 100; i++ {
		out := ""
		if i%3 == 0 {
			out += "Fizz"
		}
		if i%5 == 0 {
			out += "Buzz"
		}
		if out == "" {
			out = strconv.Itoa(i) // convert int to string
		}
		fmt.Println(out)
	}
}
