package main

import (
	"flag"
	"fmt"
)

func main() {
	v := flag.Int("value", -1, "Needs a value for the flag.")
	flag.Parse()
	fmt.Println(*v)
}

/*
 First, we define the main package.
 Then we import the flag and fmt packages.
 The v variable will reference the value for either -value or --value.
 The initial value of *v is the default value of -1 before calling flag.Parse()
 After defining the flags, you must call flag.Parse() to parse the defined flags into the
 command line.
 Calling flag.Parse() places the argument for -value into *v.
 Once you have called the flag.Parse() function, the flags will be available.
*/
