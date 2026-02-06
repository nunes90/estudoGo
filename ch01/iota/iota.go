// Enums
package main

import "fmt"

// const (
// 	sunday    = 0
// 	monday    = 1
// 	tuesday   = 2
// 	wednesday = 3
// 	thursday  = 4
// 	friday    = 5
// 	saturday  = 6
// )

/*
Using `iota` makes enums easier to create and maintain, especially if you need to add anew value to the middle of the code later. Order matters when using `iota` as it is an identifier that tells the Go compiler to start the first value at 0 and increment by 1 for each subsequente value. With `iota`, you can skip using _, start with a different offset, and even use more complicated calculations.
*/

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
}
