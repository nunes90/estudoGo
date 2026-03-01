package main

import "fmt"

/*
 * Verb | Meaning
 * %d   | Prints an integer in base-10
 * %f   | Prints a floating point number, default width, default precision
 * %t   | Prints a bool type
 * %s   | Prints a string type
 * %v   | Prints the value in default format
 * %b   | Prints the base two\binary representation
 * %x   | Prints the hex representation
 */

func main() {
	fname := "Joe"
	gpa := 3.75
	hasJob := true
	age := 24
	hourlyWage := 45.53
	fmt.Printf("%s has a gpa of %f.\n", fname, gpa)
	fmt.Printf("He has a job equals %t.\n", hasJob)
	fmt.Printf("He is %d earning %v per hour.\n", age, hourlyWage)
}
