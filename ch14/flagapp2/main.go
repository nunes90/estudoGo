package main

import (
	"flag"
	"fmt"
)

func main() {
	var v int
	flag.IntVar(&v, "value", -1, "Needs a value for the flag.")
	flag.Parse()
	fmt.Println(v)
}

/*
This code does the same as the previous snippet, however, here’s a quick breakdown:
• First, we define an integer variable v
• Use its reference as the first parameter of the IntVar function
• Parse the flags
• Print the v variable, which now does not need to be dereferenced as it is not the flag but an
actual integer
*/
