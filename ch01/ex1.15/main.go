// Ex 1.15 - function design with pointers
package main

import "fmt"

func addd5Value(count int) {
	count += 5
	fmt.Println("add5Value: ", count)
}

func add5Point(count *int) {
	*count += 5
	fmt.Println("add5Point: ", *count)
}

func main() {
	var count int
	addd5Value(count)
	fmt.Println("add5Value post: ", count)

	add5Point(&count)
	fmt.Println("add5Point post: ", count)
}
