// Ex1.12 - zero values
package main

import (
	"fmt"
	"time"
)

/*
%v  - Any value. Use this if you don't care about the type you're printing.
%+v - Values with extra information, such as struct field names.
%#v - Go syntax, such as %+v with the addition of the name of the type of the variable.
%T  - Print the variable's type.
%d  - Decimal (base 10)
%s  - String
*/

func main() {

	var count int
	fmt.Printf("Count:	  : %#v \n", count)

	var discount float64
	fmt.Printf("Discount  : %#v \n", discount)

	var debug bool
	fmt.Printf("Debug     : %#v \n", debug)

	var message string
	fmt.Printf("Message   : %#v \n", message)

	var emails []string
	fmt.Printf("Emails    : %#v \n", emails)

	var startTime time.Time
	fmt.Printf("Start     : %#v \n", startTime)
}
