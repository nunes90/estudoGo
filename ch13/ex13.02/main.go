// ex13.02 - using flags to say hello conditionally
package main

import (
	"flag"
	"fmt"
)

var (
	nameFlag  = flag.String("name", "Sam", "Name of the person to say hello to")
	quietFlag = flag.Bool("quiet", false, "Toggle to be quite when saying hello")
)

// Parse the flags and conditionally say hello pending the value of the quite flag:
func main() {
	flag.Parse()
	if !*quietFlag {
		greeting := fmt.Sprintf("Hello, %s! Welcome to the command line.", *nameFlag)
		fmt.Println(greeting)
	}
}
